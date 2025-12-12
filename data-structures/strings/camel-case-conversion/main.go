package main

import (
	"fmt"
	"strings"
)

func CamelCaseConversion(s string) string{
	words := strings.Split(s, " ")
	if len(words) == 0 {
		return ""
	}

	camel := words[0]

	for i := 1; i < len(words); i++ {
		w := words[i]

		first := strings.ToUpper(string(w[0]))
		camel += first + w[1:]
	}

	return camel
}

func CamelCaseConversionII(s string) string{
	sentence := [][]rune{}
	word := []rune{}

	for _, ch := range s {
		if ch == ' ' {
			if len(word) > 0 {
				sentence = append(sentence, word)
				word = []rune{}
			}
		} else {
			word = append(word, ch)
		}
	}

	if len(word) > 0 {
		sentence = append(sentence, word)
	}

	if len(sentence) == 0 {
		return ""
	}

	result := []rune{}

	result = append(result, sentence[0]...)

	for j := 1; j < len(sentence); j++ {
		w := sentence[j]

		if len(w) > 0 {
			first := w[0]

			if first >= 'a' && first <= 'z' {
				first -= 32
			}

			result = append(result, first)

			result = append(result, w[1:]...)

		}
	}

	return string(result)
}

func main() {
	fmt.Println(CamelCaseConversion("i got intern at geeksforgeeks"))
	fmt.Println(CamelCaseConversion("here comes the garden"))

	fmt.Println(CamelCaseConversionII("i got intern at geeksforgeeks"))
	fmt.Println(CamelCaseConversionII("here comes the garden"))
}
