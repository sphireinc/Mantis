package helper

// IntFloat is an interface for numeric types
type IntFloat interface {
	float32 | float64 | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

// Average calculates the average value of a set of numbers
func Average[T IntFloat](args ...T) float64 {
	sum := 0.0
	count := 0.0
	for _, k := range args {
		sum += float64(k)
		count++
	}
	return sum / count
}
