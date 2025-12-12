package main

import "fmt"

func BinarySearch(array []int, targetValue int) int {
	left := 0
	right := len(array)-1

	for left <= right {
		mid := (left + right) / 2

		if array[mid] == targetValue {
			return mid
		} 
		
		if array[mid] < targetValue {
			left = mid + 1
		} else {
			right = mid - 1 
		}
	}

	return -1
}

func main() {
	array := []int{1, 2, 2, 4, 5, 6, 7, 9, 9}
	targetValue := 7

	result := BinarySearch(array, targetValue)

	if result != -1 {
		fmt.Printf("value %d found at index %d\n", targetValue, result)
	} else {
		fmt.Printf("value %d not found\n", targetValue)
	}
}
