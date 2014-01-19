package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	text := "文字列"

	fmt.Println(`"` + left(text, 30) + `"`)
	fmt.Println(`"` + right(text, 30) + `"`)
	fmt.Println(`"` + center(text, 30) + `"`)
}

func left(s string, total int) string {
	return align(s, total, func(length, total int) (left, right int) {
		return 0, total - length
	})
}

func right(s string, total int) string {
	return align(s, total, func(length, total int) (left, right int) {
		return total - length, 0
	})
}

func center(s string, total int) string {
	return align(s, total, func(length, total int) (left, right int) {
		diff := total - length
		return diff / 2, diff - diff/2
	})
}

func align(s string, total int, split func(int, int) (int, int)) string {
	length := utf8.RuneCountInString(s)
	if length >= total {
		return s
	}

	left, right := split(length, total)

	ls := strings.Repeat(" ", left)
	rs := strings.Repeat(" ", right)

	return ls + s + rs
}
