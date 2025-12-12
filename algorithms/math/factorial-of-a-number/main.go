package main

import "fmt"

func FactorialNumber(n int) int {
	result := 1

	for i := 2; i <= n; i++ {
		result *= i

	}

	return result
}

func main() {
	n := 5

	result := FactorialNumber(n)

	fmt.Print(result)
}
