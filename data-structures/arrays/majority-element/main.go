package main

import "fmt"

// n/2
func MooreVotingAlgorithms(array []int) int {
	n := len(array)

	if n == 0 {
		return -1
	}

	count := 0
	candidate := 0
	for _, value := range array {
		if count == 0 {
			candidate = value
			count = 1
		} else if value == candidate {
			count++
		} else {
			count--
		}
	}

	count = 0
	for _, value := range array {
		if value == candidate {
			count++
		}
	}

	if count > n/2 {
		return candidate
	}

	return -1
}

// n/3
func MooreVotingAlgorithmsII(array []int) []int {
	var result []int
	n := len(array)

	if n == 0 {
		return result
	}

	countC1, countC2 := 0, 0
	candidate1, candidate2 := 0, 0

	for _, value := range array {
		if value == candidate1 {
			countC1++
		} else if value == candidate2 {
			countC2++
		} else if countC1 == 0 {
			candidate1 = value
			countC1 = 1
		} else if countC2 == 0 {
			candidate2 = value
			countC2 = 1
		} else {
			countC1--
			countC2--
		}
	}

	countC1, countC2 = 0, 0
	for _, value := range array {
		if value == candidate1 {
			countC1++
		} else if value == candidate2 {
			countC2++
		}
	}

	if countC1 > n/3 {
		result = append(result, candidate1)
	}

	if countC2 > n/3 {
		result = append(result, candidate2)
	}

	if len(result) == 2 {
		if result[0] > result[1] {
			result[1], result[0] = result[0], result[1]
		}
	}

	return result

}

func main() {
	array := []int{2, 2, 3, 1, 3, 2, 1, 1}

	result := MooreVotingAlgorithmsII(array)

	fmt.Print(result)
}
