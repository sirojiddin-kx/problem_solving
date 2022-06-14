package main

import (
	"fmt"
)

func formingMagicSquare(s [][]int32) int32 {
	var (
		minimalCost int32
	)

	for i := 0; i < 3; i++ {
		var sum int32 = 0
		for j := 0; j < 3; j++ {
			sum += s[i][j]
		}

		if sum < 15 {
			minimalCost += 15 - sum
		}
	}

	return minimalCost
}

func main() {
	sample := [][]int32{{5, 3, 4}, {1, 5, 8}, {6, 4, 2}}
	fmt.Println(sample[0])
	fmt.Println(formingMagicSquare(sample))
}