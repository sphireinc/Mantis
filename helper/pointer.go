package helper

// ToPtr returns a pointer value for the value passed in.
func ToPtr[T any](original T) *T {
	return &original
}

// FromPtr returns a pointer value for the value passed in.
func FromPtr[T any](original *T) T {
	return *original
}
