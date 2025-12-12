package main

import (
	"fmt"
	"strings"
	"unicode"

)

func IsPanagram(s string) bool {
	alphabet := make(map[rune]bool)
	lowercase := strings.ToLower(s)

	for _, ch := range lowercase {
		if unicode.IsLetter(ch) {
			alphabet[ch] = true
		}
	}

	return len(alphabet) == 26
}


func main() {
	s := "The quick brown fox jumps over the lazy dog"

	fmt.Print(IsPanagram(s))
}
