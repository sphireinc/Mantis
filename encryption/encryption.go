package encryption

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"hash"
	"strings"
	"time"
)

const (
	Md5 int8 = iota
	Sha224
	Sha256
	Sha384
	Sha512
	Sha512224
	Sha512256
	Hmac512
)

type mHash struct {
	input     string
	isHashed  bool
	Output    string
	algorithm int8
}

// MarshalJSON implements the JSON encoding interface
func (h *mHash) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"input":     h.input,
		"output":    h.Output,
		"algorithm": h.algorithm,
	})
}

// New returns an instance of mHash given our input and algorithm
func New(input string, algorithm int8) *mHash {
	return &mHash{
		input:     input,
		algorithm: algorithm,
	}
}

// IsHashed tells us whether our mHash has been mHash()'d
func (h *mHash) IsHashed() bool {
	return h.isHashed
}

// Algorithm returns the chosen algorithm as a string and int
func (h *mHash) Algorithm() (int8, string) {
	switch h.algorithm {
	case 0:
		return Md5, "md5"
	case 1:
		return Sha224, "SHA-224"
	case 2:
		return Sha256, "SHA-256"
	case 3:
		return Sha384, "SHA-384"
	case 4:
		return Sha512, "SHA-512"
	case 5:
		return Sha512224, "SHA-512/224"
	case 6:
		return Sha512256, "SHA-512/256"
	case 7:
		return Hmac512, "HMAC-512"
	}
	return -1, ""
}

// GetInput returns the initial input
func (h *mHash) GetInput() string {
	return h.input
}

// GetOutput returns the hashed output
func (h *mHash) GetOutput() string {
	return h.Output
}

// Hash performs our hash, fills in Output, and unsets input
func (h *mHash) Hash() {
	if h.algorithm == Hmac512 {
		hmac512 := hmac.New(sha512.New, []byte(h.input))
		hmac512.Write([]byte(h.input))
		h.Output = base64.StdEncoding.EncodeToString(hmac512.Sum(nil))
	} else {
		var hasher hash.Hash

		switch h.algorithm {
		case Md5:
			hasher = md5.New()
		case Sha224:
			hasher = sha256.New224()
		case Sha256:
			hasher = sha256.New()
		case Sha384:
			hasher = sha512.New384()
		case Sha512:
			hasher = sha512.New()
		case Sha512224:
			hasher = sha512.New512_224()
		case Sha512256:
			hasher = sha512.New512_256()
		}

		hasher.Write([]byte(h.input))
		h.isHashed = true
		h.input = ""
		h.Output = hex.EncodeToString(hasher.Sum(nil))
	}
}

// CreateRandomString generates a random string of n bytes
func CreateRandomString(bytes int) string {
	return hex.EncodeToString(CreateRandomBytes(bytes))
}

// CreateRandomBytes creates a random bytes of bytes int
func CreateRandomBytes(bytes int) []byte {
	if bytes == 0 {
		bytes = 16 // default to 16 bytes
	}

	randomBytes := make([]byte, bytes)
	_, _ = rand.Read(randomBytes[:])

	nowTime := strings.Replace(time.Now().String(), " ", "_", -1)
	return append([]byte(nowTime), randomBytes...)
}
