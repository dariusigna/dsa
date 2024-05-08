package ds

import "container/heap"

type IntMaxHeap []int

func (h IntMaxHeap) Len() int           { return len(h) }
func (h IntMaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMaxHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntMaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func NewIntMaxHeap() *IntMaxHeap {
	h := &IntMaxHeap{}
	heap.Init(h)
	return &IntMaxHeap{}
}

func (h *IntMaxHeap) Add(x int) {
	heap.Push(h, x)
}

func (h *IntMaxHeap) Remove() int {
	return heap.Pop(h).(int)
}

func (h *IntMaxHeap) Peek() int {
	return (*h)[0]
}

func (h *IntMaxHeap) IsEmpty() bool {
	return h.Len() == 0
}
