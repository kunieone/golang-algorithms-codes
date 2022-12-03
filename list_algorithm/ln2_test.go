package linkedlist

import (
	"fmt"
	"testing"

	"github.com/kr/pretty"
)

func Test_remove(t *testing.T) {
	n := NewLLN(4, 4, 4)
	// n.Remove(2)
	fmt.Println(n)

	Revm(&n, 4)

	// pretty.Println(n)
	fmt.Println(n)
}

func TestPointer(t *testing.T) {
	node := LLN{1, nil}
	refNode := &node
	refNode = &LLN{2, nil}
	fmt.Printf("%p\n", &node)   //0x1400005c5c0
	fmt.Printf("%p\n", refNode) //0x1400005c5c0
}
func TestSelect(t *testing.T) {
	node := NewLLN(-114514, 7, 2, 1, 6, 4, 5, 3)
	list := &LLNList{node}
	Select(list)
	fmt.Println(list.head)
}

func TestSplit(t *testing.T) {
	A := NewLLN(1, 2, 3, 4, 5, 6)
	B := &LLN{}
	Split(A, B)
	pretty.Println(A, B)
}

func TestList(t *testing.T) {
	f := NewFreqList()
	f.Append(2, 4, 5, 1, 12)
	visited(f, 1)
	visited(f, 1)
	visited(f, 12)
	fmt.Println(f)
}
