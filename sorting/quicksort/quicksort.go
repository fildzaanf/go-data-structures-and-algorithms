package main

import "fmt"

func Partition(array []int, low, high int) int {
	pivot := array[high]
	i := low - 1

	for j := low; j < high; j++ {
		if array[j] < pivot {
			i++
			array[i], array[j] = array[j], array[i]
		}
	}

	array[i+1], array[high] = array[high], array[i+1]

	return i + 1
}

func QuickSort(array []int, low, high int) {
	if low < high {
		pivotIndex := Partition(array, low, high)
		QuickSort(array, low, pivotIndex-1)
		QuickSort(array, pivotIndex+1, high)
	}
}

func main() {
	array := []int{64, 34, 25, 12, 22, 11, 90, 5}
	low := 0
	high := len(array) - 1

	QuickSort(array, low, high)

	fmt.Println("sorted array with quick sort: ", array)

}
