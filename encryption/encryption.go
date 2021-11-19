package encryption

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"io"
	"strings"
	"time"
)

const (
	Md5 int = iota
	Sha224
	Sha256
	Sha384
	Sha512
	Sha512224
	Sha512256
	Hmac512
)

type Hash struct {
	Input     string
	Output    string
	Algorithm int
}

func (h *Hash) Hash() {
	switch h.Algorithm {
	case 0:
		h.md5()
	case 1:
		h.sha224()
	case 2:
		h.sha256()
	case 3:
		h.sha384()
	case 4:
		h.sha512()
	case 5:
		h.sha512224()
	case 6:
		h.sha512256()
	case 7:
		h.hmac512()
	}
}

// md5 returns an MD5 string given an input string
func (h *Hash) md5() {
	hash := md5.New()
	_, _ = io.WriteString(hash, h.Input)
	h.Output = hex.EncodeToString(hash.Sum(nil))
}

// sha224 returns an SHA224 string given an input string
func (h *Hash) sha224() {
	hash := sha256.New224()
	hash.Write([]byte(h.Input))
	h.Output = hex.EncodeToString(hash.Sum(nil))
}

// sha256 returns an SHA256 string given an input string
func (h *Hash) sha256() {
	hash := sha256.New()
	hash.Write([]byte(h.Input))
	h.Output = hex.EncodeToString(hash.Sum(nil))
}

// sha384 returns an SHA384 string given an input string
func (h *Hash) sha384() {
	hash := sha512.New384()
	hash.Write([]byte(h.Input))
	h.Output = hex.EncodeToString(hash.Sum(nil))
}

// sha512 returns an SHA-512 string given an input string
func (h *Hash) sha512() {
	hash := sha512.New()
	hash.Write([]byte(h.Input))
	h.Output = hex.EncodeToString(hash.Sum(nil))
}

// sha512224 returns an SHA512-224 string given an input string
func (h *Hash) sha512224() {
	hash := sha512.New512_224()
	hash.Write([]byte(h.Input))
	h.Output = hex.EncodeToString(hash.Sum(nil))
}

// sha512256 returns an SHA512-256 string given an input string
func (h *Hash) sha512256() {
	hash := sha512.New512_256()
	hash.Write([]byte(h.Input))
	h.Output = hex.EncodeToString(hash.Sum(nil))
}

// hmac512 returns an HMAC-512 string given an input string
func (h *Hash) hmac512() {
	hmac512 := hmac.New(sha512.New, []byte(h.Input))
	hmac512.Write([]byte(h.Input))
	h.Output = base64.StdEncoding.EncodeToString(hmac512.Sum(nil))
}

// CreateRandomString generates a random string of n bytes
func CreateRandomString(bytes int) string {
	return hex.EncodeToString(CreateRandomBytes(bytes))
}

func CreateRandomBytes(bytes int) []byte {
	if bytes == 0 {
		bytes = 16 // default to 16 bytes
	}

	randomBytes := make([]byte, bytes)
	_, _ = rand.Read(randomBytes[:])

	nowTime := strings.Replace(time.Now().String(), " ", "_", -1)
	return append([]byte(nowTime), randomBytes...)
}
