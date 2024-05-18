package datastructures

import (
	"testing"
)

func TestNewHeap(t *testing.T) {
	h := NewHeap()
	if h == nil || len(h.data) != 0 {
		t.Errorf("NewHeap failed, expected a new heap with empty data, got %v", h.data)
	}
}

func TestHeapInsert(t *testing.T) {
	h := NewHeap()
	h.Insert(10)
	h.Insert(5)
	h.Insert(15)
	h.Insert(3)

	// After inserting 10, 5, 15, 3 into a min-heap, the smallest should be the root
	if h.data[0] != 3 {
		t.Errorf("HeapInsert failed, expected root to be 3, got %d", h.data[0])
	}
}

func TestHeapRemove(t *testing.T) {
	h := NewHeap()
	h.Insert(10)
	h.Insert(5)
	h.Insert(15)
	h.Insert(3)

	// Remove the root, which should be the smallest
	removed := h.Remove()
	if removed != 3 {
		t.Errorf("HeapRemove failed, expected to remove 3, got %d", removed)
	}

	// The new root should now be the next smallest
	if h.data[0] != 5 {
		t.Errorf("HeapRemove failed, expected new root to be 5, got %d", h.data[0])
	}
}

func TestHeapRemoveEmpty(t *testing.T) {
	h := NewHeap()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("HeapRemove on empty heap did not panic")
		}
	}()

	// This should panic as the heap is empty
	h.Remove()
}

func TestHeapOperationsSequence(t *testing.T) {
	h := NewHeap()
	elements := []int{25, 17, 33, 5, 21, 19, 27, 3, 100}
	expected := []int{3, 5, 17, 19, 21, 25, 27, 33, 100}

	for _, elem := range elements {
		h.Insert(elem)
	}

	for i, expectedVal := range expected {
		removed := h.Remove()
		if removed != expectedVal {
			t.Errorf("TestHeapOperationsSequence failed at %d, expected %d, got %d", i, expectedVal, removed)
		}
	}
}
