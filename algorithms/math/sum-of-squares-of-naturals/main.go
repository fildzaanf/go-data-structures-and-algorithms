package main

import "fmt"

func SumOfSquaresOfNaturals(n int) int {
	sum := 0

	for i := 1; i <= n; i++ {
		square := i * i
		sum += square
	}

	return sum
}

func main() {
	n := 2

	result := SumOfSquaresOfNaturals(n)
	fmt.Print(result)
}