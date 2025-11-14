package main

import "fmt"

// sorting algorithm with comparison based
func BubbleSort(array []int) {

	lenArray := len(array)

	for i := 0; i < lenArray-1; i++ {
		isSwapped := false

		for j := 0; j < lenArray-i-1; j++ {

			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				isSwapped = true
			}

		}

		if !isSwapped {
			break
		}

	}

}

func main() {

	array := []int{4, 7, 3, 2, 1, 9, 6, 8}

	BubbleSort(array)

	fmt.Print("sort array with bubble sort: ", array)

}
