package stack

import (
	"sync"

	"github.com/lushenle/awesome-golang-algorithm/generic"
)

// Item the type of the stack
type Item generic.Type

// ItemStack the stack of Items
type ItemStack struct {
	items []Item
	lock  sync.RWMutex
}

// New create a new ItemStack
func (s *ItemStack) New() *ItemStack {
	s.items = []Item{}
	return s
}

// Push add an Item to the top of ItemStack
func (s *ItemStack) Push(it Item) {
	s.lock.Lock()
	s.items = append(s.items, it)
	s.lock.Unlock()
}

// Pop remove am Item from the top of the stack
func (s *ItemStack) Pop() *Item {
	s.lock.Lock()
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	s.lock.Unlock()
	return &item
}

// Peek return the top Item of the stack
func (s *ItemStack) Peek() *Item {
	s.lock.Lock()
	item := s.items[len(s.items)-1]
	s.lock.Unlock()
	return &item
}

// Len return the number of items in the stack
func (s *ItemStack) Len() int {
	return len(s.items)
}

// IsEmpty return true if the stack is empty
func (s *ItemStack) IsEmpty() bool {
	return len(s.items) == 0
}
