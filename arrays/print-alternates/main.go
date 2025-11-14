package main

import "fmt"

func PrintAlternates(array []int) []int {

	var result []int

	for i := 0; i < len(array); i += 2 {

		result = append(result, array[i])

	}

	return result
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6}

	result := PrintAlternates(array)

	fmt.Print(result)
}
