package datastructures

import (
	"testing"
)

func TestStack_Push(t *testing.T) {
	stack := &Stack{}

	// Push a single item
	stack.Push(10)
	if stack.size != 1 || stack.top.value != 10 {
		t.Errorf("Push failed, expected size 1, top 10, got size %d, top %d", stack.size, stack.top.value)
	}

	// Push multiple items
	stack.Push(20)
	stack.Push(30)
	if stack.size != 3 || stack.top.value != 30 {
		t.Errorf("Push failed, expected size 3, top 30, got size %d, top %d", stack.size, stack.top.value)
	}
}

func TestStack_Pop(t *testing.T) {
	stack := &Stack{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	// Pop a single item
	val, ok := stack.Pop()
	if !ok || val != 30 || stack.size != 2 {
		t.Errorf("Pop failed, expected value 30, size 2, got value %d, size %d", val, stack.size)
	}

	// Pop remaining items
	val, ok = stack.Pop()
	if !ok || val != 20 || stack.size != 1 {
		t.Errorf("Pop failed, expected value 20, size 1, got value %d, size %d", val, stack.size)
	}

	val, ok = stack.Pop()
	if !ok || val != 10 || stack.size != 0 {
		t.Errorf("Pop failed, expected value 10, size 0, got value %d, size %d", val, stack.size)
	}

	// Pop from an empty stack
	val, ok = stack.Pop()
	if ok || val != 0 {
		t.Errorf("Pop from empty stack failed, expected false, got value %d, ok %v", val, ok)
	}
}

func TestStack_Size(t *testing.T) {
	stack := &Stack{}
	if stack.Size() != 0 {
		t.Errorf("Size failed, expected 0, got %d", stack.Size())
	}

	stack.Push(10)
	stack.Push(20)
	if stack.Size() != 2 {
		t.Errorf("Size failed, expected 2, got %d", stack.Size())
	}

	stack.Pop()
	if stack.Size() != 1 {
		t.Errorf("Size failed, expected 1, got %d", stack.Size())
	}

	stack.Pop()
	if stack.Size() != 0 {
		t.Errorf("Size failed, expected 0, got %d", stack.Size())
	}
}

func TestStack_IsEmpty(t *testing.T) {
	stack := &Stack{}
	if !stack.IsEmpty() {
		t.Errorf("IsEmpty failed, expected true, got false")
	}

	stack.Push(10)
	if stack.IsEmpty() {
		t.Errorf("IsEmpty failed, expected false, got true")
	}

	stack.Pop()
	if !stack.IsEmpty() {
		t.Errorf("IsEmpty failed, expected true, got false")
	}
}

func TestStack_Clear(t *testing.T) {
	stack := &Stack{}
	stack.Push(10)
	stack.Push(20)
	stack.Clear()

	if stack.size != 0 || stack.top != nil {
		t.Errorf("Clear failed, expected size 0, top nil, got size %d, top %v", stack.size, stack.top)
	}
}
