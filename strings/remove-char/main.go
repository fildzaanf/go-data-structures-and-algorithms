package main

import "fmt"

func RemoveCharAtPos(s string, pos int) string {
	if pos < 0 || pos > len(s) {
		return "invalid pos"
	}

	return s[:pos] + s[pos+1:]
}

func main() {
	s := "Geeks"
	pos := 2

	result := RemoveCharAtPos(s, pos)
	fmt.Println(result)
}
