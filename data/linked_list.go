package data

import "fmt"

type linkedListNode struct {
	previous *linkedListNode
	next     *linkedListNode
	key      any
}

type list struct {
	head *linkedListNode
	tail *linkedListNode
}

// NewLinkedList returns an instance of list
func NewLinkedList() *list {
	return &list{}
}

// Insert a key and value into a linked list
func (L *list) Insert(key any) {
	// create our new node
	newNode := &linkedListNode{
		next: L.head,
		key:  key,
	}

	// if we have a head, make the new node the previous
	if L.head != nil {
		L.head.previous = newNode
	}
	L.head = newNode

	node := L.head
	for node.next != nil {
		node = node.next
	}
	L.tail = node
}

// Pop returns a key from a list (FIFO)
func (L *list) Pop() any {
	// copy the tail to return the copy
	tail := L.tail

	// Reset our list positions
	node := L.tail
	for node.previous != nil {
		node = node.previous
		if node.next.key == tail.key {
			node.next = nil
			L.tail = node
		}
	}

	return tail.key
}

// Print our list
func (L *list) Print() {
	fmt.Println(L.String(""))
}

// String returns our list as a string
func (L *list) String(delimiter string) string {
	var output string

	if delimiter == "" {
		delimiter = "->"
	}

	list := L.head
	for list != nil {
		output = output + fmt.Sprintf("%+v %s ", list.key, delimiter)
		list = list.next
	}

	return output[0 : len(output)-4]
}

func (L *list) Reverse() {
	var prev *linkedListNode
	current := L.head
	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}
	L.head = prev
}
