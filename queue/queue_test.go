package queue

import "testing"

func TestQueue(t *testing.T) {
	q := New()

	if q.Len() != 0 {
		t.Errorf("error: length should be 0")
	}

	q.Enqueue(1)
	if q.Len() != 1 {
		t.Errorf("error: length should be 1")
	}

	if q.Peek().(int) != 1 {
		t.Errorf("error: enqueue value should be 1")
	}

	v := q.Dequeue()
	if v.(int) != 1 {
		t.Errorf("error: dequeued value should be 1")
	}

	if q.Peek() != nil || q.Dequeue() != nil {
		t.Errorf("empty queue should have no values")
	}

	q.Enqueue(1)
	q.Enqueue(2)

	if q.Len() != 2 {
		t.Errorf("error: length should be 0")
	}

	if q.Peek().(int) != 1 {
		t.Errorf("error: first value should be 1")
	}

	q.Dequeue()
	if q.Peek().(int) != 2 {
		t.Errorf("error: next value should be 2")
	}
}
