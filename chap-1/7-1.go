package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	text := "abcdefg 456789"
	fmt.Println(text)
	fmt.Println(reverseString1(text))
	fmt.Println(reverseString2(text))
	fmt.Println(reverseWord(text))
}

func reverseString1(s string) string {
	reversed := make([]rune, utf8.RuneCountInString(s))
	i := len(reversed) - 1

	for _, c := range s {
		reversed[i] = c
		i--
	}

	return string(reversed)
}

func reverseString2(s string) string {
	reversed := []rune(s)
	for i, j := 0, len(reversed)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}

	return string(reversed)
}

func reverseWord(s string) string {
	words := strings.Split(s, " ")
	reversed := make([]string, len(words))
	copy(reversed, words)

	for i, j := 0, len(reversed)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}

	return strings.Join(reversed, " ")
}
