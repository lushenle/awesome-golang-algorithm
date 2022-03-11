package set

import "github.com/lushenle/awesome-golang-algorithm/generic"

// Item the type of the set
type Item generic.Type

// ItemSet the set of the Items
type ItemSet struct{ items map[Item]bool }

// Add adds a new element to the set,
// return a pointer to the Set
func (s *ItemSet) Add(item Item) *ItemSet {
	if s.items == nil {
		s.items = make(map[Item]bool)
	}

	_, ok := s.items[item]
	if !ok {
		s.items[item] = true
	}

	return s
}

// Clear remove all elements from the set
func (s *ItemSet) Clear() {
	s.items = make(map[Item]bool)
}

// Has return true if the set contains the Item
func (s *ItemSet) Has(item Item) bool {
	_, ok := s.items[item]
	return ok
}

// Delete remove the Item from the set and return Has(Item)
func (s *ItemSet) Delete(item Item) bool {
	_, ok := s.items[item]
	if ok {
		delete(s.items, item)
	}

	return ok
}

// Items return the Item(s) stored
func (s *ItemSet) Items() []Item {
	var items []Item
	for i := range s.items {
		items = append(items, i)
	}

	return items
}

// Size return the size of the set
func (s *ItemSet) Size() int {
	return len(s.items)
}
