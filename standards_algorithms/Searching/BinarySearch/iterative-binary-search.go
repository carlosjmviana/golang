package main

import "fmt"

func main() {
	nums := []int{2, 3, 5, 7, 8, 9, 10, 13, 15, 17, 19, 22, 50, 101}
	targetValue := 7

	idx := binarySearch(nums, targetValue)
	if idx == -1 {
		fmt.Println("Element is not found in array")
	} else {
		fmt.Println("Element found at index ", idx)
	}
}

func binarySearch(numbers []int, target int) int {
	low := 0
	high := len(numbers) - 1

	for low <= high {
		// we can have an overflow here.
		middle := (low + high) / 2

		if target == numbers[middle] {
			return middle
		}
		if target < numbers[middle] {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}
	return -1
}
