package data

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_NewStack(t *testing.T) {
	s := NewStack()
	s.Insert(3)
	s.Insert(5)
	s.Insert(7)
	s.Insert(9)
	assert.Equal(t, s.Len(), 4)

	s = NewStack()
	data := make(map[string]string)
	data["hello"] = "world"
	s.Insert(data)
	s.Insert(data)
	s.Insert(data)
	s.Insert(data)
	assert.Equal(t, s.Len(), 4)

	s = NewStack()
	dataF := func() {}
	s.Insert(dataF)
	s.Insert(dataF)
	s.Insert(dataF)
	s.Insert(dataF)
	assert.Equal(t, s.Len(), 4)
}

func TestStack_Pop(t *testing.T) {
	s := NewStack()
	s.Insert(3)
	s.Insert(5)
	s.Insert(7)
	s.Insert(9)
	assert.Equal(t, s.Len(), 4)
	s.Pop()
	assert.Equal(t, s.Len(), 3)
	s.Pop()
	assert.Equal(t, s.Len(), 2)
	s.Pop()
	assert.Equal(t, s.Len(), 1)
	s.Pop()
	assert.Equal(t, s.Len(), 0)
	s.Pop()
	assert.Equal(t, s.Len(), 0)
}

func TestStack_Print(t *testing.T) {
	s := NewStack()
	s.Insert(3)
	s.Insert(5)
	s.Insert(7)
	s.Insert(9)
	s.Print()
}
