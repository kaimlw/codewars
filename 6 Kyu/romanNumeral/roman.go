package main

import (
	"fmt"
	"strings"
)

var romanValue = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

var romanCombination = map[string] string{
	"I": "XV",
	"X": "LC",
	"C": "DM",
}

func Decode(roman string) int {
	romanSplitted := strings.Split(roman, "")
	result := 0
	combinationFlag := false
	lastLetter := ""
	
	for i,letter := range romanSplitted{
		if len(romanCombination[letter]) != 0 && !combinationFlag && i != len(romanSplitted)-1 {
			combinationFlag = true
			lastLetter = letter
			continue		
		}

		if combinationFlag {
			if strings.Contains(romanCombination[lastLetter], letter) {
				result += romanValue[letter] - romanValue[lastLetter]
			} else {
				result += romanValue[letter] + romanValue[lastLetter]
			}
		} else {
			result += romanValue[letter]
		}

		combinationFlag = false
		lastLetter = letter
	}

	return result
}

func Decode2(roman string) int {
	romanSplitted := strings.Split(roman, "")
	romanLeft := len(romanSplitted)
	result := 0
	letter := ""
	letterAfter := ""

	if len(romanSplitted) == 1 {
		result = romanValue[romanSplitted[0]]
	}
	
	for i := 0; i < len(romanSplitted)-1; i++ {
		letter = romanSplitted[i]
		letterAfter = romanSplitted[i+1]

		fmt.Println(letter)

		if len(romanCombination[letter]) != 0 && strings.Contains(romanCombination[letter], romanSplitted[i+1]) {
			result += romanValue[letterAfter] - romanValue[letter]
			i ++
			romanLeft -= 2
		} else {
			result +=  romanValue[letter]
			romanLeft -= 1
		}

	}
 
	if romanLeft > 0 {
		result += romanValue[romanSplitted[len(romanSplitted)-1]]
	}

	return result
}
var mapping map[rune]int = map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L':50, 'C':100, 'D':500, 'M':1000}
func Decode3(roman string) int {
  var intern []int
  var result int
  for _, r := range roman {
    intern = append(intern, mapping[r])
  }
  
  for i := 1; i < len(intern); i++ {
    if intern[i-1] >= intern[i] {
      result += intern[i-1]
    } else {
      result -= intern[i-1]
    }
  }
  
  result += intern[len(intern)-1]
  
  return result
}

func main()  {
	// roman := "MMVIII"
	// roman := "MMMDCCCXL"
	// roman := "MMMCMLX"
	// roman := "I"
	// roman := "MDCLXVI"
	roman := "XCV"
	fmt.Println(Decode(roman))
	fmt.Println(Decode2(roman))
}