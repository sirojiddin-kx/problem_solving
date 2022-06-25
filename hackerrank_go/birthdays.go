package main

import (
	"fmt"
)

func sum(squares []int32) (sumOfSquares int32) {

	for _, val := range squares {
		sumOfSquares += val
	}

	return sumOfSquares
}

func birthday(s []int32, d int32, m int32) int32 {
	var (
		possibleWays int32
	)
	for i := 0; i < len(s); i++ {

		if sum(s[i:(m+int32(i))]) == d {
			possibleWays += 1
		}

		if (int32(i) + m) >= int32(len(s)) {
			break
		}
	}

	return possibleWays
}

func main() {

	squares := []int32{1, 2, 1, 3, 2}
//	nextSquare := []int32{4}
	result := birthday(squares, 3, 2)
	fmt.Println(result)
}
