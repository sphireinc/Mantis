// Package uuid generates RFC4122-compliant UUIDs, and provides functions to marshal the UUIDs into their canonical form of 16 bytes and unmarshal them from a variety of formats. See https://tools.ietf.org/html/rfc4122.
package uuid

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

// import "github.com/jet/go-mantis/uuid"

// ErrInvalidFormat is returned if the textual representation being unmarshaled is not a valid UUID
var ErrInvalidFormat = errors.New("uuid: invalid format")

// UUID is a Universally unique identifier as described in RFC4122/DCE1.1
type UUID [16]byte

// Nil UUID is an uuid with all zeros
var Nil = UUID{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// RFC4122 Predefined Namespace UUIDs https://tools.ietf.org/html/rfc4122#appendix-C
var (
	// NamespaceDNS {6ba7b810-9dad-11d1-80b4-00c04fd430c8}
	NamespaceDNS = UUID{0x6b, 0xa7, 0xb8, 0x10, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}

	// NamespaceURL {6ba7b811-9dad-11d1-80b4-00c04fd430c8}
	NamespaceURL = UUID{0x6b, 0xa7, 0xb8, 0x11, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}

	// NamespaceOID {6ba7b812-9dad-11d1-80b4-00c04fd430c8}
	NamespaceOID = UUID{0x6b, 0xa7, 0xb8, 0x12, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}

	// NamespaceX500 {6ba7b814-9dad-11d1-80b4-00c04fd430c8}
	NamespaceX500 = UUID{0x6b, 0xa7, 0xb8, 0x14, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}
)

// Version gets the version number of the UUID
func (u UUID) Version() byte {
	M := u[6]
	M &= 0xF0
	M >>= 4

	return M
}

// SetVersion sets the version number of the UUID
func (u *UUID) SetVersion(ver byte) {
	//  xxxx xxxx :    0xXX
	//  0000 1111 : && 0x0F
	//  vvvv 0000 : || 0xV0
	//  ----------:
	//  vvvv xxxx      0xVX
	u[6] = (u[6] & 0x0F) | ((ver & 0x0F) << 4)
}

// SetDCESecurity sets the security information in a Time-Based UUID as defined by
// DCE 1.1 Authentication and Security Services: http://pubs.opengroup.org/onlinepubs/9696989899/chap5.htm#tagcjh_08_02_01_01
func (u *UUID) SetDCESecurity(domain byte, id uint32) {
	if ver := u.Version(); ver < 0 || ver > 2 {
		return
	}
	u.SetVersion(2)
	u.SetVariant(VariantRFC4122) // Sets the variant bits to `10xx`
	binary.BigEndian.PutUint32(u[0:4], id)
	u[9] = domain
}

// DCESecurity gets the DCE security information on a Time-based UUID
// See: http://pubs.opengroup.org/onlinepubs/9696989899/chap5.htm#tagcjh_08_02_01_01
func (u *UUID) DCESecurity() (domain byte, id uint32) {
	if u.Version() != 2 {
		return
	}
	id = binary.BigEndian.Uint32(u[0:4])
	domain = u[9]
	return
}

// 100ns ticks since 1582-10-15 to 1970-1-1
var uuidEpocStart = uint64(time.Date(1582, time.October, 15, 0, 0, 0, 0, time.UTC).Unix() * int64(-1e7))

func timestamp64() uint64 {
	return uuidEpocStart + uint64(time.Now().UnixNano()/100)
}

func timestamp32() uint64 {
	ts := timestamp64()
	return ts & 0xFFFFFFFF00000000 // Drop time_low
}

func tick16(clk uint16) uint16 {
	return clk + 1
}

func tickHigh8(clk uint16) uint16 {
	c := (clk & 0xFF00) >> 8
	return (c + 1) << 8
}

type rfc4122Generator struct {
	initOnce sync.Once
	mu       sync.Mutex
	hwAddr   [6]byte // MAC
	lastTime uint64
	clockSeq uint16

	clockTickFunc func(uint16) uint16
	timestampFunc func() uint64
}

func (r *rfc4122Generator) String() string {
	marshaledStruct, err := json.Marshal(r)
	if err != nil {
		return err.Error()
	}
	return string(marshaledStruct)
}

var rfc4122 = rfc4122Generator{
	clockTickFunc: tick16,
	timestampFunc: timestamp64,
}
var dceSec = rfc4122Generator{
	clockTickFunc: tickHigh8,
	timestampFunc: timestamp32,
}

func (g *rfc4122Generator) newV1() (UUID, error) {
	hw, ts, clk, err := g.tick()
	if err != nil {
		return Nil, err
	}
	var u UUID
	binary.BigEndian.PutUint32(u[0:], uint32(ts))
	binary.BigEndian.PutUint16(u[4:], uint16(ts>>32))
	binary.BigEndian.PutUint16(u[6:], uint16(ts>>48))
	binary.BigEndian.PutUint16(u[8:], clk)
	copy(u[10:], hw[:])

	u.SetVersion(1)
	u.SetVariant(VariantRFC4122)

	return u, nil
}

func (g *rfc4122Generator) tick() ([6]byte, uint64, uint16, error) {
	var err error
	g.initOnce.Do(func() {
		hwAddr, e := g.getHwAddr()
		if e != nil {
			err = e
			return
		}
		g.hwAddr = hwAddr
		cbs := make([]byte, 2)
		if _, f := rand.Read(cbs); f != nil {
			err = fmt.Errorf("uuid: init clock seq failed: %v", f)
		}
		g.clockSeq = binary.BigEndian.Uint16(cbs)
	})
	if err != nil {
		return [6]byte{}, 0, 0, err
	}
	g.mu.Lock()
	defer g.mu.Unlock()
	ts := g.timestampFunc()
	if g.lastTime >= ts {
		g.clockSeq = g.clockTickFunc(g.clockSeq)
	}
	g.lastTime = ts
	return g.hwAddr, g.lastTime, g.clockSeq, nil
}

func (g *rfc4122Generator) getRandHwAddr() ([6]byte, error) {
	var hwAddr [6]byte
	if _, err := rand.Read(hwAddr[:]); err != nil {
		return [6]byte{}, err
	}
	// Set multicast bit
	hwAddr[0] |= 0x01
	return hwAddr, nil
}

func (g *rfc4122Generator) getHwAddr() ([6]byte, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return g.getRandHwAddr()
	}
	for _, iface := range ifaces {
		if len(iface.HardwareAddr) >= 6 {
			var hwAddr [6]byte
			copy(hwAddr[:], iface.HardwareAddr[0:6])

			// Set multicast bit
			hwAddr[0] |= 0x01

			return hwAddr, nil
		}
	}
	return g.getRandHwAddr()
}

