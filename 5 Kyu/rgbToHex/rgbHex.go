package main

import (
	"fmt"
	"strings"
)

func RGB(r, g, b int) string {
	result := ""
	rgb := []int{r,g,b}

	for _, item := range rgb{
		if item > 255 {
			result += "FF"
			continue
		}
		if item < 0 {
			result += "00"
			continue
		}

		hex := fmt.Sprintf("%x", item)
		if len(hex) == 1 {
			result += "0" + hex
		} else {
			result += hex
		}
	}

	return strings.ToUpper(result)
}

// TOP ANSWER
func getValid(x int) int {
  switch {
    case x < 0: return 0
    case x > 255: return 255
    default: return x
  }
}

func RGB2(r, g, b int) string {
  res := fmt.Sprintf("%02X%02X%02X", getValid(r), getValid(g), getValid(b))
  fmt.Println(res)
  return res
}
// ----------------

func main() {
	// r := 148
	// g := 0
	// b := 211

	fmt.Println(strings.ToUpper("01"))
}