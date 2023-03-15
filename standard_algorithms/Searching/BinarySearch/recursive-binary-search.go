package main

import "fmt"

func main() {
	nums := []int{2, 3, 5, 7, 8, 9, 10, 13, 15, 17, 19, 22, 50, 101}
	targetValue := 7
	low := 0
	high := len(nums) - 1

	idx := binarySearch(nums, low, high, targetValue)
	if idx == -1 {
		fmt.Println("Element is not found in array")
	} else {
		fmt.Println("Element found at index ", idx)
	}
}

func binarySearch(numbers []int, low, high, target int) int {
	if low > high {
		return -1
	}

	// we can have an overflow here.
	middle := (low + high) / 2

	if target == numbers[middle] {
		return middle
	}

	if target < numbers[middle] {
		return binarySearch(numbers, low, middle-1, target)
	}

	return binarySearch(numbers, middle+1, high, target)
}
