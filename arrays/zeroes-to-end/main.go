package main

import "fmt"

func ZeroesToEnd(array []int) []int {
	pos := 0

	for i := 0; i < len(array); i++ {
		if array[i] != 0 {
			array[i], array[pos] = array[pos], array[i]
			pos++
		}
	}

	return array
}

func main() {
	array := []int{1, 2, 0, 4, 3, 0, 5, 0}

	result := ZeroesToEnd(array)

	fmt.Print(result)
}
