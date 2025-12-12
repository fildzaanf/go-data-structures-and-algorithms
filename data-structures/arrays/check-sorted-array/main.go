package main

import "fmt"

func CheckSortedArray(array []int) bool {
	if len(array) <= 1 {
		return true
	}

	for i := 0; i < len(array)-1; i++ {
		if array[i] > array[i+1] {
			return false
		}
	}

	return true
}

func main() {
	array := []int{10, 20, 20, 30, 50}

	result := CheckSortedArray(array)

	fmt.Print(result)
}
