package main

import (
	"fmt"
	"unicode"
)

func IsPalindrome(s string) bool {
	r := []rune(s)

	left, right := 0, len(r)-1
	for left < right {

		if !unicode.IsLetter(r[left]) {
			left++
			continue
		}

		if !unicode.IsLetter(r[right]) {
			right--
			continue
		}

		if unicode.ToLower(r[left]) != unicode.ToLower(r[right]) {
			return false
		}

		left++
		right--
	}
	return true
}

func main() {
	fmt.Println(IsPalindrome("abba"))
	fmt.Println(IsPalindrome("AbBa"))
	fmt.Println(IsPalindrome("A man, a plan, a canal, Panama"))
	fmt.Println(IsPalindrome("abc"))
}
