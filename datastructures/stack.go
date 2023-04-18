package datastructures

type StackNode struct {
	value int
	next  *StackNode
}

type Stack struct {
	top  *StackNode
	size int
}

func (s *Stack) Push(val int) {
	newNode := &StackNode{value: val}
	newNode.next = s.top
	s.top = newNode
	s.size++
}

func (s *Stack) Pop() (int, bool) {
	if s.top == nil {
		return 0, false
	}
	val := s.top.value
	s.top = s.top.next
	s.size--
	return val, true
}

func (s *Stack) Size() int {
	return s.size
}
