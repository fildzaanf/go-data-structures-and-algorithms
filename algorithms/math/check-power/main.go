package main

import "fmt"

func CheckPower(x, y int) bool {

	if y == 1 {
		return true
	}

	if x == 1 || x == 0 {
		return y == 1
	}

	for y % x == 0 {
		y = y / x
	}

	return y == 1
}

func main() {
	x := 2
	y := 16

	result := CheckPower(x, y)
	fmt.Print(result)
}