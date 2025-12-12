package main

import "fmt"

func GenerateAllSubarray(array []int) [][]int {
	var result [][]int

	for start := 0; start < len(array); start++{
		for end := start; end < len(array); end++{
			subarray := make([]int, end-start+1)
			copy(subarray, array[start:end+1])
			result = append(result, subarray)
		}
	}
	return result
}

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(GenerateAllSubarray(arr))
}
