package main

import "fmt"

func print(a []int) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d ", a[i])
	}
	fmt.Println(" ")
}

func main() {
	fmt.Println("Staring sorting...")

	intArr := []int{
		20, 35, -15, 7, 55, 1, -22,
	}

	fmt.Print("Initial array: ")
	print(intArr)

	for firstUnsortedIdx := 1; firstUnsortedIdx < len(intArr); firstUnsortedIdx++ {
		newElement := intArr[firstUnsortedIdx]

		i := firstUnsortedIdx
		for ; i > 0; i-- {
			if intArr[i-1] > newElement {
				intArr[i] = intArr[i-1]
			}
		}
		intArr[i] = newElement
	}

	fmt.Print("Sorted array: ")
	print(intArr)

	fmt.Println("Sorting done!")
}
