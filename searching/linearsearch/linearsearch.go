package main

import "fmt"

func LinearSearch(array []int, targetValue int) int {
	len := len(array)

	for i := 0; i < len; i++ {
		if array[i] == targetValue {
			return i
		}
	}

	return -1
}

func main() {

	array := []int{3, 4, 2, 7, 9, 5}
	targetValue := 7

	result := LinearSearch(array, targetValue)

	if result != -1 {
		fmt.Printf("value %d found at index %d\n", targetValue, result)
	} else {
		fmt.Printf("value %d not found\n", targetValue)
	}

}
