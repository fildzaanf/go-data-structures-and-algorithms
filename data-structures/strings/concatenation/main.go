package main

import "fmt"


func Concatenation(s1, s2 string) string {
	return s1 + s2
}

func main() {
	s1 := "Hello"
	s2 := "Fildza"

	result := Concatenation(s1, s2)

	fmt.Print(result)
}
