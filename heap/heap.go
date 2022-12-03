package heap

import (
	"fmt"
	"kaoyan/utils"
)

type HeapType int

const (
	MaxHeap HeapType = iota
	MinHeap
)

type Heap[T utils.Ord] struct {
	Values []T
	model  HeapType
}

func (h Heap[T]) String() string {
	return fmt.Sprintf("---%v---\n%v\n", h.model, h.Values)
}

func (s HeapType) String() string {
	str := ""
	switch s {
	case MaxHeap:
		str = "最大堆[maxHeap]"
	case MinHeap:
		str = "最小堆[minHeap]"
	}
	return str
}

func (h *Heap[T]) compare(idxs ...int) int {
	isMax := h.model == MaxHeap
	com := func(a, b int) int {
		if h.Values[a] > h.Values[b] == isMax {
			return a
		} else {
			return b
		}
	}
	first := idxs[0]
	for i := 1; i < len(idxs); i++ {
		first = com(idxs[i], first)
	}
	return first
}
func (h *Heap[T]) swap(a, b int) {
	h.Values[a], h.Values[b] = h.Values[b], h.Values[a]
}
func (h *Heap[T]) len() int {
	return len(h.Values)
}
func (h *Heap[T]) onlyLeft(idx int) bool {
	return idx*2+1 == h.len()-1
}
func (h *Heap[T]) isLeaf(idx int) bool {
	return h.len()-1 < idx*2+1

}
func (h *Heap[T]) sinkDown(idx int) {
	if h.isLeaf(idx) {
		return
	} else {
		if h.onlyLeft(idx) {
			top := h.compare(idx, idx*2+1)
			if top != idx {
				h.swap(top, idx)
				h.sinkDown(top)
			}
		} else {
			top := h.compare(idx, idx*2+1, idx*2+2)
			if top != idx {
				h.swap(top, idx)
				h.sinkDown(top)
			}
		}
	}
}
func (h *Heap[T]) father(idx int) int {
	return (idx - 1) / 2
}
func (h *Heap[T]) siftUp(index int) {
	father := h.father(index)
	if father >= 0 {

		if h.model == MaxHeap {
			if h.Values[father] < h.Values[index] {
				h.Values[father], h.Values[index] = h.Values[index], h.Values[father]
				h.siftUp(father)
			}
		} else {
			if h.Values[father] > h.Values[index] {
				h.Values[father], h.Values[index] = h.Values[index], h.Values[father]
				h.siftUp(father)
			}
		}
	}

}

func (h *Heap[T]) Insert(item T) {
	h.Values = append(h.Values, item)
	h.siftUp(len(h.Values) - 1)
}
func (h *Heap[T]) Inserts(items ...T) {
	for _, item := range items {
		h.Values = append(h.Values, item)
		h.siftUp(len(h.Values) - 1)
	}
}
func (h *Heap[T]) beyond(idx int) bool {
	return idx >= h.len()
}

func (h *Heap[T]) Delete(idx int) bool {
	if idx < 0 || h.len() == 0 || h.beyond(idx) { //对于空堆和数组越界无法处理
		return false
	} else if h.len() == 1 {
		h.Values = []T{}
	} else {
		last := h.len() - 1
		h.Values[idx] = h.Values[last]
		h.Values = h.Values[:last] //最后一个值剪掉
		h.sinkDown(idx)
	}
	return true
}
func (h *Heap[T]) Pop() (value T, err error) {
	if h.len() == 0 {
		err = fmt.Errorf("已经没有值了")
		return
	}
	value = h.Values[0]
	h.Delete(0)
	return
}
func (h *Heap[T]) Deletes(idxs ...int) bool {
	for _, idx := range idxs {
		returnValue := h.Delete(idx)
		if !returnValue {
			return false
		}
	}
	return true
}
func NewHeap[T utils.Ord](model HeapType) *Heap[T] {
	return &Heap[T]{
		Values: []T{},
		model:  model,
	}
}
func BuildHeapSlow[T utils.Ord](model HeapType, items ...T) *Heap[T] {
	h := NewHeap[T](model)
	h.Inserts(items...)
	return h
}
func BuildHeapQuick[T utils.Ord](model HeapType, items ...T) *Heap[T] {
	h := NewHeap[T](model)                        //创建空堆
	h.Values = items                              //直接把值赋进去
	for i := h.father(h.len() - 1); i >= 0; i-- { //对所有非叶子结点sinkDown操作
		h.sinkDown(i)
	}
	return h
}
