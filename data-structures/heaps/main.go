package main

import "fmt"

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) Extract() int {
	extracted := h.array[0]

	// Check array length to avoid panic.
	if len(h.array) == 0 {
		fmt.Println("Array is zero length!")
		return -1
	}

	l := len(h.array) - 1
	// Put last value into root
	h.array[0] = h.array[l]
	// Decrease array
	h.array = h.array[:l]

	// Adjust the Heap
	h.maxHeapifyDown(0)

	return extracted
}

// We are going to heapify from bottom top
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// Will heapify from Top to Bottom
func (h *MaxHeap) maxHeapifyDown(index int) {

	lastIdx := len(h.array) - 1
	leftIdx, rightIdx := left(index), right(index)
	childToCompare := 0

	// Loop while index has at least one child
	for leftIdx <= lastIdx {
		// Case left child is the only child in tree path
		if leftIdx == lastIdx {
			childToCompare = leftIdx
		} else if h.array[leftIdx] > h.array[rightIdx] { // Case left child has larger value
			childToCompare = leftIdx
		} else {
			childToCompare = rightIdx // Case right child has larger value
		}

		// Do comparison between current index value with the larger child
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			leftIdx, rightIdx = left(index), right(index)
		} else {
			// we found the correct index value position.
			return
		}
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func (h *MaxHeap) swap(idx1, idx2 int) {
	h.array[idx1], h.array[idx2] = h.array[idx2], h.array[idx1]
}

func main() {

	mHeap := &MaxHeap{}
	fmt.Println(mHeap)
	// heap :=

	heap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	for _, v := range heap {
		mHeap.Insert(v)
		fmt.Println(mHeap)
	}

	for i := 0; i < 5; i++ {
		mHeap.Extract()
		fmt.Println(mHeap)
	}
}
