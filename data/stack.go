package data

import "fmt"

type stack struct {
	values []any
}

func NewStack() *stack {
	return &stack{}
}

func (S *stack) Insert(value any) {
	S.values = append(S.values, value)
}

func (S *stack) Pop() any {
	i := len(S.values) - 1

	if i < 0 {
		return nil
	}

	value := S.values[i]
	S.values = append(S.values[:i], S.values[i+1:]...)

	return value
}

func (S *stack) Len() int {
	return len(S.values)
}

func (S *stack) String() string {
	return fmt.Sprintf("%+v", S.values)
}
