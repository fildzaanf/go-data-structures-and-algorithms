package main

import "fmt"

func ReverseArrayGroup(array []int32, k int32) []int32 {
	n := int32(len(array))

	for i := int32(0); i < n; i += k {
		left := i

		right := i + k - 1
		if right >= n {
			right = n - 1
		}

		for left < right {
			array[left], array[right] = array[right], array[left]
			left++
			right--
		}
	}

	return array
}

func main() {

	arr1 := []int32{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(ReverseArrayGroup(arr1, 3))

	arr2 := []int32{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(ReverseArrayGroup(arr2, 3))

	arr3 := []int32{1, 2, 3}
	fmt.Println(ReverseArrayGroup(arr3, 1))

	arr4 := []int32{1, 2, 3}
	fmt.Println(ReverseArrayGroup(arr4, 10))

	arr5 := []int32{}
	fmt.Println(ReverseArrayGroup(arr5, 3))

	arr6 := []int32{9}
	fmt.Println(ReverseArrayGroup(arr6, 5))

	arr7 := []int32{1, 2, 3, 4, 5, 6}
	fmt.Println(ReverseArrayGroup(arr7, 2))

	arr8 := []int32{1, 2, 3, 4, 5, 6}
	fmt.Println(ReverseArrayGroup(arr8, 4))
}