// Variant of the UUID version
// See RFC4122 Section 4.1.1
// - https://tools.ietf.org/html/rfc4122#section-4.1.1
type Variant byte

// UnknownVariant is an unknown uuid variant
const UnknownVariant Variant = 0xFF

const (
	// VariantNCS is reserved, NCS backward compatibility.
	VariantNCS Variant = iota
	// VariantRFC4122 is the standard variant specified in RFC4122
	VariantRFC4122
	// VariantMicrosoft is reserved for Microsoft Corporation backward compatibility
	VariantMicrosoft
)

// Variant returns the variant version and the data bits of
// M in the UUID format
//     xxxxxxxx-xxxx-Mxxx-Nxxx-xxxxxxxxxxxx
//
// The variant is designated by up to 4 bits in N
// - v0. 0xxx : N = 0...7 ; returns VariantNCS
// - v1. 10xx : N = 8...b ; returns VariantRFC4122
// - v2. 110x : N = c...d ; returns VariantMicrosoft
// - any other pattern is unknown and returns UnknownVariant
func (u UUID) Variant() Variant {
	switch {
	case (u[8] >> 7) == 0x00:
		return VariantNCS
	case (u[8] >> 6) == 0x02:
		return VariantRFC4122
	case (u[8] >> 5) == 0x06:
		return VariantMicrosoft
	case (u[8] >> 5) == 0x07:
		fallthrough
	default:
		return UnknownVariant
	}
}

