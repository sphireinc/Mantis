package datastructures

import (
	"testing"
)

func TestEnqueue(t *testing.T) {
	q := &Queue{}

	// Enqueue a single item
	q.Enqueue(10)
	if q.size != 1 || q.head.value != 10 || q.tail.value != 10 {
		t.Errorf("Enqueue failed, expected size 1, head 10, tail 10, got size %d, head %d, tail %d", q.size, q.head.value, q.tail.value)
	}

	// Enqueue multiple items
	q.Enqueue(20)
	q.Enqueue(30)
	if q.size != 3 || q.head.value != 10 || q.tail.value != 30 {
		t.Errorf("Enqueue failed, expected size 3, head 10, tail 30, got size %d, head %d, tail %d", q.size, q.head.value, q.tail.value)
	}
}

func TestDequeue(t *testing.T) {
	q := &Queue{}
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	// Dequeue a single item
	val, ok := q.Dequeue()
	if !ok || val != 10 || q.size != 2 || q.head.value != 20 {
		t.Errorf("Dequeue failed, expected value 10, size 2, head 20, got value %d, size %d, head %d", val, q.size, q.head.value)
	}

	// Dequeue remaining items
	val, ok = q.Dequeue()
	if !ok || val != 20 || q.size != 1 || q.head.value != 30 {
		t.Errorf("Dequeue failed, expected value 20, size 1, head 30, got value %d, size %d, head %d", val, q.size, q.head.value)
	}

	val, ok = q.Dequeue()
	if !ok || val != 30 || q.size != 0 || q.head != nil || q.tail != nil {
		t.Errorf("Dequeue failed, expected value 30, size 0, head nil, tail nil, got value %d, size %d, head %v, tail %v", val, q.size, q.head, q.tail)
	}

	// Dequeue from an empty queue
	val, ok = q.Dequeue()
	if ok || val != 0 {
		t.Errorf("Dequeue from empty queue failed, expected false, got value %d, ok %v", val, ok)
	}
}

func TestQSize(t *testing.T) {
	q := &Queue{}
	if q.Size() != 0 {
		t.Errorf("Size failed, expected 0, got %d", q.Size())
	}

	q.Enqueue(10)
	q.Enqueue(20)
	if q.Size() != 2 {
		t.Errorf("Size failed, expected 2, got %d", q.Size())
	}

	q.Dequeue()
	if q.Size() != 1 {
		t.Errorf("Size failed, expected 1, got %d", q.Size())
	}

	q.Dequeue()
	if q.Size() != 0 {
		t.Errorf("Size failed, expected 0, got %d", q.Size())
	}
}

func TestQueueIntegration(t *testing.T) {
	q := &Queue{}

	// Enqueue and Dequeue in sequence
	q.Enqueue(1)
	q.Enqueue(2)
	val, ok := q.Dequeue()
	if !ok || val != 1 {
		t.Errorf("Integration test failed, expected value 1, got %d", val)
	}

	q.Enqueue(3)
	val, ok = q.Dequeue()
	if !ok || val != 2 {
		t.Errorf("Integration test failed, expected value 2, got %d", val)
	}

	val, ok = q.Dequeue()
	if !ok || val != 3 {
		t.Errorf("Integration test failed, expected value 3, got %d", val)
	}
}
