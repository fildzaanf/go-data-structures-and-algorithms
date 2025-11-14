package main

import "fmt"

func NthTermAP(a1, a2, n int) int {
	d := a2 - a1
	an := a1 + ((n - 1) * d)

	return an
}

func main() {
	a1 := 2
	a2 := 3
	n := 4

	result := NthTermAP(a1, a2, n)

	fmt.Print(result)

}
