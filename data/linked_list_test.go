package data

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func newLinkedListHelper() *list {
	l := NewLinkedList()
	l.Insert("first key")
	l.Insert("second key")
	l.Insert("third key")
	l.Insert("fourth key")
	return l
}

func TestNewLinkedList(t *testing.T) {
	l := newLinkedListHelper()
	assert.Equal(t, l.String(""), "fourth key -> third key -> second key -> first key")
	l.Pop()
	assert.Equal(t, l.String(""), "fourth key -> third key -> second key")
	l.Pop()
	assert.Equal(t, l.String(""), "fourth key -> third key")
	l.Pop()
	assert.Equal(t, l.String(""), "fourth key")
}

func TestReverseLinkedList(t *testing.T) {
	l := newLinkedListHelper()
	assert.Equal(t, l.String(""), "fourth key -> third key -> second key -> first key")
	l.Reverse()
	assert.Equal(t, l.String(""), "first key -> second key -> third key -> fourth key")
}
