package main

import "fmt"

func IsBinary(s string) bool {
	for _, v := range s {
		if v != '0' && v != '1' {
			return false
		}
	}
	return true
}

func main() {
	s := "01010101010"

	result := IsBinary(s)

	fmt.Print(result)
}