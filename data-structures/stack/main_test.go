package main

import (
  "testing"
	"github.com/stretchr/testify/assert"
)


func TestPush(t *testing.T) {
  myStack := Stack{}

  myStack.Push(10)
  myStack.Push(1)
  myStack.Push(20)

  expected := []int{10, 1, 20}
  assert.Equal(t, expected, myStack.items)
}
