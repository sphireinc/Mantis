package _generics

type boolSlice interface {
	type []bool
}

type numericSlice interface {
	type []int, []int8, []int16, []int32, []int64, []uint, []uint8, []uint16, []uint32, []uint64, []uintptr, []float32, []float64
}

type intsSlice interface {
	type []int, []int8, []int16, []int32, []int64, []int64, []uint, []uint8, []uint16, []uint32, []uint64
}

type signedIntsSlice interface {
	type []int, []int8, []int16, []int32, []int64, []int64
}

type unsignedIntsSlice interface {
	type []uint, []uint8, []uint16, []uint32, []uint64
}

type floatsSlice interface {
	type []float32, []float64
}

type complexesSlice interface {
	type []complex64, []complex128
}
