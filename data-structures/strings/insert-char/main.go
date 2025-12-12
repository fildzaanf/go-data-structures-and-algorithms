package main

import "fmt"

func InsertCharAtPos(s, char string, pos int) string {

	if pos < 0 || pos > len(s) {
		return "invalid pos"
	}

	return s[:pos] + char + s[pos:]
}

func main() {
	s := "Geeks"
	c := "A"
	pos := 1

	result := InsertCharAtPos(s, c, pos)
	fmt.Println(result) 
}
