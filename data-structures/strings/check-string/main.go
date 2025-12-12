package main

import "fmt"

func CheckString(s1, s2 string) bool {
	if len(s1) != len(s2) {
        return false
    }

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

func IsSameString(s1, s2 string) bool{
	return s1 == s2

}

func main() {
	s1 := "abcD"
	s2 := "abcd"

	result := CheckString(s1, s2)

	fmt.Print(result)
}
