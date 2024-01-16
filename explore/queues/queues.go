package main

import (
	"fmt"
)

// Queue represents a simple queue data structure
type Queue struct {
	items []int
}

// Enqueue adds an element to the back of the queue
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the element from the front of the queue
func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		return -1 // Assuming -1 represents an error or empty queue
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// Front returns the element at the front of the queue without removing it
func (q *Queue) Front() int {
	if len(q.items) == 0 {
		return -1 // Assuming -1 represents an error or empty queue
	}
	return q.items[0]
}

// IsEmpty returns true if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func main() {
	queue := Queue{}

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	fmt.Println("Front of the queue:", queue.Front())

	fmt.Println("Dequeue:", queue.Dequeue())
	fmt.Println("Dequeue:", queue.Dequeue())

	fmt.Println("Is the queue empty?", queue.IsEmpty())
}
