package main

import (
	"fmt"
)

type Stack struct {
	items []int
}

// Push function adds a value at the end off slice
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop function removes a value from the end and returns it.
func (s *Stack) Pop() int {
	lastIdx := len(s.items) - 1
	removedItem := s.items[lastIdx]
	s.items = s.items[:lastIdx]

	return removedItem
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)

	myStack.Push(10)
	myStack.Push(1)
	myStack.Push(20)

	fmt.Println(myStack)
}
