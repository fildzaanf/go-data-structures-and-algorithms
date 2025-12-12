package main

import (
	"fmt"
)

func ValidTriangle(a, b, c int) string {

	if a+b <= c || a+c <= b || b+c <= a {
		return "invalid"
	}

	if a == b && b == c {
		return "valid (equilateral)"
	}

	if (c*c) == ((a*a)+(b*b)) || (b*b) == (c*c)+(a*a) || (a*a) == (b*b)+(c*c) {
		return "valid (right-angled)"
	}

	return "valid"

}

func main() {
	a := 3
	b := 4
	c := 5

	result := ValidTriangle(a, b, c)

	fmt.Print(result)

}
