package binarysearchtree

import (
	"fmt"
	"sync"

	"github.com/lushenle/awesome-golang-algorithm/generic"
)

// Item the type of the binary search tree
type Item generic.Type

// Node a single node that composes the tree
type Node struct {
	key   int
	value Item
	left  *Node //left
	right *Node //right
}

// ItemBinarySearchTree the binary search tree of Items
type ItemBinarySearchTree struct {
	root *Node
	lock sync.RWMutex
}

// insertNode internal function to find the correct place for a node in a tree
func insertNode(node, newNode *Node) {
	if newNode.key < node.key {
		if node.left == nil {
			node.left = newNode
		} else {
			insertNode(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			insertNode(node.right, newNode)
		}
	}
}

// Insert inserts the Item t in the tree
func (bst *ItemBinarySearchTree) Insert(key int, value Item) {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	n := &Node{key: key, value: value, left: nil, right: nil}
	if bst.root == nil {
		bst.root = n
	} else {
		insertNode(bst.root, n)
	}
}

// inOrderTraverse internal recursive function to traverse in order
func inOrderTraverse(node *Node, f func(Item)) {
	if node != nil {
		inOrderTraverse(node.left, f)
		f(node.value)
		inOrderTraverse(node.right, f)
	}
}

// InOrderTraverse visits all nodes with in-order traversing
func (bst *ItemBinarySearchTree) InOrderTraverse(f func(item Item)) {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	inOrderTraverse(bst.root, f)
}

// internal recursive function to traverse pre-order
func preOrderTraverse(node *Node, f func(Item)) {
	if node != nil {
		f(node.value)
		preOrderTraverse(node.left, f)
		preOrderTraverse(node.right, f)
	}
}

// PreOrderTraverse visits all nodes with pre-order traversing
func (bst *ItemBinarySearchTree) PreOrderTraverse(f func(Item)) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	preOrderTraverse(bst.root, f)
}

// postOrderTraverse internal recursive function to traverse post order
func postOrderTraverse(node *Node, f func(Item)) {
	if node != nil {
		postOrderTraverse(node.left, f)
		postOrderTraverse(node.right, f)
		f(node.value)
	}
}

// PostOrderTraverse visits all nodes with post-order traversing
func (bst *ItemBinarySearchTree) PostOrderTraverse(f func(Item)) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	postOrderTraverse(bst.root, f)
}

// Min returns the Item with min value stored in the tree
func (bst *ItemBinarySearchTree) Min() *Item {
	bst.lock.RLock()
	defer bst.lock.RUnlock()

	n := bst.root
	if n == nil {
		return nil
	}

	for {
		if n.left == nil {
			return &n.value
		}
		n = n.left
	}
}

// Max returns the Item with max value stored in the tree
func (bst *ItemBinarySearchTree) Max() *Item {
	bst.lock.RLock()
	defer bst.lock.RUnlock()

	n := bst.root
	if n == nil {
		return nil
	}

	for {
		if n.right == nil {
			return &n.value
		}
		n = n.right
	}
}

// search internal recursive function to search an item in the tree
func search(node *Node, key int) bool {
	if node == nil {
		return false
	}

	if key < node.key {
		return search(node.left, key)
	}
	if key > node.key {
		return search(node.right, key)
	}

	return true
}

// Search returns true if the Item t exists in the tree
func (bst *ItemBinarySearchTree) Search(key int) bool {
	bst.lock.RLock()
	defer bst.lock.RUnlock()

	return search(bst.root, key)
}

// remove internal recursive function to remove an item
func remove(node *Node, key int) *Node {
	if node == nil {
		return nil
	}

	if key < node.key {
		node.left = remove(node.left, key)
		return node
	}
	if key > node.key {
		node.right = remove(node.right, key)
		return node
	}

	// key == node.key
	if node.left == nil && node.right == nil {
		node = nil
		return nil
	}
	if node.left == nil {
		node = node.right
		return node
	}
	if node.right == nil {
		node = node.left
		return node
	}

	leftMostRightSide := node.right
	for {
		// find the smallest value on the right side
		if leftMostRightSide != nil && leftMostRightSide.left != nil {
			leftMostRightSide = leftMostRightSide.left
		} else {
			break
		}
	}

	node.key, node.value = leftMostRightSide.key, leftMostRightSide.value
	node.right = remove(node.right, node.key)

	return node
}

// Remove removes the Item with key `key` from the tree
func (bst *ItemBinarySearchTree) Remove(key int) {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	remove(bst.root, key)
}

// stringify internal recursive function to print a tree
func stringify(node *Node, level int) {
	if node != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(node.left, level)
		fmt.Printf(format+"%d\n", node.key)
		stringify(node.right, level)
	}
}

// String prints a visual representation of the tree
func (bst *ItemBinarySearchTree) String() {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	fmt.Println("----------------------------------")
	stringify(bst.root, 0)
	fmt.Println("----------------------------------")
}
