package main

import "fmt"

func swap(a []int, i, j int) {
	if i == j {
		return
	}

	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}

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

	for lastUnsortedIdx := len(intArr) - 1; lastUnsortedIdx > 0; lastUnsortedIdx-- {
		for i := 0; i < lastUnsortedIdx; i++ {
			if intArr[i] > intArr[i+1] {
				swap(intArr, i, i+1)
			}
		}
	}

	fmt.Print("Sorted array: ")
	print(intArr)

	fmt.Println("Sorting done!")
}
