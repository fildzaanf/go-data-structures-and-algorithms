package main

import "fmt"


func main() {

	matrix := [3][3]int{
		{1, 2, 3}, // row 0
		{4, 5, 6}, // row 1
		{7, 8, 9}, // row 2
	}

	value := matrix[1][2] // row = 1; col = 2; v = 6

	fmt.Print(value)
}
