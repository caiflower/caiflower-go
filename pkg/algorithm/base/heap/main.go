package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // 最小堆
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	// 创建一个堆
	h := &IntHeap{2, 1, 5}
	heap.Init(h) // 初始化堆

	// 推入元素
	heap.Push(h, 3)
	fmt.Printf("推入 3 后： %d\n", (*h)[0]) // 最小堆，堆顶为最小元素

	// 弹出元素
	fmt.Printf("弹出: %d\n", heap.Pop(h))
	fmt.Printf("弹出后堆顶: %d\n", (*h)[0])
}
