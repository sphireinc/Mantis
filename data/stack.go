package data

import "fmt"

// Stack holds our data
type Stack struct {
	values []any
}

// NewStack creates a new Stack
func NewStack() *Stack {
	return &Stack{}
}

// Insert a value into our stack
func (S *Stack) Insert(value any) {
	S.values = append(S.values, value)
}

// Pop a value from the top of our stack
func (S *Stack) Pop() any {
	i := len(S.values) - 1

	if i < 0 {
		return nil
	}

	value := S.values[i]
	S.values = append(S.values[:i], S.values[i+1:]...)

	return value
}

// Len returns the length of our stack
func (S *Stack) Len() int {
	return len(S.values)
}

// String converts our stack into a list
func (S *Stack) String() string {
	return fmt.Sprintf("%+v", S.values)
}
