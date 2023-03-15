package main

import (
	"fmt"
)

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

func (q *Queue) Dequeue() int {
	removedItem := q.items[0]
	q.items = q.items[1:]

	return removedItem
}

func main() {
	myQueue := Queue{}
	myQueue.Enqueue(10)
	myQueue.Enqueue(1)
	myQueue.Enqueue(20)

	fmt.Printf("Queue\n")
	fmt.Println(myQueue)

	fmt.Printf("\nRemoving item")
	removedItem := myQueue.Dequeue()
	fmt.Printf("\nItem removed: %d\n", removedItem)
	fmt.Println(myQueue)
}
