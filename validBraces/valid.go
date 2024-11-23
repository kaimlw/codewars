package main

import (
	"fmt"
	"strings"
)

func ValidBraces(str string) bool {
	validBrace := map[string]string{
		")" : "(",
		"}" : "{",
		"]" : "[",
	}

	stack := []string{}

	for _,brace := range strings.Split(str, ""){
		if brace == "{" || brace == "[" || brace == "(" {
			stack = append(stack, brace)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		if validBrace[brace] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}else{
			return false
		}
	}

	return len(stack) == 0
}


func main() {
	// input := "(){}[]"
	// input := "(}"
	// input := "[(])"
	// input := "([{}])"
	input := "[({})]()"

	fmt.Println(ValidBraces(input))
}