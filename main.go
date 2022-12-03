package main

import (
	"fmt"
)

func main() {
	list := NewList(1, 2, 3, 4)
	ave := Solve20(list)
	fmt.Println("平均值:", ave)
	fmt.Println(list)
	/* 19 */
	list2 := NewList(1, 2, -3, -4, 5, -6, 1, -9, -8)
	Solve19(list2)
	fmt.Println(list2)

	/* 翻转链表 */
	list3 := NewList(1, 2, 3, 4, 5, 6)
	list3.Reverse()
	fmt.Println(list3)
	h := list3.head
	list3.head = RecReverse(h)
	fmt.Println(list3)

	/* 18  */
	list4 := NewList(1, 3, 5, 2, 4, 6)
	fmt.Printf("%v =>%v\n", list4, Solve18(list4))
	list4 = NewList(2, 4, 6, 8, 0)

	fmt.Printf("%v =>%v\n", list4, Solve18(list4))
	list4 = NewList(2, 4, 6, 8, 1)
	fmt.Printf("%v =>%v\n", list4, Solve18(list4))
	list4 = NewList(2, 1)
	fmt.Printf("%v =>%v\n", list4, Solve18(list4))
	list4 = NewList(1, 2)
	fmt.Printf("%v =>%v\n", list4, Solve18(list4))
	Iterate(NewList(1, 2, 3, 4))

	/* 16 */
	TestSolve16()

	/* 14 */
	Solve14(44387)

	/* 13 */
	list5 := NewList(1, 2, 3, 1)
	list5 = Solve13(list5, 1)
	fmt.Println(list5)
	list6 := NewList(1, 2, 3, 1, 1, 1, 1, 1, 1, 1)
	Solve13_1(&list6.head, 1)
	fmt.Println(list6)
	// 	for i := range [3]interface{} {
	// i
	// 	}
	listT1 := NewList(2, 3, 7, 8, 9)

	fmt.Println(T1(listT1.head, 9), listT1)

}
