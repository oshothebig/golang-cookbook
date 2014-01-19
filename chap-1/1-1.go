package main

import "fmt"

func main() {
	text := "ようこそ世界"

	for _, char := range text {
		fmt.Printf("%c\n", char)
	}
}
