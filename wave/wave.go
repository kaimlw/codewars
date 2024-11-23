package main

import (
	"fmt"
	"strings"
)

func wave(words string) []string {
	result := []string{}
	upperIndex := 0

	for _, letter1 := range words {
		if string(letter1) == " " {
			upperIndex++
			continue
		}

		waveWord := ""
		for idx, letter2 := range words{
			if idx == upperIndex {
				waveWord += strings.ToUpper(string(letter2))
			}else {
				waveWord += strings.ToLower(string(letter2))
			}
		}

		upperIndex++
		result = append(result, waveWord)
	}
	
	return result 
}

func wave2(words string) (wave []string) {
  wave = []string{} // leaving the array uninitialised would be better practice
  for i, c := range words {
    if c == ' ' {
      continue
    }
    upperC := string(c-'a'+'A')
    wave = append(wave, words[:i] + upperC + words[i+1:])
  }
  return
}

func main() {
	// str := "hello"
	// str := " x yz"
	fmt.Println(string('r'-'a'+'A'))
}