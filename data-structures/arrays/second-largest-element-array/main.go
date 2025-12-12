package main

import "fmt"

func SecondLargestElement(array []int32) int32 {
	if len(array) < 2 {
		return int32(-1)
	}

	firstLargest := int32(-1)
	secondLargest := int32(-1)

	for _, v := range array {
		if v > firstLargest {
			secondLargest = firstLargest
			firstLargest = v
		} else if v > secondLargest && v < firstLargest {
			secondLargest = v
		}
	}

	return secondLargest
}

func main() {
	array := []int32{5, 10, 10, 10}

	result := SecondLargestElement(array)

	fmt.Print(result)
}
