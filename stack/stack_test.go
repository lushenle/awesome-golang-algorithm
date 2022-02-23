package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := New()

	if s.Len() != 0 {
		t.Errorf("error: length should be 0")
	}

	s.Push(1)
	if s.Len() != 1 {
		t.Errorf("error: length should be 1")
	}

	if s.Peek().(int) != 1 {
		t.Errorf("error: top value should be 1")
	}

	if s.Pop().(int) != 1 {
		t.Errorf("error: top value should be 1")
	}

	s.Push(2)

	if s.Len() != 2 {
		t.Errorf("error: length should be 2")
	}

	if s.Peek().(int) != 2 {
		t.Errorf("error: top value should be 2")
	}

	if s.IsEmpty() {
		t.Errorf("error: empty stack have no values")
	}

	if s.length == 0 {
		t.Errorf("error:")
	}

	if s.Pop().(int) != 2 {
		t.Errorf("error: top value should be 2")
	}
}
