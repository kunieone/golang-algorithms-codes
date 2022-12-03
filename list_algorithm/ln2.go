package linkedlist

import (
	"fmt"
)

type LLN struct {
	data int
	nex  *LLN
}

type LLNList struct {
	head *LLN
}

// type CharNode struct {
// 	data rune
// 	next *CharNode
// }

func NewLLN(v ...int) *LLN {
	n := &LLN{v[0], nil}
	p := n
	for i := 1; i < len(v); i++ {
		println("添加")
		p.nex = &LLN{v[i], nil}
		p = p.nex
	}
	return n
}
func (n *LLN) String() string {
	str := "[head]"
	if n == nil {
		return str + "[空]"
	}
	for {
		if n == nil {
			break
		}
		str += fmt.Sprintf("->%v", n.data)
		n = n.nex
	}
	return str
}
func (n *LLN) insert(v int) {
	n.nex = &LLN{v, nil}
}

func Revm(n **LLN, x int) {
	flag := false
	RemoveHelper(*n, x, &flag)
	if flag {
		*n = nil
	}
}
func RemoveHelper(n *LLN, x int, flag *bool) bool {
	if n.data == x && n.nex == nil {
		fmt.Println("最后一个结点也是，需要全部删除")
		*flag = true
		return false
	}
	// if n == nil {
	// 	return true
	// }
	if n.data == x && n.nex != nil {
		*n = *n.nex
	}
	if n.nex != nil {
		if n.nex.data == x {
			n.nex = n.nex.nex
			RemoveHelper(n, x, flag)
		} else {
			RemoveHelper(n.nex, x, flag)
		}
	}
	return true
}

/*
2. 编写一个算法在基于单链表表示的待排序关键字序列上进行简单选择排序。
*/
func Select(list *LLNList) {
	for p := list.head; p.nex != nil; p = p.nex {
		minPtr := p.nex
		for q := p.nex; q != nil; q = q.nex {
			fmt.Print(q.data)
			if minPtr.data > q.data {
				minPtr = q
			}
		}
		println("--", minPtr.data)
		p.nex.data, minPtr.data = minPtr.data, p.nex.data
	}

}

/*
### 11
设单链表不带表头结点，编写递归算法将一个单链表的元素按奇数节点和偶数节点拆分成 2 个表。
*/

func Split(n *LLN, receive *LLN) {
	if n.nex == nil {
		return
	}
	if n.data%2 == 0 {
		receive.nex = n
		*n = *n.nex
		Split(n, receive.nex)
	} else {
		Split(n.nex, receive)
	}

}

// ### 04
// 设有一个双向链表 L 每个节点中含有数值域 data 和访问频度域 freq，
// 在链表被起用前，频度域 freq 初始化为 0，而当对链表进行了
//  visited(L,X)的操作之后 data 值 X 的节点频度域增 1，同时调整
// 链表的节点次序使得按照频度值递减次序排列，编写复合要求的 visited(L,X)算法

type FreqNode struct {
	freq int
	data int
	next *FreqNode
	pre  *FreqNode
}
type FreqLinkedList struct {
	head *FreqNode
}

func (f *FreqNode) addNode(data int) *FreqNode {
	f.next = &FreqNode{0, data, nil, f}
	return f.next
}
func NewFreqList() *FreqLinkedList {
	return &FreqLinkedList{&FreqNode{0, -1, nil, nil}}
}
func (f *FreqLinkedList) Append(values ...int) {
	p := f.head
	for i := range values {
		p = p.addNode(values[i])
	}
}
func (list *FreqLinkedList) String() string {
	str := "[head]"
	for p := list.head; p.next != nil; p = p.next {
		str += "->" + fmt.Sprintf("(%v)", p.next.data)
	}
	return str
}
func visited(f *FreqLinkedList, data int) {
	p := f.head
	for p.next != nil {
		if p.next.data == data {
			break
		}
		p = p.next
	}
	p.next.freq += 1
	for p.pre != nil {
		// 1->2
		if p.next.freq <= p.freq {
			break
		}
		if p.next.freq > p.freq {
			// 交换p和p.next
			p.data, p.next.data = p.next.data, p.data
			p.freq, p.next.freq = p.next.freq, p.freq

		}
		p = p.pre

	}
}
