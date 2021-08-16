package main

import "fmt"

func main() {
	var array = []int{4, 8, 9, 1, 2, 5, 7, 6, 3, 0}
	var heap = new(Heap)
	fmt.Println(array)
	heap.HeapSort(array)
	fmt.Println(array)
}

type Heap struct {

}

func (heap *Heap) HeapSort(array []int) {
	heap.BuildMaxHeap(array)
	for length:= len(array); length > 1; length-- {
		heap.Pop(array, length)
	}
}

func (heap *Heap) BuildMaxHeap(array []int) {
	for i := len(array) / 2; i >= 0; i-- {
		heap.Heapify(array, i, len(array))
	}
}

func (heap *Heap) Pop(array []int, length int) {
	var lastIndex = length - 1
	array[0], array[lastIndex] = array[lastIndex], array[0]
	heap.Heapify(array, 0, lastIndex)
}

func (heap *Heap) Heapify(array []int, root, length int) {
	var max = root
	var l, r = heap.LeftChild(array, root), heap.RightChild(array, root)
	if l < length && array[l] > array[max] {
		max = l
	}
	if r < length && array[r] > array[max] {
		max = r
	}
	if max != root {
		array[root], array[max] = array[max], array[root]
		heap.Heapify(array, max, length)
	}
}

func (*Heap) LeftChild(array []int, root int) int {
	return (root * 2) + 1
}

func (*Heap) RightChild(array []int, root int) int {
	return (root * 2) + 2
}