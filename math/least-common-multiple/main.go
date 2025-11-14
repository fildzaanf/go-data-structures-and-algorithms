package main

import "fmt"

func LeastCommonMultiple(a, b int) int {

	factorsA := PrimeFactors(a)
	factorsB := PrimeFactors(b)

	commonFactors := FindCommonFactors(factorsA, factorsB)

	return commonFactors

}

func FindCommonFactors(a, b []int) int {
	var common int

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				return common
			}
		}
	}

	return common
}

func PrimeFactors(n int) []int {
	var factors []int
	for i := 2; n > 1; i++ {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}

	if n > 1 {
		factors = append(factors, n)
	}

	return factors
}

func main() {
	a := 10
	b := 5

	result := LeastCommonMultiple(a, b)

	fmt.Print(result)
}
