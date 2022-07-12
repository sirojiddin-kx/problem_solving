package main

import (
	"fmt"
)

func abs(num int32) int32 {

	if num > 0 {

		return num
	}

	return -(num)
}

func taumBday(b int32, w int32, bc int32, wc int32, z int32) int64 {
	// var (
	// 	overAllSum int64  = int64(b)*int64(bc) + int64(w)*int64(wc)
	// 	minimumSum int64
	// )
	// fmt.Println(b * bc)
	// fmt.Println(w * wc)
	// fmt.Println("OverAllSum : ", overAllSum)
	// if bc < z && wc < z {

	// 	return int64(overAllSum)
	// }

	// if bc == z && wc == z {

	// 	return overAllSum
	// }

	// if bc < wc {
	// 	minimumSum = int64(b*bc + w*bc + w*z)

	// 	if minimumSum < overAllSum {

	// 		return int64(minimumSum)
	// 	} else {

	// 		return int64(overAllSum)
	// 	}
	// } else {
	// 	minimumSum = int64(b*wc + w*wc + b*z)

	// 	if minimumSum < overAllSum {

	// 		return int64(minimumSum)
	// 	} else {

	// 		return int64(overAllSum)
	// 	}
	// }
	var result int64
	if bc > wc+z && z < bc+wc {
		result = int64(w)*int64(wc) + int64(b)*(int64(wc)+int64(z))
	} else if wc > bc+z && z < bc+wc {
		result = int64(b)*int64(bc) + int64(w)*(int64(bc)+int64(z))
	} else {
		result = int64(b)*int64(bc) + int64(w)*int64(wc)
	}
	return result
}

func main() {
	res := taumBday(int32(71805), int32(9169), int32(905480), int32(255669), int32(334440))
	fmt.Println(res)
}
