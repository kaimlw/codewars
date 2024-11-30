package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Order(sentence string) string {
	if sentence == "" {
		return sentence
	}

	splittedSentence1 := strings.Split(sentence, " ")
	splittedSentence2 := strings.Split(sentence, " ")
	result := []string{}
	for i := 1; i <= len(splittedSentence1); i++ {
		notMatch := []string{}
		for _,word := range splittedSentence2{
			if strings.Contains(word, strconv.Itoa(i)) {
				result = append(result, word)
			}else{
				notMatch = append(notMatch, word)
			}
		}
		splittedSentence2 = notMatch
	}

	return strings.Join(result, " ")
}
func Order2(sentence string) string {
	if sentence == "" {
		return sentence
	}

	splittedSentence := strings.Split(sentence, " ")
	result := []string{}
	for i := 1; i <= len(splittedSentence); i++ {
		for _,word := range splittedSentence{
			if strings.Contains(word, strconv.Itoa(i)) {
				result = append(result, word)
			}
		}
	}

	return strings.Join(result, " ")
}

func main() {
	text := "4of Fo1r pe6ople g3ood th5e the2"
	// text := "is2 Thi1s T4est 3a"
	fmt.Println(Order(text))
	fmt.Println(Order2(text))
}