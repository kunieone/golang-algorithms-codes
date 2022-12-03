package linkedlist

import (
	"fmt"
	"testing"

	"github.com/kr/pretty"
)

func TestNewLL(t *testing.T) {
	ll := NewLL(1, 2, 2, 3)
	pretty.Println(ll)
	ll.DeleteOne(2)
	fmt.Println(ll)
}

func TestDeleteAll(t *testing.T) {
	ll := NewLL(1, 2, 2, 3)
	pretty.Println(ll)
	ll.DeleteAll(2)
	fmt.Println(ll)
}

func Test12_1_1(t *testing.T) {
	A := NewLL(1, 2, 3, 4, 5, 6)
	B := NewLL(7, 8, 9)
	InsertAjToBjWithLen(A, B, 7, 1, 3) //=> B: 4,5,6,7,8,9
	// fmt.Println(B)
}

func TestDeleteRepeat(t *testing.T) {
	ll := NewLL(2, 2, 3, 3, 4, 4, 5, 6, 7, 7, 8, 9, 9, 1)
	ll.DeleteRepeat()
	fmt.Println(ll)
}
func TestDeleteRepeat2(t *testing.T) {
	ll := NewLL(2, 2, 3, 3, 4, 4, 5, 6, 7, 7, 8, 9, 9, 1)
	ll.DeleteRepeat2()
	fmt.Println(ll)
}

func TestSieve(t *testing.T) {
	A := NewLL(1, 2, 3, 4, 5, 6)
	B := NewLL(1, 2, 3, 4, 5)
	Sieve(A, B)
	fmt.Println(A)
}

func TestRecDel(t *testing.T) {
	A := *NewLL(0, 1, 0)

	noHeadNode := A.head.next
	fmt.Printf("p:%p\n", &noHeadNode)
	RecusiveDel(&noHeadNode, 0)
	pretty.Println(noHeadNode)
}

func TestRecDel_Wrong(t *testing.T) {
	A := *NewLL(0, 1, 0)

	noHeadNode := A.head.next
	fmt.Printf("p:%p\n", noHeadNode)
	// RecusiveDel(&noHeadNode, 0)
	RecusiveDel_Wrong(noHeadNode, 0)
	pretty.Println(&noHeadNode)
}

func TestRecDel_3rdSolution(t *testing.T) {
	A := *NewLL(0, 1, 0)
	node := *A.head.next
	RecusiveDel_Wrong(&node, 1)
	pretty.Println(node)

}
