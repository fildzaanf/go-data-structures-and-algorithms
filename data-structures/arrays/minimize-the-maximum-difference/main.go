package main

import "fmt"

func MinimizeMaximumDiff(arr []int, k int) int {
	n := len(arr)
	if n <= 1 {
		return 0
	}

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	diff := arr[n-1] - arr[0]

	small := arr[0] + k 
	big := arr[n-1] - k

	if small > big {
		small, big = big, small
	}

	for i := 1; i < n-1; i++ {
		sub := arr[i] - k 
		add := arr[i] + k 

		if sub < 0 { 
			continue
		}

		if sub >= small && add <= big { 
			continue
		}

		if add-big < small-sub { 
			big = add 
		} else {
			small = sub 
		}
	}

	if big-small < diff { 
		diff = big - small 
	}

	return diff
}

func main() {
	arr := []int{1, 5, 8, 10}
	k := 2 
	fmt.Println(MinimizeMaximumDiff(arr, k))
	fmt.Println(arr)
}
