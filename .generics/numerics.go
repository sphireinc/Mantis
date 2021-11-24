package _generics

type numerics interface {
	type int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64
}

type ints interface {
	type int, int8, int16, int32, int64, int64, uint, uint8, uint16, uint32, uint64
}

type signedInts interface {
	type int, int8, int16, int32, int64, int64
}

type unsignedInts interface {
	type uint, uint8, uint16, uint32, uint64
}

type floats interface {
	type float32, float64
}

type complexes interface {
	type complex64, complex128
}