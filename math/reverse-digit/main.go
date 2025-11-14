package main

import "fmt"

func ReverseDigit(n int) int {

	result := 0

	for n > 0 {
		last := n % 10

		result = (result * 10) + last

		n = n / 10
	}

	return result
}

func main() {
	n := 12345

	result := ReverseDigit(n)

	fmt.Print(result)
}