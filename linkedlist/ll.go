package linkedlist

import (
	"fmt"
)

type LNode struct {
	value int
	next  *LNode
}
type List struct {
	head *LNode
}

func NewLL(values ...int) *List {
	ll := &List{
		head: &LNode{},
	}
	p := ll.head
	for i := 0; i < len(values); i++ {
		p.next = &LNode{value: values[i]}
		p = p.next
	}
	return ll
}

// 删除一个结点
func (l *List) DeleteOne(item int) bool {

	for p := l.head; p.next != nil; p = p.next {
		if p.next.value == item {
			p.next = p.next.next
			return true
		}
	}
	return false
}

// 删除一个结点
func (l *List) DeleteAll(item int) bool {
	hasKey := false
	p := l.head
	for p.next != nil {
		if p.next.value == item {
			p.next = p.next.next
			hasKey = true
		} else {
			p = p.next
		}
	}
	return hasKey
}

// 删除重复结点 返回是否存在重复结点
func (l *List) DeleteRepeat() bool {
	compareArr := []int{}
	flag := false
	hasRepeat := func(item int) bool {
		for _, v := range compareArr {
			if v == item {
				return true
			}
		}
		return false
	}
	p := l.head
	for p.next != nil {
		if !hasRepeat(p.next.value) {
			compareArr = append(compareArr, p.next.value)
			p = p.next
		} else {
			flag = true
			p.next = p.next.next
		}
	}
	return flag
}

func (list *List) String() string {
	str := "[head]"
	for p := list.head; p.next != nil; p = p.next {
		str += "->" + fmt.Sprintf("(%v)", p.next.value)
	}
	return str
}

/*
### 05
试编写在单向链表中删除值相同的多余节点的方法（要求不使用辅助空间）
### 06
*/
func (l *List) DeleteRepeat2() {
	hasRepeat := func(p *LNode) bool {

		for q := p; q.next.next != nil; q = q.next {
			if p.next.value == q.next.next.value {
				return true
			}
		}
		return false
	}
	p := l.head
	for p.next != nil {
		if !hasRepeat(p) {
			p = p.next
		} else {
			p.next = p.next.next
		}
	}
}

/*
InsertAjToBjWithLen
### 12-1
1. 已知两个单链表 A 和 B，其头指针分别为 heada 和 headb，编写一个过程从单链表 A 中删除自第 i 个元素起共 len 个元素，然后将单链表 A 插入到单链表 B 的第 j 个元素之前。
2. 编写一个算法在基于单链表表示的待排序关键字序列上进行简单选择排序。
*/
func InsertAjToBjWithLen(A *List, B *List, i int, len int, j int) {
	ptrA := A.head
	ptrB := B.head
	for ii := 0; ii < i-1 && ptrA != nil; ii++ {
		ptrA = ptrA.next
	}

	for l := 0; l < len && ptrA.next != nil; l++ {
		ptrA.next = ptrA.next.next
	}
	fmt.Println("ptrA:", ptrA)
	for ptrA.next != nil {
		ptrA = ptrA.next //ptrA指针到尾部的前驱
	}
	fmt.Println("ptrA:", ptrA)
	for jj := 0; jj < j-1 && ptrB != nil; jj++ { //移动到带插入结点的前驱
		ptrB = ptrB.next
	}
	fmt.Println(ptrA, ptrB, "\n", "A.head.next != nil , ptrA.next != nil:", A.head.next != nil, ",", ptrA.next != nil)
	if A.head.next != nil && ptrA != nil {
		right := ptrB.next
		ptrB.next = A.head.next
		ptrA.next = right
	}
	fmt.Println(A, B)
}

/*
Sieve
07
设 L1 和 L2 为单链表表示的有序表，试编写算法求解 L1 交 L2，结果放在 L1 中。
input  A:1,3,5,6,7 B:2,3,4,5,7,8
output  A:3,5,7 B不变
*/
// 查询是否存在
// 取交集
func Sieve(A *List, B *List) {
	ptrB := B.head
	ptrA := A.head
	for ptrA.next != nil {
		flag := false
		for ptrB.next != nil { //->1.成功 2.失败 3.循环走完->(还是失败)
			if ptrB.next.value == ptrA.next.value { // 成功
				flag = true
				break
			} else if ptrA.next.value < ptrB.next.value { //已经失败,删掉这个元素
				flag = false
				break
			} else {
				ptrB = ptrB.next
			}
		}
		if flag {
			ptrA = ptrA.next
		} else {
			ptrA.next = ptrA.next.next
		}
	}
}

/*
RecusiveDel
### 10 设单链表不带表头结点，编写递归算法删除单链表中所有值为 x 的元素
*/
// 这里*LNode是链表结点的引用类型,p是**LNode，是指向链表结点的指针，通过解引用*p可以修改LNode的值
func RecusiveDel(p **LNode, x int) {
	fmt.Printf("p:%p", p)
	if (*p) == nil {
		return
	}
	if (*p).value == x {
		*p = (*p).next
		RecusiveDel(p, x)
	} else {
		RecusiveDel((&(*p).next), x)
	}
}

// 这段是错误的，无法对第一个结点处理
func RecusiveDel_Wrong(p *LNode, x int) {
	fmt.Printf("p:%p", p)
	if p == nil {
		return
	}
	if p.value == x {
		p = p.next
		RecusiveDel_Wrong(p, x)
	} else {
		RecusiveDel_Wrong(p.next, x)
	}
}
