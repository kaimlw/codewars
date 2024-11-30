package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func IpsBetween(start, end string) int {
	startSplitted := strings.Split(start, ".")
	endSplitted := strings.Split(end, ".")
	result := 0

	for i := 3; i > 0; i-- {
		startInt,_ := strconv.Atoi(startSplitted[i])
		endInt,_ := strconv.Atoi(endSplitted[i])

		if i < 3 {
			result += int(math.Pow(256, float64(3-i))) * (endInt-startInt)
		} else {
			result += endInt - startInt
		}
	}

	return result
}

func main() {
	// start := "10.0.0.0"
	// end := "10.0.0.50"
	start := "20.0.1.10"
	end := "20.1.0.20"
	fmt.Println(IpsBetween(start,end))
}