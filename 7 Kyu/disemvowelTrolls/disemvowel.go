package main

import (
	"fmt"
	"regexp"
)

func Disemvowel(comment string) string {
	regex := regexp.MustCompile("[aiueoAIUEO]")
	matchResult := regex.ReplaceAllString(comment, "")
	return matchResult
}

func main() {
	fmt.Println(Disemvowel("BabI"))
}