/*
LINKED LISTS

Example 1: Given the head of a linked list with an odd number
of nodes head, return the value of the node in the middle.

For example, given a linked list that
represents 1 -> 2 -> 3 -> 4 -> 5, return 3.
*/

package main

import "fmt"

// Node represents a node in the linked list
type Node struct {
	data int
	next *Node
}

// LinkedList represents a singly linked list
type LinkedList struct {
	head *Node
}

// Append adds a new node with the given data to the end of the linked list
func (ll *LinkedList) Append(data int) {
	newNode := &Node{data: data}

	if ll.head == nil {
		ll.head = newNode
		return
	}

	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// FindMiddle returns the value of the middle node in the linked list
func (ll *LinkedList) FindMiddle() int {
	slow, fast := ll.head, ll.head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	return slow.data
}

func main() {
	linkedList := LinkedList{}
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)
	linkedList.Append(4)
	linkedList.Append(5)

	middleValue := linkedList.FindMiddle()
	fmt.Println("Middle Value:", middleValue)
}
