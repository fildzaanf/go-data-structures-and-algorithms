package main

import "fmt"

func SumOfNaturals(n int) int {
    return n * (n + 1) / 2
}

func main() {
	n := 3

	result := SumOfNaturals(n)
	fmt.Print(result)
}