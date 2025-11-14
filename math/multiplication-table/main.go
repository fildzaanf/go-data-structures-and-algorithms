package main

import "fmt"

func MultiplicationTable(n, limit int) {
	for i := 1; i <= limit; i++ {
		fmt.Printf("%d * %d = %d\n", n, i, n*i)
	}
}

func main() {
	n := 8
	limit := 10

	MultiplicationTable(n, limit)
}
