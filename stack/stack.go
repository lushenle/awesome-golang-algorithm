package stack

// node storage data of the stack
type node struct {
	value interface{}
	prev  *node
}

// Stack storage of the stack
type Stack struct {
	top    *node
	length int
}

// New create a new stack
func New() *Stack {
	return &Stack{nil, 0}
}

// Len return the number of items in the stack
func (s *Stack) Len() int {
	return s.length
}

// Peek queries the top element of the stack
func (s *Stack) Peek() interface{} {
	return s.length
}

// Push adds new item to top of existing / empty stack
func (s *Stack) Push(value interface{}) {
	n := &node{value: value, prev: s.top}
	s.top = n
	s.length++
}

// Pop removes most recent item(top) from stack
func (s *Stack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}

	n := s.top
	s.top = n.prev
	return n.value
}

// IsEmpty returns whether the stack is empty or not
func (s *Stack) IsEmpty() bool {
	return s.length == 0
}
