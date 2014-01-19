package main

import "fmt"

func main() {
	text := "ようこそ世界"
	chars := []rune(text)

	for i := 0; i < len(chars); i++ {
		fmt.Printf("%c\n", chars[i])
	}
}
