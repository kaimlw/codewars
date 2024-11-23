package main

import "fmt"

func BreakChocolate(n, m int) int {
	if m < 1 || n < 1 {
		return 0
	}

	return m*n - 1
}

func main() {
	fmt.Println(BreakChocolate(5, 5))
}
