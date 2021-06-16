//Main package to demonstrate Linked List Data Structure
package main

import "fmt"

//In order to group the nodes we use type - struct
type linkedList struct {
	head *node
	tail *node
}

//Each data element is stored in a node which is of type - struct
type node struct {
	next *node
	data interface{}
}

//New() : Function to initialise an instance of LinkedList struct
func New() *linkedList {
	emptyNode := &node {
		next: nil,
		data: nil,
	}
	return &linkedList {
		head: emptyNode,
		tail: emptyNode,
	}
}

//AppendNode() : Function to add a new Node to existing Linked List
func (ll *linkedList) AppendNode(d interface{}) *linkedList {
	nextNode := &node {
		next: nil,
		data: d,
	}
	if ll.head.data == nil {
		ll.head = nextNode
	} else {
		ll.tail.next = nextNode
	}
	ll.tail = nextNode
	return ll
}

//DeleteNodeWithValue() : Function to delete a node from an existing Linked List
func (ll *linkedList) DeleteNodeWithValue(v interface{}) *linkedList {
	var element =ll.head
	if element.data == v {
		ll.head = ll.head.next
		return ll
	}
	for {
		if v == element.next.data {
			if element.next.next != nil {
				element.next = element.next.next
				break
			}
			element.next = nil
			break
		}
		element = element.next
	}
	return ll
}

//PrintAllNodes(): Function to print all nodes.
func (ll *linkedList) PrintAllNodes() {
	var element = ll.head
	for {
		fmt.Println("Elements are: ",element.data)
		if element.next == nil {
			return 
		}
		element = element.next
	}
}

func main() {
	linkedlistObj := New()
	linkedlistObj.AppendNode(1).AppendNode(2).AppendNode(3).PrintAllNodes()
	linkedlistObj.PrintAllNodes()
	linkedlistObj.DeleteNodeWithValue(2)
	linkedlistObj.PrintAllNodes()
	linkedlistObj.DeleteNodeWithValue(1)
	linkedlistObj.PrintAllNodes()
}