package main 

import (
	"fmt"
)

func nums(num int32) []int32 {
	var (
		result []int32
	)
	for num != 0 {
		rem := num % 10
		num /= 10
		result = append(result, rem)
	}

	return result
}

func findDigits(n int32) int32 {
	var (
		count int32
	)

	nums := nums(n)
	for _, num := range nums {
		if num != 0 && n % num == 0 {
			count += 1
		}
	}

	return count
}

func main() {
	res := findDigits(1012)
	fmt.Println(res)
}