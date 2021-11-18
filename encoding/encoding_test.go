package encoding

import (
	"fmt"
	"testing"
)

func TestBase64EncodeStd(t *testing.T) {
	tests := []struct {
		rawStr     string
		encodedStr string
	}{
		{"abc123!?$*&()'-=@~", "YWJjMTIzIT8kKiYoKSctPUB+"},
		{"", ""},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			encodedStr := Base64EncodeStd(test.rawStr)
			if encodedStr != test.encodedStr {
				t.Fatalf("expected '%s', got '%s'", test.encodedStr, encodedStr)
			}
		})
	}
}

func TestBase64EncodeUrl(t *testing.T) {
	tests := []struct {
		rawStr     string
		encodedStr string
	}{
		{"abc123!?$*&()'-=@~", "YWJjMTIzIT8kKiYoKSctPUB-"},
		{"", ""},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			encodedStr := Base64EncodeUrl(test.rawStr)
			if encodedStr != test.encodedStr {
				t.Fatalf("expected '%s', got '%s'", test.encodedStr, encodedStr)
			}
		})
	}
}

func TestBase64Decode(t *testing.T) {
	tests := []struct {
		encodedStr string
		rawStr     string
	}{
		{"YWJjMTIzIT8kKiYoKSctPUB-", "abc123!?$*&()'-=@~"},
		{"YWJjMTIzIT8kKiYoKSctPUB+", "abc123!?$*&()'-=@~"},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			decodedStr, _ := Base64Decode(test.encodedStr)
			if test.rawStr != string(decodedStr) {
				t.Fatalf("expected '%s', got '%s'", test.rawStr, string(decodedStr))
			}
		})
	}
}
