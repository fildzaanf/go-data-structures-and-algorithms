package main

import "fmt"

func GenerateSubstrings(s string) []string {
	var substrings []string

	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
            substrings = append(substrings, s[i:j])
        }
	}

	return substrings

}

func main() {
	s := "abc"

	substrings := GenerateSubstrings(s)

	fmt.Print(substrings)
}
