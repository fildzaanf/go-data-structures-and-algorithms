package main

import "fmt"

func LengthOfString(s string) int {
	count := 0

	for range s {
		count++
	}

	return count
}

func main() {
	s := "GeeksforGeeks"
	result := LengthOfString(s)

    fmt.Printf("built-in: %d\n", len(s))
	fmt.Printf("manual: %d\n", result) 
}
