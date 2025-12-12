package main

import (
	"fmt"
	"strings"
)

func IsSubstrings(s, sub string) bool {
	for i := 0; i <= len(s)-len(sub); i++ {
		if strings.HasPrefix(s[i:], sub) {
			return true
		}
	}
	return false
}

func IsSubstring(s, sub string) bool {
	n, m := len(s), len(sub)
	if m == 0 {
		return true
	}
	if m > n {
		return false
	}

	for i := 0; i <= n-m; i++ {
		match := true
		for j := 0; j < m; j++ {
			if s[i+j] != sub[j] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}

	return false
}

func main() {
	s := "geeksforgeeks"
	sub := "eks"

	fmt.Print(IsSubstring(s, sub))
}
