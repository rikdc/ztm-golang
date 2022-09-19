package queue

import "testing"

func TestAddQueue(t *testing.T) {
	q := New(3)
	for i := 0; i < 3; i++ {
		if len(q.items) != i {
			t.Errorf("incorrect queue element count: %v, want %v", len(q.items), i)
		}

		if !q.Append(i) {
			t.Errorf("failed to appned item %v to queue", i)
		}
	}

	if q.Append(4) {
		t.Errorf("should not be able to add to a full queue")
	}
}

func TestNext(t *testing.T) {
	q := New(3)
	for i := 0; i < 3; i++ {
		q.Append(i)
	}

	for i := 0; i < 3; i++ {
		item, ok := q.Next()
		if !ok {
			t.Errorf("could not get item %v from the queue", i)
		}

		if item != i {
			t.Errorf("unexpected item returned from queue, expected %v, got %v", i, item)
		}
	}

	item, ok := q.Next()
	if ok {
		t.Errorf("queue should be empty. got: %v", item)
	}
}
