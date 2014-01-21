package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "abcdefgh"
	set1 := "xyz"
	set2 := "abc"

	fmt.Printf("Standard library: %v\n", strings.ContainsAny(text, set1))
	fmt.Printf("Standard library: %v\n", strings.ContainsAny(text, set2))

	fmt.Printf("containsAny: %v\n", containsAny(text, set1))
	fmt.Printf("containsAny: %v\n", containsAny(text, set2))

	fmt.Printf("containsAnyWithMap: %v\n", containsAnyWithMap(text, set1))
	fmt.Printf("containsAnyWithMap: %v\n", containsAnyWithMap(text, set2))
}

func containsAny(text, set string) bool {
	for _, c := range text {
		for _, s := range set {
			if c == s {
				return true
			}
		}
	}

	return false
}

func containsAnyWithMap(text, set string) bool {
	// make set
	m := make(map[rune]struct{})
	for _, c := range set {
		m[c] = struct{}{}
	}

	// then check
	for _, c := range text {
		if _, ok := m[c]; ok {
			return true
		}
	}

	return false
}
