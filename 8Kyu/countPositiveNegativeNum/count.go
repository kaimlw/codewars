package main

func CountPositivesSumNegatives(numbers []int) []int {
	res := make([]int, 2)
	if numbers == nil {
		return res
	}

	for _,num := range numbers {
		if num > 0 {
			res[0]++
		}
		if num < 0  {
			res[1] += num
		}
	}

	return res
}

func main()  {
	
}