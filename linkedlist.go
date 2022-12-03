package main

import (
	"fmt"
)

type LNode[T comparable] struct {
	value T
	next  *LNode[T]
}
type List[T comparable] struct{ head *LNode[T] }

func NewList[T comparable](values ...T) *List[T] {

	list := &List[T]{head: &LNode[T]{}}
	list.head.value = values[0]
	ptr := list.head
	for i := 1; i < len(values); i++ {
		ptr.next = &LNode[T]{value: values[i]}
		ptr = ptr.next
	}
	return list
}

func (list *List[T]) String() string {
	pool := "[head]"
	if list.head != nil {
		for p := list.head; p != nil; p = p.next {
			pool += "->" + fmt.Sprintf("(%v)", p.value)
		}
	}
	return pool + "\n"
}

/*
一个存放实数的顺序表以带头结点的单链表形式存储，编写算法删除表中一个值最大结点和一个最小结点，并计算平均值。
*/

func Solve20_2(list *List[int]) float32 {
	sum, len := list.head.value, 0
	minNode, maxNode := list.head, list.head
	for p := list.head; p.next != nil; p = p.next {
		len++
		sum += p.next.value
		if minNode.value > p.next.value {
			minNode = p /* 这里获得的其实是它的前驱 */
		}
		if maxNode.value < p.next.value {
			maxNode = p /* 这里获得的其实是它的前驱 */
		}
		println(p.value)
	}
	sum -= (maxNode.next.value + minNode.next.value)
	maxNode.next = maxNode.next.next
	minNode.next = minNode.next.next
	if len != 2 {
		return float32(sum) / float32(len-2)
	} else {
		return 0
	}
}
func Solve20(list *List[int]) float32 {
	minNode, maxNode := &list.head, &list.head
	len, sum := 0, 0
	// 这是第2种方法 是使用对最值结点指针的指针，并对指针的指针解引用(即结点指针本身)赋值，从而直接修改指针(多比较了一次)
	for p := &list.head; *p != nil; p = &(*p).next {
		len++
		sum += (*p).value
		// fmt.Println(len, sum)
		if (*p).value < (*minNode).value {
			minNode = p
		}
		if (*p).value > (*maxNode).value {
			maxNode = p
		}
	}
	res := sum - ((*minNode).value + (*maxNode).value)
	*minNode = (*minNode).next
	*maxNode = (*maxNode).next
	if len != 2 {
		return float32(res) / float32(len-2)
	} else {
		return 0
	}
}

/*
一个带有头结点的单链表，结点数据类型为整形。

(1)编写算法：结点元素为负整数的放在链表前面，将结点元素为正整数的放在链表后面。

(2)对于上述问题使用何种物理结构实现比较好
*/

func Solve19(list *List[int]) {
	for p := list.head; p.next != nil; {
		if p.next.value < 0 {
			nextNode := p.next
			p.next = p.next.next
			list = list.appendToHead(nextNode)
		} else {
			p = p.next
		}
	}
}
func (l *List[int]) appendToHead(node *LNode[int]) *List[int] {
	newList := &List[int]{
		head: node,
	}
	newList.head.next = l.head
	return newList
}
func (l *List[T]) Reverse() {
	for p := l.head; p.next != nil; {
		nextNode := p.next
		p.next = p.next.next
		*l = *l.appendToHead(nextNode)
	}
}

/* 递归翻转 */
func RecReverse[T comparable](head *LNode[T]) *LNode[T] {
	if head == nil || head.next == nil {
		return head
	}
	next := head.next
	newHead := RecReverse(head.next)
	next.next = head
	head.next = nil
	return newHead
}

/*
一个线性表的元素均为正整数，使用带头指针的单链表实现。编写算法：判断该链表是否符合：所有奇数在前面，偶数在后面。
*/
func Solve18(l *List[int]) bool {
	evenAppear := false
	//链表的传统迭代:
	for p := l.head; p != nil; p = p.next {
		if p.value%2 == 0 {
			evenAppear = true
		}
		// println(p.value)
		if evenAppear && p.value%2 == 1 {
			return false
		}
	}
	return true

}

func Iterate(l *List[int]) {
	for p := l.head; p != nil; p = p.next {
		fmt.Println(p.value)
	}
}

/*
已知带头结点的单链表 A，
要求设计算法生成一个新的链表 B，使得 B 含有 A 中的元素，且次序不变。
但不包含 A 中的重复元素。如 A 的元素为（1，7，7，3，5，3，1），则 B 的元素依次为（1，7，3，5）
*/
func Solve16(l *List[int]) *List[int] {
	unitArr := []int{l.head.value}
	hasRepeat := func(elem int) bool {
		for i := 0; i < len(unitArr); i++ {
			if elem == unitArr[i] {
				return true
			}
		}
		return false
	}

	for p := l.head; p.next != nil; {
		fmt.Println("比较:", p.next.value)
		if !hasRepeat(p.next.value) {

			unitArr = append(unitArr, p.next.value)
			p = p.next
		} else {
			fmt.Println("--重复:", p.next.value)
			p.next = p.next.next
		}
	}
	return l
}
func TestSolve16() {
	listA := NewList(1, 7, 7, 3, 5, 3, 1)
	listB := Solve16(listA)
	fmt.Println(listB)
}

/*
### 15
一个由整数元素构成的递增有序线性表存放在一个双向链表中，
设计一个时间复杂度为 O(n)的算法，在链表中获得两个和为 x
的结点的值，并以 x=a+b 的形式输出；若不存在则给出提示信息。
*/
func Solve15() {
}

/*
### 14

将数的质因数分解按递减顺序写成一个有序单链表，如 2100->7.5.5.3.2.2
*/
func Solve14(n int) {
	fmt.Println(n, ":")
	var emptyList *List[int] = &List[int]{head: nil}
	for i := 2; i < n; {
		if n%i == 0 {
			emptyList = emptyList.appendToHead(&LNode[int]{value: i})
			n /= i
		} else {
			i++
		}
	}
	emptyList = emptyList.appendToHead(&LNode[int]{value: n})
	fmt.Println(emptyList)
}

/*
	### 13

写出递归删除单链表中所有值为 item 的算法
*/
/*
void del(LinkList *&L,ElemType x){
	  Node *t;
   if (L==NULL) return;
   if (L->data==x)
    {  t=L;
	  L=L->next;
	  free(t);
	  return;
    }
   else del(L->next,x);
}

*/

/*
	## T1

查找链表（带头结点）中是否存在一个值为 x 的结点，如果存在，则删除结点并返回 1，否则返回 0
*/
func T1(l *LNode[int], x int) int {
	var p *LNode[int] = l
	for p.next != nil {
		if p.next.value == x {
			break
		}
		p = p.next
	}
	if p.next == nil {
		return 0
	} else {
		// 删除结点
		p.next = p.next.next //原先p.next的引用计数为0，自动释放
		return 1
	}

}

func Solve13_1[T comparable](L **LNode[T], item T) {
	// var t LNode[T]
	if L == nil {
		return
	}
	if (*L).value == item {
		// t = **L
		*L = (*L).next
		return
	} else {
		Solve13_1(&(*L).next, item)
	}
}

/* 13-迭代解法 */
func Solve13[T comparable](list *List[T], item T) *List[T] {
	if list.head.value == item {
		list.head = list.head.next
	}
	if list.head.next != nil {
		for p := list.head; p.next != nil; {
			if p.next.value == item {
				p.next = p.next.next
			} else {
				p = p.next
			}
		}
	}
	return list
}
