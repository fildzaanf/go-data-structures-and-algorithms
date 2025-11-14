package main

import "fmt"


func GreatestCommonDivisor(a, b int) int {
	
	factorsA := PrimeFactors(a)
	factorsB := PrimeFactors(b)

	commonFactors := FindCommonFactors(factorsA, factorsB)

	result := 1
	for _, factor := range commonFactors {
		result *= factor
	}

	return result
}

func PrimeFactors(n int) []int {
	var factors []int
	for i := 2; n > 1; i++ { 
		for n % i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}
	
	if n > 1 { 
		factors = append(factors, n)
	}

	return factors
}

func FindCommonFactors(a, b []int) []int {
	var common []int
	used := make([]bool, len(b)) 

	for _, fa := range a {
		for j, fb := range b {
			if !used[j] && fa == fb {
				common = append(common, fa)
				used[j] = true 
				break
			}
		}
	}
	return common
}

func main() {
	a := 12
	b := 18

	result := GreatestCommonDivisor(a, b)

	fmt.Print(result)
}