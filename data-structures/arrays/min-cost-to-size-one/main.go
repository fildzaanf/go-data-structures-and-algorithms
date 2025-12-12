package main

import "fmt"

func MinCostToSizeOne(array []int) int {
	if len(array) <= 1 {
		return 0
	}

	minValue := array[0]

	for i := 1; i < len(array); i++ {
		if array[i] < minValue {
			minValue = array[i]
		}
	}

	return minValue * (len(array) - 1)

}

func main() {
	array := []int{4, 4, 2, 1, 8, 1, 9, 5}

	result := MinCostToSizeOne(array)

	fmt.Print(result)

}
