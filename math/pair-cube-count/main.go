package main

import "fmt"

func PairCubeCount(n int) int {
	count := 0

	for a := 1; a*a*a <= n; a++ {
		for b := a; b*b*b <= n; b++ {
			if a*a*a+b*b*b == n {
				fmt.Printf("found pair: a=%d, b=%d → %d³ + %d³ = %d\n", a, b, a, b, n)
				count++
			}
		}
	}

	return count
}

func main() {
	n := 1729

	result := PairCubeCount(n)

	fmt.Printf("total: %d", result)
}