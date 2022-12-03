package tree

import "kaoyan/stack"

type TreeNode[T any] struct {
	Value  T
	LChild *TreeNode[T]
	RChild *TreeNode[T]
}

func NewTree[T any](v T) *TreeNode[T] {
	return &TreeNode[T]{v, nil, nil}
}
func (p *TreeNode[T]) PreOrder(visit func(T)) {
	s1 := stack.NewStack[*TreeNode[T]]()
	s1.Push(p)
	for !s1.IsEmpty() {
		node := s1.Pop()
		visit(node.Value)
		if node.RChild != nil {
			s1.Push(node.RChild)
		}
		if node.LChild != nil {
			s1.Push(node.LChild)
		}
	}
}

func (p *TreeNode[T]) PostOrder(visit func(T)) {
	s1 := stack.NewStack[*TreeNode[T]]()
	s2 := stack.NewStack[*TreeNode[T]]()
	s1.Push(p)
	for !s1.IsEmpty() {
		node := s1.Pop()
		s2.Push(node)

		if node.LChild != nil {
			s1.Push(node.LChild)
		}

		if node.RChild != nil {
			s1.Push(node.RChild)
		}

	}
	for !s2.IsEmpty() {
		visit(s2.Pop().Value)
	}
}
func (p *TreeNode[T]) MidOrder(visit func(T)) {
	stack := stack.NewStack[*TreeNode[T]]()
	stack.Push(p)
	for !stack.IsEmpty() || p != nil {
		for p != nil {
			stack.Push(p)
			p = p.LChild
		}
		node := stack.Pop()
		visit(node.Value)
		if node.RChild != nil {
			p = node.RChild
		}
	}
}
