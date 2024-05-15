package datastructures

// QueueNode defines a queue node
type QueueNode struct {
	value int
	next  *QueueNode
}

// Queue defines our queue structure
type Queue struct {
	head *QueueNode
	tail *QueueNode
	size int
}

// Enqueue adds an item onto our queue
func (q *Queue) Enqueue(val int) {
	newNode := &QueueNode{value: val}
	if q.tail == nil {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
	q.size++
}

// Dequeue removes an item from our queue
func (q *Queue) Dequeue() (int, bool) {
	if q.head == nil {
		return 0, false
	}
	val := q.head.value
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	q.size--
	return val, true
}

// Size returns our queue size
func (q *Queue) Size() int {
	return q.size
}
