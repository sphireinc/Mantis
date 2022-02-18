package byte

// BytesUnit represents a unit of magnitude for Bytes in either SI (base-10) or NIST (base-2)
type BytesUnit int64

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
