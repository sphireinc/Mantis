package data

import "testing"

func TestNewLinkedList(t *testing.T) {
	l := NewLinkedList()
	l.Insert("first key")
	l.Insert("second key")
	l.Insert("third key")
	l.Insert("fourth key")
	l.Display()
	l.Pop()
	l.Display()
	l.Pop()
	l.Display()
	l.Pop()
	l.Display()
}
