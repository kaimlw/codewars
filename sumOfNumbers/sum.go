package main

import "fmt"

func GetSum(a, b int) int {
	if a == b {
		return a
	}

	// if a > b {
	// 	a,b = b,a
	// }

	result := 0
	if a < b {
		for i := a; i <= b; i++ {
			result += i
		}
	} else {
		for i := b; i <= a; i++ {
			result += i
		}
	}

	return result
}

func main() {
	fmt.Println(GetSum(-1,2))
}