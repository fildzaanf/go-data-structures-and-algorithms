package main

import (
	"fmt"
)

func RemoveAllOccurences(c rune, s string) string {
	var result []rune

	for _, v := range s{
		if v != c {
			result = append(result, v)
		}
	}

	return string(result)
}

func main() {
	ch := '😊'
	s := "H😊e😊l😊l😊o"


	result := RemoveAllOccurences(ch, s)

	fmt.Print(result)

}
