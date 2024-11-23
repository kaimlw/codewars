package main

import (
	"fmt"
	"math"
)

func Persistent(n int) int {
	digits := int(math.Log10(float64(n)))+1
	counter := 0
	number := n
	
	for digits > 1 {
		counter += 1
		result := 1

		for i := digits; i > 0; i-- {
			result *= number%10
			number = number/10
		}

		number = result
		digits = int(math.Log10(float64(number)))+1
	}

	return counter
}

func Persistent2(n int) int {
	counter := 0

	for n >= 10 {
		result_temp := 1
		
		for n > 0{
			result_temp *= n%10
			n /= 10 
		}
		
		n = result_temp
		counter ++
	}

	return counter
}

func main() {
	// Persistent(39)

	fmt.Println(Persistent2(999))
}