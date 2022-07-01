package main

import (
	"fmt"
	"math"
)

func lenLoop(i int32) int32 {
    if i == 0 {
        return 1
    }
    count := 0
    for i != 0 {
        i /= 10
        count++
    }
    return int32(count)
}

func kaprekarNumbers(p int32, q int32) {
    var (
		div int32 = 10
		result []int32
	)

	for i := p; i <= q; i++ {

		divison := math.Pow(float64(div), float64(lenLoop(i)))
		square := int64(i * i)

		left  := square / int64(int32(divison))
		right := square % int64(int32(divison))

		sumOfTwo := left + right
		if int32(sumOfTwo) == i {
			result = append(result, i)
		}
	}

	if len(result) > 0 {
		fmt.Println(result)
	} else {
		fmt.Println("INVALID RANGE")
	}

}

func main() {
	kaprekarNumbers(1, 99999)
}