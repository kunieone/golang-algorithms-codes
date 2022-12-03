package tree

import (
	"fmt"
	"testing"
)

func TestTreePostOrderTraverse(t *testing.T) {
	tree := NewTree(2)
	tree.LChild = NewTree(3)
	tree.LChild.RChild = NewTree(4)
	tree.RChild = NewTree(1)
	tree.PreOrder(func(i int) { fmt.Println(i) })
	fmt.Println("#############")
	tree.MidOrder(func(i int) { fmt.Println(i) })
	fmt.Println("#############")
	tree.PostOrder(func(i int) { fmt.Println(i) })
	fmt.Println("#############")
}
