package main

import "fmt"

// sorting algorithm with comparison based
func InsertionSort(array []int) {

	lenArray := len(array)

	for i := 1; i < lenArray; i++ {

		insertIndex := i
		currentValue := array[insertIndex]

		for j := i - 1; j >= 0; j-- {
			if array[j] > currentValue {
				array[j+1] = array[j]
				insertIndex = j
			} else {
				break
			}
		}

		array[insertIndex] = currentValue
	}

}

func main() {

	array := []int{9, 2, 8, 1, 3, 7, 4}

	InsertionSort(array)

	fmt.Print("sort array with insertion sort: ", array)
}