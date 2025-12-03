package main

import "fmt"

func ThirdLargestElement(array []int32) int32 {
	if len(array) < 3 {
		return -1
	}

	first := int32(-1)
	second := int32(-1)
	third := int32(-1)

	for _, v := range array {
		if v > first { 
			third = second 
			second = first 
			first = v         
		} else if v > second && v < first { 
			third = second 
			second = v 
		} else if v < second && v > third { 
			third = v
		}
	}

	return third
}

func main() {
	array := []int32{19, -10, 20, 14, 2, 16, 10}

	result := ThirdLargestElement(array)

	fmt.Print(result)
}
