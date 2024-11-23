package main

import (
	"fmt"
	"sort"
	"strings"
)

func InArray(array1 []string, array2 []string) []string {
	inArray := []string{}
	seenSubstr := map[string]int{}

	for _, substr := range array1 {
		if seenSubstr[substr] == 1 {
			continue
		}
		seenSubstr[substr] = 1

		for _, word := range array2 {
			if strings.Contains(word, substr) {
				inArray = append(inArray, substr)
				break
			}
		}
	}
	
	sort.Strings(inArray)
	return inArray
}

func main() {
	array1 := []string{"live", "arp", "strong"}
	// array1 := []string{"tarp", "mice", "bull"}

	array2 := []string{"lively", "alive", "harp", "sharp", "armstrong"}

	fmt.Println(InArray(array1, array2))
}