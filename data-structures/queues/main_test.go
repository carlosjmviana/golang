package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnqueue(t *testing.T) {
	myQueue := Queue{}
	myQueue.Enqueue(10)
	myQueue.Enqueue(1)
	myQueue.Enqueue(20)

	expected := []int{10, 1, 20}
	assert.Equal(t, expected, myQueue.items)
}

func TestDequeue(t *testing.T) {
	myQueue := Queue{}
	myQueue.Enqueue(10)
	myQueue.Enqueue(1)
	myQueue.Enqueue(20)

	myQueue.Dequeue()

	expected := []int{1, 20}
	assert.Equal(t, expected, myQueue.items)
}
