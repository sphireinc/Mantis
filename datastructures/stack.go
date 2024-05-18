package datastructures

// StackNode defines our stack node
type StackNode struct {
	value int
	next  *StackNode
}

// Stack defines our stack
type Stack struct {
	top  *StackNode
	size int
}

// Push a value onto our stack
func (s *Stack) Push(val int) {
	newNode := &StackNode{value: val}
	newNode.next = s.top
	s.top = newNode
	s.size++
}

// Pop the top-most item from our stack
func (s *Stack) Pop() (int, bool) {
	if s.top == nil {
		return 0, false
	}
	val := s.top.value
	s.top = s.top.next
	s.size--
	return val, true
}

// Size returns the size of our stack
func (s *Stack) Size() int {
	return s.size
}

// IsEmpty returns true if the stack is empty
func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

// Clear the stack
func (s *Stack) Clear() {
	s.top = nil
	s.size = 0
}
