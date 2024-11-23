package main

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	// "strings"
)

func ToJadenCase(s string) /*string*/ {
	var result []rune
	runes := []rune(s)
	// fmt.Println(runes)

	for _,letter := range runes{
		if letter == 32 {
			continue
		}

		letter = cases.Upper(letter)
	} 

	// return strings.ToTitle(s)
}

func main() {
	str := "How can mirrors be real if our eyes aren't real"
	// fmt.Println(ToJadenCase(str))
	ToJadenCase(str)
}