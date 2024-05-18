package datastructures

import (
	"testing"
)

func TestInsert(t *testing.T) {
	ll := &LinkedList{}
	values := []int{10, 20, 30}

	// Insert values into the linked list
	for _, val := range values {
		ll.Insert(val)
	}

	// Check the size of the linked list
	if ll.Size() != len(values) {
		t.Errorf("Insert failed, expected size %d, got %d", len(values), ll.Size())
	}

	// Check the string representation of the linked list
	expectedStr := "10 -> 20 -> 30 -> nil"
	if ll.String() != expectedStr {
		t.Errorf("Insert failed, expected string %s, got %s", expectedStr, ll.String())
	}
}

func TestLLRemove(t *testing.T) {
	ll := &LinkedList{}
	values := []int{10, 20, 30}
	for _, val := range values {
		ll.Insert(val)
	}

	// Remove a middle element
	ll.Remove(20)
	expectedStr := "10 -> 30 -> nil"
	if ll.String() != expectedStr {
		t.Errorf("Remove failed, expected string %s, got %s", expectedStr, ll.String())
	}

	// Remove a non-existent element
	if ll.Remove(40) {
		t.Errorf("Remove failed, should not successfully remove non-existent element")
	}

	// Remove the head
	ll.Remove(10)
	expectedStr = "30 -> nil"
	if ll.String() != expectedStr {
		t.Errorf("Remove head failed, expected string %s, got %s", expectedStr, ll.String())
	}
}

func TestReverse(t *testing.T) {
	ll := &LinkedList{}
	values := []int{10, 20, 30}
	for _, val := range values {
		ll.Insert(val)
	}

	// Reverse the linked list
	ll.Reverse()
	expectedStr := "30 -> 20 -> 10 -> nil"
	if ll.String() != expectedStr {
		t.Errorf("Reverse failed, expected string %s, got %s", expectedStr, ll.String())
	}
}

func TestLLSize(t *testing.T) {
	ll := &LinkedList{}
	if ll.Size() != 0 {
		t.Errorf("Size failed, expected 0, got %d", ll.Size())
	}
	ll.Insert(10)
	ll.Insert(20)
	if ll.Size() != 2 {
		t.Errorf("Size failed, expected 2, got %d", ll.Size())
	}
	ll.Remove(10)
	if ll.Size() != 1 {
		t.Errorf("Size failed, expected 1, got %d", ll.Size())
	}
}

func TestString(t *testing.T) {
	ll := &LinkedList{}
	values := []int{1, 2, 3}
	for _, val := range values {
		ll.Insert(val)
	}
	expectedStr := "1 -> 2 -> 3 -> nil"
	if ll.String() != expectedStr {
		t.Errorf("String failed, expected %s, got %s", expectedStr, ll.String())
	}
}
