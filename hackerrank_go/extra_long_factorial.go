package main

import (
	"fmt"
	"math/big"
)

func extraLongFactorials(n int32) {
	N := big.NewInt(0)
	N.MulRange(1, int64(n))
	fmt.Println(N)
}

func main() {
	extraLongFactorials(30)
}