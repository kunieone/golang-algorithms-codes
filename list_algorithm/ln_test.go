package linkedlist

import (
	"testing"

	"github.com/kr/pretty"
)

func TestSolution(t *testing.T) {
	node := NewLL(0, 1, 0).head.next
	Solution(node, 1)
	pretty.Println(node) //expect : 1
	node2 := NewLL(0, 1, 0, 2, 3).head.next
	Solution(node2, 2)
	pretty.Println(node) //expect : 0,1,0,3
}
