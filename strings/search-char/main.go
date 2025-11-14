package main

import "fmt"

func SearchCharInString(ch byte, s string) int {
	for i := 0; i < len(s); i++ {
		if s[i] == ch {
			return i
		}
	}

	return -1
}

func main() {
	ch := byte('z')
	s := "geeksforgeeks"

	result := SearchCharInString(ch, s)

	fmt.Print(result)

}