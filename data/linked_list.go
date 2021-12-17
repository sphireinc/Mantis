package data

import (
	"fmt"
)

type node struct {
	previous *node
	next     *node
	key      any
}

type list struct {
	head *node
	tail *node
}

func NewLinkedList() *list {
	return &list{}
}

func (L *list) Insert(key any) {
	newNode := &node{
		next: L.head,
		key:  key,
	}
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

func (L *list) Pop() any {
	tail := L.tail

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

func (L *list) Display() {
	var output string
	list := L.head
	for list != nil {
		output = output + fmt.Sprintf("%+v -> ", list.key)
		list = list.next
	}

	output = output[0 : len(output)-3]
	fmt.Println(output)
}

func (L *list) Reverse() {

}
