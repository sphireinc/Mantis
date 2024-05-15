package datastructures

import (
	"fmt"
)

// LinkedListNode defines our LL node
type LinkedListNode struct {
	value int
	next  *LinkedListNode
}

// LinkedList defines our parent structure
type LinkedList struct {
	head *LinkedListNode
	size int
}

// Insert a new value into our LL
func (l *LinkedList) Insert(val int) {
	newNode := &LinkedListNode{value: val}
	if l.head == nil {
		l.head = newNode
	} else {
		currNode := l.head
		for currNode.next != nil {
			currNode = currNode.next
		}
		currNode.next = newNode
	}
	l.size++
}

// Remove a value from our LL
func (l *LinkedList) Remove(val int) bool {
	if l.head == nil {
		return false
	}
	if l.head.value == val {
		l.head = l.head.next
		l.size--
		return true
	}
	currNode := l.head
	for currNode.next != nil {
		if currNode.next.value == val {
			currNode.next = currNode.next.next
			l.size--
			return true
		}
		currNode = currNode.next
	}
	return false
}

// Size returns the size of our LL
func (l *LinkedList) Size() int {
	return l.size
}

// String converts our LL into a string
func (l *LinkedList) String() string {
	str := ""
	currNode := l.head
	for currNode != nil {
		str += fmt.Sprintf("%d -> ", currNode.value)
		currNode = currNode.next
	}
	str += "nil"
	return str
}
