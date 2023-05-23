package main

import (
	"fmt"
)

type minHeap struct {
	heap []int
}

func main() {
	heap := minHeap{}
	a := []int{33, 12, 56, 78, 11, 13, 24}
	for _,v:=range a{

		heap.insert(v)
	}
	fmt.Println(heap.heap)

}
func (h *minHeap) insert(data int) {
	h.heap = append(h.heap, data)
	h.insertHelper(len(h.heap) - 1)
}
func (h *minHeap) insertHelper(position int) {
	if h.heap[position] < h.heap[parent(position)] {
		swap(h.heap, position, parent(position))
		h.insertHelper(parent(position))
	}
}

func leftChild(parent int) int {
	return (2 * parent) + 1
}
func rightChild(parent int) int {
	return (2 * parent) + 2
}
func parent(child int) int {
	return (child - 1) / 2
}

func swap(a []int, x, y int) {
	a[x], a[y] = a[y], a[x]
}
//func Heapify()
