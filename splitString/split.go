package main

import (
	"fmt"
	"strings"
)

func ToCamelCase1(s string) string {
	var splittedString []string
	if strings.Contains(s, "_") {
		splittedString = strings.Split(s, "_")
	} else if strings.Contains(s, "-") {
		splittedString = strings.Split(s, "-")
	} else {
		return s
	}

	result := ""
	for i,word := range(splittedString) {
		if i == 0 {
			result += word
			continue
		}
		result += strings.ToUpper(word[:1])+strings.ToLower(word[1:])
	}

	return result
}

func ToCamelCase2(s string) {
	// var result []rune
	runes := []rune(s)

	fmt.Println(runes)
	// re
}

func main() {	
	splittedString := ToCamelCase1("The_stealth_warrior")
	
	fmt.Println(splittedString)
}