// SetVariant sets the variant version and the data bits
// of N in the UUID format:
//     xxxxxxxx-xxxx-Mxxx-Nxxx-xxxxxxxxxxxx
func (u *UUID) SetVariant(v Variant) {
	switch v {
	case VariantNCS:
		u[8] = u[8]&(0xff>>1) | (0x00 << 7)
	case VariantRFC4122:
		u[8] = u[8]&(0xff>>2) | (0x02 << 6)
	case VariantMicrosoft:
		u[8] = u[8]&(0xff>>3) | (0x06 << 5)
	case UnknownVariant:
		fallthrough
	default:
		u[8] = u[8]&(0xff>>3) | (0x07 << 5)
	}
}

// Time extracts the timestamp from a Version1 or Version2 UUID
//
// For versions that are not RFC4122 complaint, or not Version1 or Version2, this will a zero time: `time.Time{}`
func (u UUID) Time() time.Time {
	if u.Variant() != VariantRFC4122 { // Not RFC4122 Compliant
		return time.Time{}
	}
	ver := u.Version()
	if ver < 1 || ver > 2 {
		return time.Time{} // Not a Time-based UUID
	}

	// Timestamp
	var ts [8]byte
	if ver == 1 {
		copy(ts[4:8], u[0:4]) // low
	}
	copy(ts[2:4], u[4:6]) // mid
	copy(ts[0:2], u[6:])  // high
	ts[0] &= 0x0F         // clear version
	ticks := binary.BigEndian.Uint64(ts[:])
	fmt.Printf("%x\n", ticks)
	// convert to unix nanos
	return time.Unix(0, int64((ticks-uuidEpocStart)*100))
}

// In its canonical textual representation, the sixteen octets of a UUID are represented as 32 hexadecimal (base 16) digits, displayed in five groups separated by hyphens, in the form 8-4-4-4-12 for a total of 36 characters (32 alphanumeric characters and four hyphens). For example:
//     123e4567-e89b-12d3-a456-426655440000
//     xxxxxxxx-xxxx-Mxxx-Nxxx-xxxxxxxxxxxx
// ...
// The four bits of digit M indicate the UUID version,
// and the one to three most significant bits of digit N indicate the UUID
// variant. In the example, M is 1 and N is a (10xx),
// meaning that the UUID is a variant 1, version 1 UUID;
// that is, a time-based DCE/RFC 4122 UUID.
// ...
// The canonical 8-4-4-4-12 format string is based on the "record layout"
// for the 16 bytes of the UUID
func (u UUID) String() string {
	return fmt.Sprintf("%s-%s-%s-%s-%s",
		hex.EncodeToString(u[0:4]),
		hex.EncodeToString(u[4:6]),
		hex.EncodeToString(u[6:8]),
		hex.EncodeToString(u[8:10]),
		hex.EncodeToString(u[10:]),
	)
}

// Format formats the UUID according to the fmt.Formatter interface.
//
//    %s    canonical form lowercase (123e4567-e89b-12d3-a456-426655440000)
//    %+s   canonical form UPPERCASE (123E4567-E89B-12D3-A456-426655440000)
//    %x    hash-like lowercase (123e4567e89b12d3a456426655440000)
//    %X    hash-like UPPERCASE (123E4567E89B12D3A456426655440000)
//    %v    equivalent to %s
//    %+v   equivalent to %+s
//    %q    equivalent to %s enclosed in double-quotes
//    %+q   equivalent to %+s enclosed in double-quotes
func (u UUID) Format(s fmt.State, verb rune) {
	switch verb {
	case 's':
		switch {
		case s.Flag('+'):
			_, _ = fmt.Fprintf(s, strings.ToUpper(u.String()))
		default:
			_, _ = fmt.Fprintf(s, u.String())
		}
	case 'x':
		_, _ = fmt.Fprintf(s, hex.EncodeToString(u[:]))
	case 'X':
		_, _ = fmt.Fprintf(s, strings.ToUpper(hex.EncodeToString(u[:])))
	case 'v':
		u.Format(s, 's')
	case 'q':
		_, _ = fmt.Fprintf(s, `"`)
		u.Format(s, 's')
		_, _ = fmt.Fprintf(s, `"`)
	}
}

