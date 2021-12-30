package data

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_NewLinkedList(t *testing.T) {
	l := NewLinkedList()
	l.Insert("first")
	l.Insert("second")
	l.Insert("third")
	l.Insert("fourth")
	assert.Equal(t, l.String(""), "fourth -> third -> second -> first")
}

func TestList_Pop(t *testing.T) {
	l := NewLinkedList()
	l.Insert("first")
	l.Insert("second")
	l.Insert("third")
	l.Insert("fourth")
	assert.Equal(t, l.String(""), "fourth -> third -> second -> first")
	l.Pop()
	assert.Equal(t, l.String(""), "fourth -> third -> second")
	l.Pop()
	assert.Equal(t, l.String(""), "fourth -> third")
	l.Pop()
	assert.Equal(t, l.String(""), "fourth")
}

func TestList_Reverse(t *testing.T) {
	l := NewLinkedList()
	l.Insert("first")
	l.Insert("second")
	l.Insert("third")
	l.Insert("fourth")
	assert.Equal(t, l.String(""), "fourth -> third -> second -> first")
	l.Reverse()
	assert.Equal(t, l.String(""), "first -> second -> third -> fourth")
}

func TestList_Print(t *testing.T) {
	s := NewLinkedList()
	s.Insert(3)
	s.Insert(5)
	s.Insert(7)
	s.Insert(9)
	s.Print()
}
