package linkedlist

import (
	"errors"
	"fmt"
	"sync"

	"github.com/lushenle/awesome-golang-algorithm/generic"
)

// Item the type of the linked list
type Item generic.Type

// Node a single node that composes the list
type Node struct {
	content Item
	next    *Node
}

// ItemLinkedList the linked list of Items
type ItemLinkedList struct {
	head *Node
	size int
	lock sync.RWMutex
}

// Head return a pointer to the first node of the list
func (l *ItemLinkedList) Head() *Node {
	l.lock.Lock()
	defer l.lock.Unlock()

	return l.head
}

// Append add an Item to the end of the linked list
func (l *ItemLinkedList) Append(item Item) {
	l.lock.Lock()
	defer l.lock.Unlock()

	node := Node{content: item, next: nil}
	if l.head == nil {
		l.head = &node
	} else {
		last := l.head
		for {
			if last.next == nil {
				break
			}
			last = last.next
		}
		last.next = &node
	}
	l.size++
}

// Insert add an Item at position i
func (l *ItemLinkedList) Insert(i int, item Item) error {
	l.lock.Lock()
	defer l.lock.Unlock()

	if i < 0 || i > l.size {
		return errors.New("index out of bounds")
	}

	addNode := Node{content: item, next: nil}
	if i == 0 {
		addNode.next = l.head
		l.head = &addNode
		return nil
	}
	node := l.head
	j := 0
	for j < i-2 {
		j++
		node = node.next
	}
	addNode.next = node.next
	node.next = &addNode
	l.size++

	return nil
}

// RemoveAt remove a node at position
func (l *ItemLinkedList) RemoveAt(i int) (*Item, error) {
	l.lock.Lock()
	defer l.lock.Unlock()

	if i < 0 || i > l.size {
		return nil, errors.New("index out of bounds")
	}

	node := l.head
	j := 0
	for j < i-1 {
		j++
		node = node.next
	}
	remove := node.next
	node.next = remove.next
	l.size--

	return &remove.content, nil
}

// IndexOf return the position of the Item
func (l *ItemLinkedList) IndexOf(item Item) int {
	l.lock.Lock()
	defer l.lock.Unlock()

	node := l.head
	j := 0
	for {
		if node.content == item {
			return j
		}
		if node.next == nil {
			return -1
		}
		node = node.next
		j++
	}
}

// IsEmpty return true if the list is empty
func (l *ItemLinkedList) IsEmpty() bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	return l.head == nil
}

// Size return the linked list size
func (l *ItemLinkedList) Size() int {
	l.lock.Lock()
	defer l.lock.Unlock()

	size, last := 1, l.head
	for {
		if last == nil || last.next == nil {
			break
		}
		last = last.next
		size++
	}

	return size
}

// String format output Items
func (l *ItemLinkedList) String() {
	l.lock.Lock()
	defer l.lock.Unlock()

	j, node := 0, l.head
	for {
		if node == nil {
			break
		}
		j++
		fmt.Print(node.content)
		fmt.Print(" ")
		node = node.next
	}
	fmt.Println()
}
