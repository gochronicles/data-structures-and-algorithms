package main

import (
	"fmt"
	"time"
	"math/rand"
	"math"
	"strings"
)

type HeapNode struct {
	heap  *MinHeap
	index int
}

type MinHeap struct {
	arr []int
}

func NewMinHeap(values ...int) *MinHeap {
	h := &MinHeap{
		arr: append([]int{}, values...),
	}

	for i := len(h.arr)/2 - 1; i >= 0; i-- {
		h.down(i)
	}

	return h
}

func (h *MinHeap) Count() int {
	return len(h.arr)
}

func (h *MinHeap) Push(value int) {
	h.arr = append(h.arr, value)
	h.up(len(h.arr) - 1)
}

func (h *MinHeap) Pop() (int, bool) {
	if len(h.arr) == 0 {
		return 0, false
	}

	val := h.arr[0]
	h.arr[0], h.arr[len(h.arr)-1] = h.arr[len(h.arr)-1], h.arr[0]
	h.arr = h.arr[:len(h.arr)-1]
	h.down(0)

	return val, true
}

func (h *MinHeap) up(idx int) {
	for {
		parentIdx := (idx - 1) / 2
		if idx == 0 || h.arr[parentIdx] <= h.arr[idx] {
			break
		}
		h.arr[idx], h.arr[parentIdx] = h.arr[parentIdx], h.arr[idx]
		idx = parentIdx
	}
}
func (h *MinHeap) down(idx int) {
	for {
		// pick child to swap (smaller one)
		childIdx := idx*2 + 1                       // left child
		if childIdx >= len(h.arr) || childIdx < 0 { // <0 int overflow
			break
		}
		rightIdx := childIdx + 1
		if rightIdx < len(h.arr) && h.arr[childIdx] >= h.arr[rightIdx] {
			childIdx = rightIdx
		}

		// swap
		if h.arr[childIdx] >= h.arr[idx] {
			break
		}
		h.arr[idx], h.arr[childIdx] = h.arr[childIdx], h.arr[idx]
		idx = childIdx
	}
}


func (h *MinHeap) Root() *HeapNode {
	return h.nodeAt(0)
}

func (h *MinHeap) Height() int {
	return int(math.Floor(math.Log2(float64(len(h.arr)))))
}

func (h *MinHeap) PrintTree() {
	for i := 0; i < len(h.arr); i++ {
		node := h.nodeAt(i)

		// left-side whitespaces
		if leftChild := node.LeftChild(); leftChild != nil {
			if node.isRightChild() {
				fmt.Print(strings.Repeat("-", leftChild.calcPrintWidth()))
			} else {
				fmt.Print(strings.Repeat(" ", leftChild.calcPrintWidth()))
			}

		}

		// node value
		fmt.Printf(" %d ", node.Value())

		// right-side whitespaces
		if rightChild := node.RightChild(); rightChild != nil {
			if i == 0 || node.isRightChild() {
				fmt.Print(strings.Repeat(" ", rightChild.calcPrintWidth()))
			} else {
				fmt.Print(strings.Repeat("-", rightChild.calcPrintWidth()))
			}
		}

		if node.isRightMost() {
			if i == len(h.arr)-1 && !node.isRightChild() {
				if parent := node.Parent(); parent != nil {
					if vw := parent.calcValueWidth(); vw > 0 {
						fmt.Print(strings.Repeat("-", vw/2))
						fmt.Print("+")
					}
				}
			}
			fmt.Println()
		} else {
			if rightParent := node.findRightParent(); rightParent != nil {
				if vw := rightParent.calcValueWidth(); vw > 0 {
					if node.isRightChild() {
						fmt.Print(strings.Repeat(" ", vw))
					} else {
						fmt.Print(strings.Repeat("-", vw/2))
						fmt.Print("+")
						fmt.Print(strings.Repeat("-", (vw-1)/2))
					}
				}
			}
		}
	}
}

func (h *MinHeap) Validate() bool {
	if root := h.Root(); root != nil {
		return root.validate()
	}
	return true
}

func (h *MinHeap) nodeAt(index int) *HeapNode {
	if index < 0 || index >= len(h.arr) {
		return nil
	}

	return &HeapNode{heap: h, index: index}
}

func (n *HeapNode) Value() int {
	return n.heap.arr[n.index]
}

func (n *HeapNode) Parent() *HeapNode {
	return n.heap.nodeAt((n.index+1)/2 - 1)
}

func (n *HeapNode) LeftChild() *HeapNode {
	return n.heap.nodeAt((n.index+1)*2 - 1)
}

func (n *HeapNode) RightChild() *HeapNode {
	return n.heap.nodeAt((n.index + 1) * 2)
}

func (n *HeapNode) Depth() int {
	return int(math.Floor(math.Log2(float64(n.index + 1))))
}

func (n *HeapNode) Height() int {
	height := 0
	if n.LeftChild() != nil {
		leftHeight := n.LeftChild().Height()
		height = leftHeight + 1
	}
	if n.RightChild() != nil {
		rightHeight := n.RightChild().Height()
		if rightHeight+1 > height {
			height = rightHeight + 1
		}
	}
	return height
}

func (n *HeapNode) calcValueWidth() int {
	v := n.Value()
	if v == 0 {
		return 3
	}

	return int(math.Log10(float64(v))) + 3
}

func (n *HeapNode) calcPrintWidth() int {
	width := n.calcValueWidth()

	if n.LeftChild() != nil {
		width += n.LeftChild().calcPrintWidth()
	}
	if n.RightChild() != nil {
		width += n.RightChild().calcPrintWidth()
	}

	return width
}

func (n *HeapNode) isRightChild() bool {
	return n.index > 0 && n.index%2 == 0
}

func (n *HeapNode) isRightMost() bool {
	return n.index == len(n.heap.arr)-1 ||
		n.index == int(math.Exp2(float64(n.Depth()+1)))-2
}

func (n *HeapNode) findRightParent() *HeapNode {
	node := n
	for {
		if !node.isRightChild() {
			return node.Parent()
		}
		node = node.Parent()
	}
	return nil
}

func (n *HeapNode) validate() bool {
	if leftChild := n.LeftChild(); leftChild != nil {
		if leftChild.Value() < n.Value() || !leftChild.validate() {
			fmt.Printf("SELF: %d / LEFT: %d\n", n.Value(), leftChild.Value())
			return false
		}
	}
	if rightChild := n.RightChild(); rightChild != nil {
		if rightChild.Value() < n.Value() || !rightChild.validate() {
			fmt.Printf("SELF: %d / RIGHT: %d\n", n.Value(), rightChild.Value())
			return false
		}
	}
	return true
}


func main() {
	valMax := 100
	rand.Seed(time.Now().Unix())

	input := []int{}
	for i := 0; i < 20; i++ {
		input = append(input, rand.Intn(valMax))
	}
	heap := NewMinHeap(input...)

	for t := 0; t <= 100; t++ {
		heap.PrintTree()
		fmt.Println()
		time.Sleep(100 * time.Millisecond)

		if !heap.Validate() {
			panic("Invalid heap")
		}

		if rand.Int()%2 == 0 {
			val := rand.Intn(valMax)
			fmt.Printf("PUSH %d\n", val)
			heap.Push(val)
		} else {
			val, ok := heap.Pop()
			if ok {
				fmt.Printf("POP %d\n", val)
			}
		}

		fmt.Println()
	}
}