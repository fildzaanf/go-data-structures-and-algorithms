package main

import "fmt"

func MissingNumber(number []int) int {
	n := len(number) + 1

	sum := 0
	for _, num := range number {
		sum += num
	}

	expsum := 0
	for i := 1; i <= n; i++ {
		expsum += i
	}

	return expsum - sum

}

func main() {
	number := []int{1, 2, 3, 5}
	fmt.Println(MissingNumber(number))
}
