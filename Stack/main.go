package main

import (
	"container/list"
	"fmt"
)

type dataStack struct {
	stack *list.List
}

func (d *dataStack) Push(value string) {
	d.stack.PushFront(value)
}

func (d *dataStack) Pop() error {
	if d.stack.Len() > 0 {
		element := d.stack.Front()
		d.stack.Remove(element)
	}
	return fmt.Errorf("Pop Error: Stack is Empty")
}

func (d *dataStack) Peek() (string, error) {
	if d.stack.Len() > 0 {
		if value, ok := d.stack.Front().Value.(string); ok {
			return value, nil
		}
		return "", fmt.Errorf("Peek Error: Stack data type is incorrect")
	}
	return "", fmt.Errorf("Peek Error: Stack is empty")
}

func (d *dataStack) Size() int {
	return d.stack.Len()
}

func main() {
	dataObject := &dataStack{
		stack: list.New(),
	}
	fmt.Printf("Push: Data-1\n")
	dataObject.Push("Data-1")
	fmt.Printf("Push: Data-2\n")
	dataObject.Push("Data-2")
	fmt.Printf("Size: %d\n", dataObject.Size())
	for dataObject.Size() > 0 {
		topVal, _ := dataObject.Peek()
		fmt.Printf("Front: %s\n", topVal)
		fmt.Printf("Pop: %s\n", topVal)
		dataObject.Pop()
	}
	fmt.Printf("Size: %d\n", dataObject.Size())
}
