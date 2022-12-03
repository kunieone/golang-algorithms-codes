package queue

import (
	"fmt"
	"testing"
)

func TestSQueue_EnQueue(t *testing.T) {
	stacQueue := NewSQueue()
	stacQueue.EnQueue(2)
	stacQueue.EnQueue(14)
	stacQueue.EnQueue(21)
	s1 := stacQueue.Dequeue()
	s2 := stacQueue.Dequeue()
	fmt.Println(stacQueue.IsEmpty())
	s3 := stacQueue.Dequeue()
	stacQueue.EnQueue(4)
	fmt.Println(s1, s2, s3, stacQueue.IsEmpty())
	stacQueue.Dequeue()
	fmt.Println(stacQueue.IsEmpty())

}
