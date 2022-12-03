package queue

import "fmt"

type Queue struct {
	Queue []int
	Size  int
	state int
}

func NewQueue(size int) *Queue {
	return &Queue{[]int{}, size, 0}
}

func (q *Queue) Enqueue(v int) bool {
	if q.IsFull() {
		fmt.Println("队列已满")
		return false
	}
	q.Queue = append(q.Queue, v)
	q.state++
	return true
}
func (q *Queue) IsEmpty() bool {
	return q.state == 0
}
func (q *Queue) IsFull() bool {
	return q.Size == q.state
}
func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		panic("队列已空，无法出队")
	}
	first := q.Queue[0]
	q.Queue = q.Queue[1:] //把0号元素截下来
	q.state--
	return first
}
