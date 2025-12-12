package main

import "fmt"
func CommonElements(first, second, third []int) []int {
	count :=  make(map[int]int)
	result := []int{}

	for _, v := range first {
		count[v]++
	}

	for _, v := range second {
		count[v]++
	}

	for _, v := range third {
		count[v]++
	}

	for key, value := range count {
		if value == 3 {
			result = append(result, key)
		}
	}

	return result
}

func main() {
	first := []int{1, 5, 10, 20, 30}
	second := []int{5, 13, 15, 20}
	third := []int{5, 20}

	result := CommonElements(first, second, third)
	fmt.Println(result)
}
