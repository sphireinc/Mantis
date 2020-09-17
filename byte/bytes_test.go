package byte

import (
	"fmt"
	"testing"
)

func TestBytesStringer(t *testing.T) {
	tests := []struct {
		bytes    Bytes
		expected string
	}{
		{Bytes(1), "1 B"},
		{Bytes(0x400), "1 KiB"},
		{Bytes(0x100000), "1 MiB"},
		{Bytes(0x40000000), "1 GiB"},
		{Bytes(0x10000000000), "1 TiB"},
		{Bytes(0x4000000000000), "1 PiB"},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			act := test.bytes.String()
			if act != test.expected {
				t.Fatalf("expected '%s', got '%s'", test.expected, act)
			}
		})
	}
}

func TestFormatBytes(t *testing.T) {
	tests := []struct {
		format   string
		bytes    Bytes
		expected string
	}{
		{"%v", Bytes(1), "1 B"},
		{"%v", Bytes(0x400), "1 KiB"},
		{"%v", Bytes(0x100000), "1 MiB"},
		{"%v", Bytes(0x40000000), "1 GiB"},
		{"%v", Bytes(0x10000000000), "1 TiB"},
		{"%v", Bytes(0x4000000000000), "1 PiB"},

		{"%v", Bytes(0x200), "512 B"},
		{"%v", Bytes(0x40000), "0.25 MiB"},
		{"%v", Bytes(0x20000000), "0.5 GiB"},
		{"%v", Bytes(0x9000000000), "0.5625 TiB"},
		{"%v", Bytes(0xAA00000000000), "2.65625 PiB"},

		{"%s", Bytes(0xAA00000000000), "2.65625 PiB"},
		{"%s", Bytes(-0xAA00000000000), "-2.65625 PiB"},
		{"%.3s", Bytes(-0xAA00000000000), "-2.656 PiB"},

		{"%+v", Bytes(1), "1 B"},
		{"%+v", Bytes(10), "10 B"},
		{"%+v", Bytes(1e2), "100 B"},
		{"%+v", Bytes(1e3), "1 KB"},
		{"%+v", Bytes(1e4), "10 KB"},
		{"%+v", Bytes(1e5), "100 KB"},
		{"%+v", Bytes(1e6), "1 MB"},
		{"%+v", Bytes(1e7), "10 MB"},
		{"%+v", Bytes(1e8), "100 MB"},
		{"%+v", Bytes(1e9), "1 GB"},
		{"%+v", Bytes(1e10), "10 GB"},
		{"%+v", Bytes(1e11), "100 GB"},
		{"%+v", Bytes(1e12), "1 TB"},
		{"%+v", Bytes(1e13), "10 TB"},
		{"%+v", Bytes(1e14), "100 TB"},
		{"%+v", Bytes(1e15), "1 PB"},
		{"%+v", Bytes(1e16), "10 PB"},
		{"%+v", Bytes(1e17), "100 PB"},

		{"%d", Bytes(1), "1 B"},
		{"%d", Bytes(10), "10 B"},
		{"%d", Bytes(1e2), "100 B"},
		{"%d", Bytes(1e3), "1000 B"},
		{"%d", Bytes(1e4), "10000 B"},
		{"%d", Bytes(1e5), "100000 B"},
		{"%d", Bytes(1e6), "1000000 B"},
		{"%d", Bytes(1e7), "10000000 B"},
		{"%d", Bytes(1e8), "100000000 B"},
		{"%d", Bytes(1e9), "1000000000 B"},
		{"%d", Bytes(1e10), "10000000000 B"},
		{"%d", Bytes(1e11), "100000000000 B"},
		{"%d", Bytes(1e12), "1000000000000 B"},
		{"%d", Bytes(1e13), "10000000000000 B"},
		{"%d", Bytes(1e14), "100000000000000 B"},
		{"%d", Bytes(1e15), "1000000000000000 B"},
		{"%d", Bytes(1e16), "10000000000000000 B"},
		{"%d", Bytes(1e17), "100000000000000000 B"},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			act := fmt.Sprintf(test.format, test.bytes)
			if act != test.expected {
				t.Fatalf("expected '%s', got '%s'", test.expected, act)
			}
		})
	}
}

func TestParseBytes(t *testing.T) {
	tests := []struct {
		expected Bytes
		parse    string
	}{
		{Bytes(0x200), "512 B"},
		{Bytes(0x40000), "0.25 MiB"},
		{Bytes(0x20000000), "0.5 GiB"},
		{Bytes(0x9000000000), "0.5625 TiB"},
		{Bytes(0xAA00000000000), "2.65625 PiB"},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			act, err := ParseBytes(test.parse)
			if err != nil {
				t.Fatal(err)
			}
			if act != test.expected {
				t.Fatalf("expected '%d', got '%d'", test.expected, act)
			}
		})
	}
}

func TestParseBytesError(t *testing.T) {
	tests := []struct {
		expected string
		parse    string
	}{
		{"", ""},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if bytes, err := ParseBytes(test.parse); err != nil && bytes == 1 {
				t.Fatalf("expected error parsing '%s'", test.parse)
			}
		})
	}
}
