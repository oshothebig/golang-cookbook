package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	text := "\n空白削除  \n"

	fmt.Printf("%#v\n", text)
	fmt.Printf("%#v\n", strings.TrimFunc(text, unicode.IsSpace))
	fmt.Printf("%#v\n", strings.TrimLeftFunc(text, unicode.IsSpace))
	fmt.Printf("%#v\n", strings.TrimRightFunc(text, unicode.IsSpace))
}
