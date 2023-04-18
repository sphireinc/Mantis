package datastructures

type QueueNode struct {
	value int
	next  *QueueNode
}

type Queue struct {
	head *QueueNode
	tail *QueueNode
	size int
}

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

func (q *Queue) Size() int {
	return q.size
}
