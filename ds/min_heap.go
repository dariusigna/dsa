package ds

import "container/heap"

type IntMinHeap []int

func (h IntMinHeap) Len() int           { return len(h) }
func (h IntMinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMinHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntMinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func NewIntMinHeap() *IntMinHeap {
	h := &IntMinHeap{}
	heap.Init(h)
	return &IntMinHeap{}
}

func (h *IntMinHeap) Add(x int) {
	heap.Push(h, x)
}

func (h *IntMinHeap) Remove() int {
	return heap.Pop(h).(int)
}

func (h *IntMinHeap) Peek() int {
	return (*h)[0]
}

func (h *IntMinHeap) IsEmpty() bool {
	return h.Len() == 0
}
