//Main Package to demostrate Queue data structure
package main

import "fmt"

// In order to group our data we use type - struct
type queue struct {
	data []int
}

// Function for intializing an instance of Queue struct
func New() *queue {
	return &queue{
		data: []int{},
	}
}

// Function to check if the queue data structure is empty
func (q *queue) IsEmpty() bool {
	return len(q.data) == 0
}

// Peek : Return the first Element in the queue as per FIFO logic
func (q *queue) Peek() (int, error) {
	if len(q.data) == 0 {
		return 0, fmt.Errorf("Queue is empty.Please add elements")
	}
	return q.data[0], nil
}

// Enqueue : Function to add elements into the Queue
func (q *queue) Enqueue(n int) *queue {
	q.data = append(q.data, n)
	return q
}

// Dequeue : Function to remove elements from the Queue
func (q *queue) Dequeue() (int, error) {
	if len(q.data) == 0 {
		return 0, fmt.Errorf("Queue is empty.Please add elements")
	}
	element := q.data[0]
	q.data = q.data[1:]
	return element, nil
}

// Main Function
func main() {
	queue := New()
	result, _ := queue.Enqueue(1).Enqueue(2).Enqueue(3).Peek()
	fmt.Println(result)
	fmt.Println(queue.IsEmpty())
	result, _ = queue.Dequeue()
	fmt.Println(result)
	result, _ = queue.Dequeue()
	fmt.Println(result)
	queue.Dequeue()
	fmt.Println(queue.IsEmpty())
	_, err := queue.Peek()
	fmt.Println(err)
}