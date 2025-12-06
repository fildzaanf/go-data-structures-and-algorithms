package main

import "fmt"

func NextPermutation(array []int) []int {
	n := len(array)

	pivot := -1
	for i := n - 2; i >= 0; i-- {
		if array[i] < array[i+1] {
			pivot = i
			break
		}
	}

	if pivot == -1 {
		for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
			array[l], array[r] = array[r], array[l]
		}
		return array
	}

	for j := n - 1; j > pivot; j-- {
		if array[j] > array[pivot] {
			array[j], array[pivot] = array[pivot], array[j]
			break
		}
	}

	for l, r := pivot+1, n-1; l < r; l, r = l+1, r-1 {
		array[l], array[r] = array[r], array[l]
	}

	return array
}

func main() {
	array := []int{2, 4, 1, 7, 5, 0}

	result := NextPermutation(array)

	fmt.Print(result)
}
