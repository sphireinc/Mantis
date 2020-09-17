package byte

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// ParseErr returned invalid strings are parsed
type ParseErr struct {
	Func string
	Str  string
	Err  error
}

func (e *ParseErr) Error() string {
	return "bytes." + e.Func + ": " + "parsing " + strconv.Quote(e.Str) + ": " + e.Err.Error()
}

// BytesUnit represents a unit of magnitude for Bytes in either SI (base-10) or NIST (base-2)
type BytesUnit int64

const (
	// B byte (Unit)
	B BytesUnit = 1
	// KB Kilobyte (SI) - 1000 bytes
	KB BytesUnit = 1000
	// MB Megabyte (SI) - 1e6 bytes
	MB BytesUnit = 1e6
	// GB Gigabyte (SI) - 1e9 bytes
	GB BytesUnit = 1e9
	// TB Terabyte (SI) - 1e12 bytes
	TB BytesUnit = 1e12
	// PB Petabyte (SI) - 1e15 bytes
	PB BytesUnit = 1e15

	// KiB Kibibyte (NIST) - 1024 bytes
	KiB BytesUnit = 1024
	// MiB Mebibyte (NIST) - 2^20 bytes
	MiB = 1024 * KiB
	// GiB Gibibyte (NIST) - 2^30 bytes
	GiB = 1024 * MiB
	// TiB Tebibyte (NIST) - 2^40 bytes
	TiB = 1024 * GiB
	// PiB Pebibyte (NIST) - 2^50 bytes
	PiB = 1024 * TiB
)

var strToByteUnit = map[string]BytesUnit{
	"":    B,
	"B":   B,
	"KB":  KB,
	"KIB": KiB,
	"MB":  MB,
	"MIB": MiB,
	"GB":  GB,
	"GIB": GiB,
	"TB":  TB,
	"TIB": TiB,
	"PB":  PB,
	"PIB": PiB,
}

var byteUnitToStr = map[BytesUnit]string{
	B:   "B",
	KB:  "KB",
	KiB: "KiB",
	MB:  "MB",
	MiB: "MiB",
	GB:  "GB",
	GiB: "GiB",
	TB:  "TB",
	TiB: "TiB",
	PB:  "PB",
	PiB: "PiB",
}

var bytesUnitThresholds = []struct {
	unit BytesUnit
	less Bytes
}{
	{B, 0x400},
	{KiB, Bytes(uint64(MiB) >> 4)},
	{MiB, Bytes(uint64(GiB) >> 4)},
	{GiB, Bytes(uint64(TiB) >> 4)},
	{TiB, Bytes(uint64(PiB) >> 4)},
	{PiB, 0x7FFFFFFFFFFFFFFF},
}
var bytesSIUnitThresholds = []struct {
	unit BytesUnit
	less Bytes
}{
	{B, 5e2},
	{KB, 5e5},
	{MB, 5e8},
	{GB, 5e11},
	{TB, 5e14},
	{PB, 0x7FFFFFFFFFFFFFFF},
}

// Convert a byte count into the (possibly fractional) count in this unit
func (u BytesUnit) Convert(b Bytes) float64 {
	if u == B {
		return float64(b)
	}
	return float64(b) / float64(u)
}

func (u BytesUnit) String() string {
	return byteUnitToStr[u]
}

// Bytes is a formattable byte count
type Bytes int64

// Format bytes
func (b Bytes) Format(s fmt.State, verb rune) {
	switch verb {
	case 'd':
		_, _ = fmt.Fprintf(s, "%d B", int64(b))
	case 's':
		fallthrough
	case 'v':
		if b < 0 {
			_, _ = s.Write([]byte{'-'})
			b = b * -1
		}
		thresholds := bytesUnitThresholds
		if s.Flag('+') {
			thresholds = bytesSIUnitThresholds
		}
		for _, t := range thresholds {
			if b < t.less {
				f := t.unit.Convert(b)
				if prec, ok := s.Precision(); ok {
					format := fmt.Sprintf("%%.%df %%s", prec)
					_, _ = fmt.Fprintf(s, format, f, t.unit)
				} else {
					_, _ = fmt.Fprintf(s, "%s %s", strconv.FormatFloat(f, 'f', -1, 64), t.unit)
				}
				return
			}
		}
	}
}

func (b Bytes) String() string {
	return fmt.Sprintf("%s", b)
}

var parseBytesRegex = regexp.MustCompile(`(?i)(?P<value>[0-9.+eE+-]+)\s*(?P<unit>[a-z]?(?:i?b)?)?`)

// ParseBytes parses the given string with byte units into a Byte count
func ParseBytes(s string) (Bytes, error) {
	gs := parseBytesRegex.FindStringSubmatch(s)
	if len(gs) == 0 {
		return 0, &ParseErr{
			Func: "ParseBytes",
			Str:  s,
			Err:  errors.New("bad string format"),
		}
	}
	vs := strings.ToLower(gs[1])
	us := gs[2]

	unit, ok := strToByteUnit[strings.ToUpper(strings.TrimSpace(us))]
	if !ok {
		return 0, &ParseErr{
			Func: "ParseBytes",
			Str:  s,
			Err:  errors.New("unknown bytes unit"),
		}
	}

	vs = strings.Replace(vs, ",", "", -1)
	if strings.Contains(vs, ".") || strings.Contains(vs, "e") {
		f, err := strconv.ParseFloat(vs, 64)
		if err != nil {
			return 0, &ParseErr{
				Func: "ParseBytes",
				Str:  s,
				Err:  err,
			}
		}
		return Bytes(f * float64(unit)), nil
	}
	i, err := strconv.ParseInt(vs, 10, 64)
	if err != nil {
		return 0, &ParseErr{
			Func: "ParseBytes",
			Str:  s,
			Err:  err,
		}
	}
	return Bytes(i * int64(unit)), nil
}
