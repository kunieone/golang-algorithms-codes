package linkedlist

import (
	"fmt"
)

type LinkNode struct {
	value int
	next  *LinkNode
}

type LinkedList struct {
	head *LinkNode
}

func (list *LinkedList) String() string {
	str := "[head]"
	for p := list.head; p.next != nil; p = p.next {
		str += "->" + fmt.Sprintf("(%v)", p.next.value)
	}
	return str
}

func NewLL(values ...int) *LinkedList {
	ll := &LinkedList{
		head: &LinkNode{},
	}
	p := ll.head
	for i := 0; i < len(values); i++ {
		p.next = &LinkNode{value: values[i]}
		p = p.next
	}
	return ll
}

/////////////////////////////////////////////////////////////////

/*
# 设单链表不带表头结点，编写递归算法删除单链表中所有值为 x 的元素
input: LinkNode(0,1,0) , 0
output: LinkNode(1)
input: LinkNode(0,1,0,2,3) , 2
output: LinkNode(0,1,0,3)
*/
func Solution(n *LinkNode, x int) {
	if n == nil {
		return
	}
	if n.value == x {
		n.value = n.next.value
		n.next = n.next.next
	}
	if n.next.value == x && n.next.next == nil {
		n.next = nil
		return
	}
	Solution(n.next, x)
}

func Solution2(p **LinkNode, x int) { //使用二级指针的方法
	fmt.Printf("p:%p", p)
	if (*p) == nil {
		return
	}
	if (*p).value == x {
		*p = (*p).next
		Solution2(p, x)
	} else {
		Solution2((&(*p).next), x)
	}
}
