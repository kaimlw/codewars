package main

import (
	"fmt"
	"strings"
)

func FindShort(s string) int {
	words := strings.Split(s, " ");
	shortest := len(words[0])

	for _, word := range words {
		if len(word) < shortest {
			shortest = len(word)
		}
	}

	return shortest
}

func main() {
	// text := "bitcoin take over the world maybe who knows perhaps"
	// text := "turns out random test cases are easier than writing out basic ones"
	text := "Let's travel abroad shall we"
	fmt.Println(FindShort(text))
}