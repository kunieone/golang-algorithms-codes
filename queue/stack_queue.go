package queue

import "kaoyan/stack"

type SQueue struct {
	queue [2]*stack.Stack[int]
}

func NewSQueue() *SQueue {
	stack2 := [2]*stack.Stack[int]{stack.NewStack[int](), stack.NewStack[int]()}
	return &SQueue{queue: stack2}
}

func (s *SQueue) EnQueue(v int) {
	s.queue[0].Push(v)
}
func (s *SQueue) IsEmpty() bool {
	return s.queue[0].IsEmpty() && s.queue[1].IsEmpty()
}
func (s *SQueue) Dequeue() int {
	for !s.queue[0].IsEmpty() {
		s.queue[1].Push(s.queue[0].Pop())
	}
	return s.queue[1].Pop()
}
