package main

import (
	"fmt"
)

func plusMinus(arr []int32) {
	var (
		zeroValue, positiveValue, negativeValue int32
	)

	for _, value := range arr {
		if value > 0 {
			positiveValue += 1
		} else if value < 0 {
			negativeValue += 1
		} else {
			zeroValue += 1
		}
	}
	fmt.Printf("%6f\n", float64(positiveValue)/float64(len(arr)))
	fmt.Printf("%6f\n", float64(negativeValue)/float64(len(arr)))
	fmt.Printf("%6f\n", float64(zeroValue)/float64(len(arr)))
}

func main() {
	arr := []int32{1, 1, 0, -1, -2}
	plusMinus(arr)
}
