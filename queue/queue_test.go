package queue

import (
	"reflect"
	"testing"
)

var q ItemQueue

func initQueue() *ItemQueue {
	if q.items == nil {
		q = ItemQueue{}
		q.New()
	}

	return &q
}

func TestItemQueue_Enqueue(t *testing.T) {
	q := initQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if length := q.Len(); length != 3 {
		t.Errorf("wrong count, expected 3 and got %d", length)
	}
}

func TestItemQueue_Peek(t *testing.T) {
	if item := q.Peek(); reflect.DeepEqual(item, 3) {
		t.Errorf("Peek() = %v, want %v", item, 3)
	}
}

func TestItemQueue_Dequeue(t *testing.T) {
	q.Enqueue(1)
	q.Enqueue(1)
	q.Dequeue()

	if length := len(q.items); length != 4 {
		t.Errorf("wrong count, expected 2 and got %d", length)
	}

	q.Dequeue()
	q.Dequeue()

	if length := q.Len(); length != 2 {
		t.Errorf("wrong count, expected 0 and got %d", length)
	}
	q.Dequeue()
	q.Dequeue()

	if !q.IsEmpty() {
		t.Error("IsEmpty should return true")
	}
}
