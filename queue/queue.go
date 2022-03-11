package queue

import (
	"sync"

	"github.com/lushenle/awesome-golang-algorithm/generic"
)

// Item the type of the queue
type Item generic.Type

// ItemQueue the queue of Items
type ItemQueue struct {
	items []Item
	lock  sync.RWMutex
}

// New create a new ItemQueue
func (q *ItemQueue) New() *ItemQueue {
	q.items = []Item{}
	return q
}

// Enqueue add an Item to the end of the queue
func (q *ItemQueue) Enqueue(it Item) {
	q.lock.Lock()
	q.items = append(q.items, it)
	q.lock.Unlock()
}

// Dequeue remove an Item from the start of the queue
func (q *ItemQueue) Dequeue() *Item {
	q.lock.Lock()
	item := q.items[0]
	q.items = q.items[1:len(q.items)]
	q.lock.Unlock()
	return &item
}

// Peek return the item next in the queue, without removing it
func (q *ItemQueue) Peek() *Item {
	q.lock.Lock()
	item := q.items[0]
	q.lock.Unlock()
	return &item
}

// IsEmpty return true if the queue is empty
func (q *ItemQueue) IsEmpty() bool {
	return len(q.items) == 0
}

// Len return the number of Items in the queue
func (q *ItemQueue) Len() int {
	return len(q.items)
}
