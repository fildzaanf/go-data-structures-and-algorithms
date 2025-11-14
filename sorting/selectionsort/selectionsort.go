package main

import (
	"fmt"
)

func SelectionSort(array []int) {

	lenArray := len(array)

	for i := 0; i < lenArray-1; i++ {
		minIndex := i
		for j := i + 1; j < lenArray; j++ {
			if array[j] < array[minIndex] {
				minIndex = j
			}
		}

		array[i], array[minIndex] = array[minIndex], array[i]
	}

}

func main() {

	array := []int{9, 3, 4, 2, 8, 7, 1}

	SelectionSort(array)

	fmt.Print("sorted array with selection sort: ", array)
}
