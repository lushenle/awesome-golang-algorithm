package binarysearchtree

import (
	"fmt"
	"testing"
)

var bst ItemBinarySearchTree

func fillTree(bst *ItemBinarySearchTree) {
	bst.Insert(8, "8")
	bst.Insert(4, "4")
	bst.Insert(10, "10")
	bst.Insert(2, "2")
	bst.Insert(6, "6")
	bst.Insert(1, "1")
	bst.Insert(3, "3")
	bst.Insert(5, "5")
	bst.Insert(7, "7")
	bst.Insert(9, "9")
}

func TestItemBinarySearchTree_Insert(t *testing.T) {
	fillTree(&bst)
	bst.String()

	bst.Insert(11, "11")
	bst.String()
}

// isSameSlice returns true if the 2 slices are identical
func isSameSlice(a, b []string) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func Test_InOrderTraverse(t *testing.T) {
	var result []string
	bst.InOrderTraverse(func(item Item) {
		result = append(result, fmt.Sprintf("%s", item))
	})

	if !isSameSlice(result, []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"}) {
		t.Errorf("Traversal order incorrect, got %v", result)
	}
}

func Test_PreOrderTraverse(t *testing.T) {
	var result []string
	bst.PreOrderTraverse(func(item Item) {
		result = append(result, fmt.Sprintf("%s", item))
	})
	if !isSameSlice(result, []string{"8", "4", "2", "1", "3", "6", "5", "7", "10", "9", "11"}) {
		t.Errorf("Traversal order incorrect, got %v instead of %v", result, []string{"8", "4", "2", "1", "3", "6", "5", "7", "10", "9", "11"})
	}
}

func Test_PostOrderTraverse(t *testing.T) {
	var result []string
	bst.PostOrderTraverse(func(item Item) {
		result = append(result, fmt.Sprintf("%s", item))
	})
	if !isSameSlice(result, []string{"1", "3", "2", "5", "7", "6", "4", "9", "11", "10", "8"}) {
		t.Errorf("Traversal order incorrect, got %v instead of %v", result, []string{"1", "3", "2", "5", "7", "6", "4", "9", "11", "10", "8"})
	}
}

func Test_Min(t *testing.T) {
	if fmt.Sprintf("%s", *bst.Min()) != "1" {
		t.Errorf("min should be 1")
	}
}

func Test_Max(t *testing.T) {
	if fmt.Sprintf("%s", *bst.Max()) != "11" {
		t.Errorf("max should be 11")
	}
}

func Test_Search(t *testing.T) {
	if !bst.Search(1) || !bst.Search(8) || !bst.Search(11) {
		t.Errorf("search not working")
	}
}

func Test_Remove(t *testing.T) {
	bst.Remove(1)
	if fmt.Sprintf("%s", *bst.Min()) != "2" {
		t.Errorf("min should be 2")
	}
}
