package main

import "fmt"

func ReverseArray(array []int) []int {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
	return array
}

func main() {
	array := []int{1, 4, 3, 2, 6, 5}

	result := ReverseArray(array)

	fmt.Print(result)
}