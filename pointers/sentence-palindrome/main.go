package main

import "fmt"

func IsPalindromeASCII(s string) bool {
	clean := []byte{}

	for i := 0; i < len(s); i++ {
		ch := s[i]

		if 'A' <= ch && ch <= 'Z' {
			ch += 32
		}

		if ('a' <= ch && ch <= 'z') || ('0' <= ch && ch <= '9') {
			clean = append(clean, ch)
		}
	}

	left, right := 0, len(clean)-1
	for left < right {
		if clean[left] != clean[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func main() {
	fmt.Println(IsPalindromeASCII("Too hot to hoot."))   // true
	fmt.Println(IsPalindromeASCII("Abc 012..##  10cbA")) // true
	fmt.Println(IsPalindromeASCII("ABC $. def01ASDF..")) // false
}
