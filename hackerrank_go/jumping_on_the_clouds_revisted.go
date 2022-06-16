package main 

import (
	"fmt"
)

func jumpingOnClouds(c []int32, k int32) int32 {
	var energy int32 = 100
	for i := 0; i < len(c); i += int(k) {

		if c[i] == 1 {
			energy -= 3
		} else {
			energy -= 1
		}

		if i + int(k)> len(c) {
			break
		}
	}

	return energy
}

func main() {
	c := []int32{0, 0, 1, 0, 0, 1, 1, 0}

	fmt.Println(jumpingOnClouds(c, 2))
}