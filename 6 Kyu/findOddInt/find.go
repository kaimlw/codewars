package main

import "fmt"

func FindOdd(seq []int) int {
	appearsCounter := map[int]int{}

	for _,num := range seq {
		appearsCounter[num] += 1
	}

	for i,item := range appearsCounter {
		if item % 2 != 0 {
			return i
		}
	}

	return 0
}

func FindOdd2(seq []int) int {
	res := 0
	for _, x := range seq {
			//"^" is XOR, Only true if two item diff, e.g: 5 ^ 5 = 0; 2 ^ 3 ^ 2 = 3;
			res ^= x 
	}
	return res
}

func main() {
	fmt.Println(FindOdd([]int{1,2,2,3,3,3,4,3,3,3,2,2,1}))
}