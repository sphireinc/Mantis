package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack_NewStack(t *testing.T) {
	s := NewStack()
	s.Insert(3)
	s.Insert(5)
	s.Insert(7)
	s.Insert(9)
	assert.Equal(t, 4, s.Len())

	s = NewStack()
	data := make(map[string]string)
	data["hello"] = "world"
	s.Insert(data)
	s.Insert(data)
	s.Insert(data)
	s.Insert(data)
	assert.Equal(t, 4, s.Len())

	s = NewStack()
	dataF := func() {}
	s.Insert(dataF)
	s.Insert(dataF)
	s.Insert(dataF)
	s.Insert(dataF)
	assert.Equal(t, 4, s.Len())
}

func TestStack_Pop(t *testing.T) {
	s := NewStack()
	s.Insert(3)
	s.Insert(5)
	s.Insert(7)
	s.Insert(9)
	assert.Equal(t, 4, s.Len())
	s.Pop()
	assert.Equal(t, 3, s.Len())
	s.Pop()
	assert.Equal(t, 2, s.Len())
	s.Pop()
	assert.Equal(t, 1, s.Len())
	s.Pop()
	assert.Equal(t, 0, s.Len())
	s.Pop()
	assert.Equal(t, 0, s.Len())
}

func TestStack_String(t *testing.T) {
	s := NewStack()
	s.Insert(3)
	s.Insert(5)
	s.Insert(7)
	s.Insert(9)
	assert.Equal(t, s.String(), "[3 5 7 9]")
}
