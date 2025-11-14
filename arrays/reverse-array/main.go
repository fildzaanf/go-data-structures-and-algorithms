package main

import "fmt"

func ReverseArray(array []int) []int {
	left := 0
	right := len(array) - 1

	for left < right {
		array[left], array[right] = array[right], array[left]
		left++
		right--
	}

	return array
}

func main() {
	array := []int{1, 4, 3, 2, 6, 5}

	result := ReverseArray(array)

	fmt.Print(result)
}