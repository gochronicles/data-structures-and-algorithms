//Main Package to demostrate Queue data structure
package main

import (
	"container/list"
	"fmt"
)

// In order to group our data we use type - struct
type dataStack struct {
	stack *list.List
}

// Push : Function to add elements into the stack
func (d *dataStack) Push(value string) {
	d.stack.PushFront(value)
}

// Pop : Function to remove elements from the stack
func (d *dataStack) Pop() error {
	if d.stack.Len() > 0 {
		element := d.stack.Front()
		d.stack.Remove(element)
	}
	return fmt.Errorf("Pop Error: Stack is Empty")
}

// Peek : Return the last Element in the state as per LIFO logic
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

func (d *dataStack) IsEmpty() bool {
	return d.stack.Len() == 0
}

func main() {
	dataObject := &dataStack{
		stack: list.New(),
	}
	fmt.Printf("Push: Data-1\n")
	dataObject.Push("Data-1")
	fmt.Printf("Push: Data-2\n")
	dataObject.Push("Data-2")
	fmt.Printf("Is Empty: %t\n", dataObject.IsEmpty())
	fmt.Printf("Size: %d\n", dataObject.Size())
	for dataObject.Size() > 0 {
		topVal, _ := dataObject.Peek()
		fmt.Printf("Front: %s\n", topVal)
		fmt.Printf("Pop: %s\n", topVal)
		dataObject.Pop()
	}
	fmt.Printf("Is Empty: %t\n", dataObject.IsEmpty())
	fmt.Printf("Size: %d\n", dataObject.Size())
}
