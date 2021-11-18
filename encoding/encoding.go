package encoding

import (
	"encoding/base64"
	"errors"
	"strings"
)

// Base64EncodeStd encodes a string into a standard Base64 string
func Base64EncodeStd(data string) string {
	dataToEncode := []byte(data)
	return base64.StdEncoding.EncodeToString(dataToEncode)
}

// Base64EncodeUrl encodes a string into a Base64 string that is url-compatible
func Base64EncodeUrl(data string) string {
	dataToEncode := []byte(data)
	return base64.URLEncoding.EncodeToString(dataToEncode)
}

// Base64Decode decodes a base64 encoded string
func Base64Decode(encodedData string) ([]byte, error) {
	// + and / for toBase64 && - and _ for toBase64URL
	minusSign := strings.Index(encodedData, "-")
	underscore := strings.Index(encodedData, "_")

	if minusSign >= 0 || underscore >= 0 {
		return base64.URLEncoding.DecodeString(encodedData)
	}

	decoded, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		return nil, errors.New("invalid base64 encoded string")
	}
	return decoded, err
}
