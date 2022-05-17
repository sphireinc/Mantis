package helper

type IntFloat interface {
	float32 | float64 | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

func Average[T IntFloat](args ...T) float64 {
	sum := 0.0
	count := 0.0
	for _, k := range args {
		sum += float64(k)
		count += 1
	}
	return sum / count
}
