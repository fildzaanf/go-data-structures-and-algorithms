package main

import "fmt"

func SwapTwoNumbers(a, b int) (int, int) {
	return b, a
}

func main() {
	a := 7
	b := 9

	a, b = SwapTwoNumbers(a, b)
	fmt.Printf("a = %d \nb = %d", a, b)
}