// GenerateV1 creates a Time-based UUID.
//
// A Version 1 UUID is arranged like:
//
//    0                   1                   2                   3
//     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                          time_low                             |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |       time_mid                |         time_hi_and_version   |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |clk_seq_hi_res |  clk_seq_low  |         node (0-1)            |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                         node (2-5)                            |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//
func GenerateV1() (UUID, error) {
	return rfc4122.newV1()
}

// GenerateV2 generates a version 2 time-based UUID with DCE Security Information
//
// *Warning*: You should not exceed a call-rate of about 1 per 7 seconds. Why?
//
// The clock value truncated to the 28 most significant bits, compared to 60 bits in version 1
// Therefore, it will "tick" only once every 429.49 seconds: a little more than 7 minutes
// Additionally, the clock sequence number that prevents duplicate ids for the same timestamp is only 6 bits
// compared to 14 bits in version 1; so you can only call this 64 times in a 7 minute period (6.7 seconds per UUID)
//
// A Version 2 UUID is arranged like:
//    0                   1                   2                   3
//     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                              id                               |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |       time_mid                |         time_hi_and_version   |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |  clk_seq_res  |     domain    |         node (0-1)            |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                         node (2-5)                            |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//
// See: https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_2_(date-time_and_MAC_address,_DCE_security_version)
// RFC4122 does not formally specify Version 2, but references it from DCE 1.1 Authentication and Security Services (http://pubs.opengroup.org/onlinepubs/9696989899/chap5.htm#tagcjh_08_02_01_01),
// and mentions that *"Nothing in this document should be construed to override the DCE standards that defined UUIDs."*
func GenerateV2(domain byte, id uint32) (UUID, error) {
	u, err := dceSec.newV1()
	if err != nil {
		return u, err
	}
	u.SetDCESecurity(domain, id)
	return u, nil
}

// GenerateV3 generates a UUID by hashing the `ns` UUID and the input byte slice using MD5.
func GenerateV3(ns UUID, n []byte) UUID {
	h := md5.New()
	h.Write(ns[:])
	h.Write(n)
	hs := h.Sum(nil)
	var u UUID
	copy(u[:], hs[:16])
	u.SetVersion(3)
	u.SetVariant(VariantRFC4122)
	return u
}

// GenerateV4 generates a random (Version 4) UUID
// A version 4 UUID is randomly generated by grabbing a random 16-byte sequence from `crypto/rand`.
// An error may be returned if it fails to get the random bytes.
func GenerateV4() (UUID, error) {
	var u UUID
	_, err := rand.Read(u[:])
	if err != nil {
		return Nil, err
	}
	u.SetVersion(4)
	u.SetVariant(VariantRFC4122)
	return u, nil
}

// GenerateV5 generates a UUID by hashing the namespace UUID and the input byte slice using SHA1 (truncated to 16 bytes)
func GenerateV5(ns UUID, n []byte) UUID {
	h := sha1.New()
	h.Write(ns[:])
	h.Write(n)
	hs := h.Sum(nil)
	var u UUID
	copy(u[:], hs[:16])
	u.SetVersion(5)
	u.SetVariant(VariantRFC4122)
	return u
}

