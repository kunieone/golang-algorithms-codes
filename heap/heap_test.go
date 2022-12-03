package heap

import (
	"fmt"
	"kaoyan/utils"
	"testing"
	"time"
)

const (
	RANDOM_NUM = 100_000
)

func Test_PopValue(t *testing.T) {
	arrN := utils.SpawnArrN(5)
	heap := BuildHeapQuick(MaxHeap, arrN...)
	v1, _ := heap.Pop()
	v2, _ := heap.Pop()
	v3, _ := heap.Pop()
	v4, _ := heap.Pop()
	v5, err := heap.Pop()
	v6, err2 := heap.Pop()
	v7, err3 := heap.Pop()
	fmt.Println(heap, v1, v2, v3, v4, v5, v6, v7)
	fmt.Println(err, err2, err3)

}

// 基准测试- 快速建堆法（自底向上沉底）O(n)
func Test_BuildHeapQuick(t *testing.T) {
	arrN := utils.SpawnArrN(RANDOM_NUM)
	start := time.Now() // 获取当前时间
	heap := BuildHeapQuick(MaxHeap, arrN...)
	cost := time.Since(start) // 计算此时与start的时间差
	fmt.Println("开销:", cost, heap.len())
	fmt.Println([]int{1, 2, 3}[:2])
}

// 基准测试- Floyd建堆法（自顶向下上浮）O(nlogn)
func Test_BuildHeapSlow(t *testing.T) {
	arrN := utils.SpawnArrN(RANDOM_NUM)
	start := time.Now() // 获取当前时间
	heap := BuildHeapSlow(MaxHeap, arrN...)
	cost := time.Since(start) // 计算此时与start的时间差
	fmt.Println("开销:", cost, heap.len())

}

// 测试删除成功
func Test_DeleteHeapElement(t *testing.T) {
	arrN := utils.SpawnArrN(12)
	heap := BuildHeapQuick(MaxHeap, arrN...)
	fmt.Println(heap.Delete(-1), heap)
	fmt.Println(heap.Delete(-2), heap)
	fmt.Println(heap.Delete(12), heap)
	fmt.Println(heap.Delete(13), heap)
	fmt.Println(heap.Delete(6), heap)
	fmt.Println(heap.Delete(7), heap)
}
