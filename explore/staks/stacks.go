package main

import (
	"fmt"
)

// Stack represents a simple stack data structure
type Stack struct {
	items []int
}

// Push adds an element to the top of the stack
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes and returns the element from the top of the stack
func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		return -1 // Assuming -1 represents an error or empty stack
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// Peek returns the element at the top of the stack without removing it
func (s *Stack) Peek() int {
	if len(s.items) == 0 {
		return -1 // Assuming -1 represents an error or empty stack
	}
	return s.items[len(s.items)-1]
}

// IsEmpty returns true if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func main() {
	stack := Stack{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println("Top of the stack:", stack.Peek())

	fmt.Println("Pop:", stack.Pop())
	fmt.Println("Pop:", stack.Pop())

	fmt.Println("Is the stack empty?", stack.IsEmpty())
}
