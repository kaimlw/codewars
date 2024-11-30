package main

import (
	"fmt"

)

func IsValidWalk(walk []rune) bool {
	if len(walk) != 10 {
		return false
	}

	vertical := 0
	horizontal := 0

	for _,direction := range walk {
		switch string(direction) {
		case "n":
			vertical += 1
		case "e":
			horizontal += 1
		case "s":
			vertical -= 1
		case "w":
			horizontal -=1
		}
	}

	return vertical == 0 && horizontal == 0
}

func main() {
	walk := []rune{'s', 'n', 'n', 's', 's', 'n', 'n', 'w', 'n', 's'}
	// walk := []rune{'s', 'e', 'w', 'n', 'n', 's', 'e', 'w', 'n', 's'}
	fmt.Println(IsValidWalk(walk))
}