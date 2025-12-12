package main

import "fmt"

func RemoveDuplicates(array []int) []int {
	if len(array) == 0 {
		return []int{}
	}

	result := []int{array[0]}

	for i := 1; i < len(array); i++{
		if array[i] != array[i-1]{
			result =  append(result, array[i])
		}
	}

	return result
}

func main() {
	array := []int{1, 2, 2, 3, 4, 4, 4, 5, 5}

	result := RemoveDuplicates(array)

	fmt.Print(result)
}
