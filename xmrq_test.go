package xmrq

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue()
	for i := 0; i < 100; i++ {
		q.Add(i)
	}

	for i := 0; i < 100; i++ {
		if q.Peek().(int) != i {
			t.Error(q.Peek())
		}
		q.Remove()
	}
}
