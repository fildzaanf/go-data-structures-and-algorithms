package main

import (
	"fmt"
)

func SumOfDigits(n int) int {
	sum := 0

	for n > 0 {
		last := n % 10

		sum = sum + last

		n = n / 10
	}

	return sum
}

func main() {
	n := 12345

	result := SumOfDigits(n)
	
	fmt.Print(result)
}
