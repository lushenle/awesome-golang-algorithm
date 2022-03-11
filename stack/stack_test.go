package stack

import (
	"reflect"
	"testing"
)

var s ItemStack

func initStack() *ItemStack {
	if s.items == nil {
		s = ItemStack{}
		s.New()
	}

	return &s
}

func TestItemStack_Push(t *testing.T) {
	s := initStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	if length := len(s.items); length != 3 {
		t.Errorf("wrong count, expected 3 and got %d", length)
	}
}

func TestItemStack_Peek(t *testing.T) {
	if item := s.Peek(); reflect.DeepEqual(item, 3) {
		t.Errorf("Peek() = %v, want %v", item, 3)
	}
}

func TestItemStack_Len(t *testing.T) {
	s.Pop()
	if length := s.Len(); length != 2 {
		t.Errorf("wrong count, expected 2 and got %d", length)
	}
}

func TestItemStack_Pop(t *testing.T) {
	s.Pop()
	s.Pop()
	if size := len(s.items); size != 0 {
		t.Errorf("wrong count, expected 0 and got %d", size)
	}

	if !s.IsEmpty() {
		t.Error("IsEmpty should return true")
	}
}