// GenerateV4String generates a UUID and serializes it to a string
// in the standard format:
// xxxxxxxx-xxxx-Mxxx-Nxxx-xxxxxxxxxxxx
func GenerateV4String() (string, error) {
	u, err := GenerateV4()
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

// UnmarshalText implements encoding.UnmarshalText
// The text can be in a few formats
// - hex only  : e8c8cec324e9445aa086f021ecbac4dd
// - canonical : e8c8cec3-24e9-445a-a086-f021ecbac4dd
// - braced    : {e8c8cec3-24e9-445a-a086-f021ecbac4dd}
// - urn:uuid  :
//   - urn:uuid:e8c8cec324e9445aa086f021ecbac4dd
//   - urn:uuid:e8c8cec3-24e9-445a-a086-f021ecbac4dd
func (u *UUID) UnmarshalText(text []byte) error {
	switch len(text) {
	case 32: // e8c8cec324e9445aa086f021ecbac4dd
		return u.unmarshalHex(text)
	case 36: // e8c8cec3-24e9-445a-a086-f021ecbac4dd
		return u.unmarshalCanonical(text)
	case 38: // {e8c8cec3-24e9-445a-a086-f021ecbac4dd}
		return u.unmarshalBraced(text)
	case 41: // urn:uuid:e8c8cec324e9445aa086f021ecbac4dd
		fallthrough
	case 45: // urn:uuid:e8c8cec3-24e9-445a-a086-f021ecbac4dd
		return u.unmarshalURN(text)
	}
	return ErrInvalidFormat
}

// MarshalText marshalls this UUID's value as text
func (u *UUID) MarshalText() ([]byte, error) {
	if u == nil {
		return nil, nil
	}
	return []byte(u.String()), nil
}

// Equals compares two pointers to a UUID, with the additional logic for equating `nil` to `uuid.Nil`
// For non-pointer comparison, a UUID can be directly compared using `==` since it is of type `[16]byte`
func (u *UUID) Equals(o *UUID) bool {
	if u == nil {
		return Nil.Equals(o)
	}
	if o == nil {
		return u.Equals(&Nil)
	}
	return *o == *u
}

var (
	urnPrefix  = []byte("urn:uuid:")
	byteGroups = []int{8, 4, 4, 4, 12}
)

func (u *UUID) unmarshalPlain(text []byte) error {
	switch len(text) {
	case 32:
		return u.unmarshalHex(text)
	case 36:
		return u.unmarshalCanonical(text)
	}
	return ErrInvalidFormat
}

func (u *UUID) unmarshalHex(text []byte) error {
	_, err := hex.Decode(u[:], text)
	return err
}

func (u *UUID) unmarshalCanonical(text []byte) error {
	i := 0
	j := 0
	for _, bgl := range byteGroups {
		if i > 0 {
			if text[i] != '-' {
				return ErrInvalidFormat
			}
			i++
		}
		ii := i + bgl
		jj := j + bgl/2
		_, err := hex.Decode(u[j:jj], text[i:ii])
		if err != nil {
			return ErrInvalidFormat
		}
		i += bgl
		j += bgl / 2
	}
	return nil
}

func (u *UUID) unmarshalBraced(text []byte) error {
	n := len(text) - 1
	if text[0] != '{' || text[n] != '}' {
		return ErrInvalidFormat
	}
	return u.unmarshalCanonical(text[1:n])
}

func (u *UUID) unmarshalURN(text []byte) error {
	if !bytes.HasPrefix(text, urnPrefix) {
		return ErrInvalidFormat
	}
	n := len(urnPrefix)
	return u.unmarshalPlain(text[n:])
}

// MustParseUUIDString parses an uuid string and returns the UUID
// If the parse fails, it panics
func MustParseUUIDString(s string) UUID {
	u, err := ParseUUIDString(s)
	if err != nil {
		panic(err)
	}
	return u
}

// ParseUUIDString parses a string into a UUID
func ParseUUIDString(s string) (UUID, error) {
	var u UUID
	err := u.UnmarshalText([]byte(s))
	return u, err
}
