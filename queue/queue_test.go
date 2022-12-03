package queue

import (
	"fmt"
	"testing"
)

func TestQueue_Queue(t *testing.T) {
	q := NewQueue(5)
	fmt.Println(q, q.IsEmpty(), q.IsFull())
	q.Enqueue(1)
	fmt.Println(q, q.IsEmpty(), q.IsFull())
	q.Enqueue(2)
	fmt.Println(q, q.IsEmpty(), q.IsFull())
	q.Enqueue(3)
	fmt.Println(q, q.IsEmpty(), q.IsFull())
	q.Enqueue(4)
	fmt.Println(q, q.IsEmpty(), q.IsFull())
	q.Enqueue(5)
}
