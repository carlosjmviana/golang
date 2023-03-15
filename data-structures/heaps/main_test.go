package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
  maxHeap := MaxHeap{}

  maxHeap.Insert(10)
  maxHeap.Insert(20)
  maxHeap.Insert(30)
 
  expected := []int{30, 10, 20}
  assert.Equal(t, expected, maxHeap.array)
}

func TestExtract(t *testing.T) {
  maxHeap := MaxHeap{}

  maxHeap.Insert(10)
  maxHeap.Insert(20)
  maxHeap.Insert(30)

  maxHeap.Extract()

  expected := []int{20, 10}
  assert.Equal(t, expected, maxHeap.array)
}
