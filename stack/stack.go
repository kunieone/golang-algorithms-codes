package stack

type IStack interface {
	push()
	pop()
}

type Stack[T any] struct {
	arr []T
	len int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{make([]T, 0), 0}
}
func (s *Stack[T]) Push(v T) {
	s.arr = append(s.arr, v)
	s.len++
}
func (s *Stack[T]) IsEmpty() bool {
	return s.GetLen() == 0
}
func (s *Stack[T]) GetLen() int {
	return s.len
}
func (s *Stack[T]) Top() T {
	if s.len == 0 {
		panic("空栈")
	}
	return s.arr[len(s.arr)-1] //把最后一个元素拿出来

}
func (s *Stack[T]) Pop() T {
	if s.len == 0 {
		panic("空栈")
	}
	last := s.arr[len(s.arr)-1]  //把最后一个元素拿出来
	s.arr = s.arr[:len(s.arr)-1] //调整大小，删去最后一个元素 :后面是不包含.
	s.len--
	return last
}
