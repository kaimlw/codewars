package main

import "fmt"

func CountBy(x, n int) []int {
	result := []int{}
	for i := 1; i <= n; i++ {
		result = append(result, i*x)
	}
	return result
}

func main() {
	fmt.Println(CountBy(2,5))
}