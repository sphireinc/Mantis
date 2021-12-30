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
	assert.Equal(t, "fourth -> third -> second -> first", l.String(""))
}

func TestList_Pop(t *testing.T) {
	l := NewLinkedList()
	l.Insert("first")
	l.Insert("second")
	l.Insert("third")
	l.Insert("fourth")
	assert.Equal(t, "fourth -> third -> second -> first", l.String(""))
	l.Pop()
	assert.Equal(t, "fourth -> third -> second", l.String(""))
	l.Pop()
	assert.Equal(t, "fourth -> third", l.String(""))
	l.Pop()
	assert.Equal(t, "fourth", l.String(""))
}

func TestList_Reverse(t *testing.T) {
	l := NewLinkedList()
	l.Insert("first")
	l.Insert("second")
	l.Insert("third")
	l.Insert("fourth")
	assert.Equal(t, "fourth -> third -> second -> first", l.String(""))
	l.Reverse()
	assert.Equal(t, "first -> second -> third -> fourth", l.String(""))
}

func TestList_Print(t *testing.T) {
	s := NewLinkedList()
	s.Insert(3)
	s.Insert(5)
	s.Insert(7)
	s.Insert(9)
	s.Print()
}
