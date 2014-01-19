package main

import "fmt"

func main() {
	//	Convert string to Unicode codepoints
	codePoints := []rune("こんにちは世界")

	//	Convert Unicode codepoint to string
	text := string(codePoints)

	fmt.Println(codePoints)
	fmt.Println(text)
}
