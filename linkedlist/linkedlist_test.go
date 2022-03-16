package linkedlist

import (
	"fmt"
	"testing"
)

var l ItemLinkedList

func TestItemLinkedList_Append(t *testing.T) {
	if !l.IsEmpty() {
		t.Error("list should be empty")
	}

	l.Append("first")
	if l.IsEmpty() {
		t.Error("list should not be empty")
	}

	if size := l.Size(); size != 1 {
		t.Errorf("wrong count, expected 1 and got %d", size)
	}

	l.Append("second")
	l.Append("third")
	if size := l.Size(); size != 3 {
		t.Errorf("wrong count, expected 3 and got %d", size)
	}
}

func TestRemoveAt(t *testing.T) {
	_, err := l.RemoveAt(1)
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}

	if size := l.Size(); size != 2 {
		t.Errorf("wrong count, expected 2 and got %d", size)
	}
}

func TestInsert(t *testing.T) {
	//test inserting in the middle
	err := l.Insert(2, "second2")
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}
	if size := l.Size(); size != 3 {
		t.Errorf("wrong count, expected 3 and got %d", size)
	}

	//test inserting at head position
	err = l.Insert(0, "zero")
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}
}

func TestIndexOf(t *testing.T) {
	if i := l.IndexOf("zero"); i != 0 {
		t.Errorf("expected position 0 but got %d", i)
	}
	if i := l.IndexOf("first"); i != 1 {
		t.Errorf("expected position 1 but got %d", i)
	}
	if i := l.IndexOf("second2"); i != 2 {
		t.Errorf("expected position 2 but got %d", i)
	}
	if i := l.IndexOf("third"); i != 3 {
		t.Errorf("expected position 3 but got %d", i)
	}
}

func TestHead(t *testing.T) {
	l.Append("zero")
	h := l.Head()
	if "zero" != fmt.Sprint(h.content) {
		t.Errorf("Expected `zero` but got %s", fmt.Sprint(h.content))
	}
}

func TestItemLinkedList_String(t *testing.T) {
	l.String()
}
