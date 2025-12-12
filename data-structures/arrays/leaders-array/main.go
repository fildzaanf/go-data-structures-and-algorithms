package main

import "fmt"

func LeadersArray(array []int) []int {
	var result []int

	for i := 0; i < len(array); i++ {
		isLeader := true

		for j := i+1; j < len(array); j++ {
			if array[j] > array[i] {
				isLeader = false
				break
			}
		}

		if isLeader {
			result = append(result, array[i])
		}
	}

	return result
}

func main() {
	array := []int{16, 17, 4, 3, 5, 2}

	result := LeadersArray(array)

	fmt.Print(result)
}
