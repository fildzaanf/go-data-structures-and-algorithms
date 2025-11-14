package main

import "fmt"

func CountingSort(array []int) []int {

	if len(array) == 0 {
		return array
	}

	maximumValue := array[0]

	for _, value := range array {
		if value > maximumValue {
			maximumValue = value
		}
	}

	count := make([]int, maximumValue+1)

	for _, value := range array {
		count[value]++
	}

	sortedArray := []int{}

	for value, frequency := range count {
		for i := 0; i < frequency; i++ {
			sortedArray = append(sortedArray, value)
		}
	}

	return sortedArray
}

func main() {

	array := []int{4, 4, 2, 1, 8, 1, 9, 5}

	CountingSort(array)

	fmt.Println("sorted array with counting sort: ", array)

}
