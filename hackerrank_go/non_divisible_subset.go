package main

import (
	"fmt"
	"sort"
)

func convertMapToSlice(mp map[int32]int32) (resp []int) {

	for k := range mp {
		resp = append(resp, int(k))
	}
	sort.Sort((sort.IntSlice((resp))))
	return resp
}

func nonDivisibleSubset(k int32, s []int32) int32 {
	var (
		frequencyOfNumbers = make(map[int32]int32)
		result             int32
	)

	for _, val := range s {
		frequencyOfNumbers[(val % k)] += 1
	}

	mpKeys := convertMapToSlice(frequencyOfNumbers)
	left := 0
	right := len(mpKeys) - 1
	mid := len(mpKeys) / 2
	fmt.Println("mp keys := ", mpKeys)
	if len(mpKeys)%2 != 0 {
		result += frequencyOfNumbers[int32(mpKeys[mid])]
	}

	for i := 0; i < mid; i++ {
		fmt.Println(left, right)
		if (mpKeys[left]+mpKeys[right])%int(k) == 0 {

			if frequencyOfNumbers[int32(mpKeys[left])] > frequencyOfNumbers[int32(mpKeys[right])] {
				fmt.Println("left")
				result += frequencyOfNumbers[int32(mpKeys[left])]
			} else if frequencyOfNumbers[int32(mpKeys[left])] < frequencyOfNumbers[int32(mpKeys[right])] {
				fmt.Println("right")
				result += frequencyOfNumbers[int32(mpKeys[right])]
			} else {
				fmt.Println("eq")
				result += frequencyOfNumbers[int32(mpKeys[right])]
			}
		} else {
			fmt.Println("not eq")
			fmt.Println("res", result)
			result += (frequencyOfNumbers[int32(mpKeys[left])] + frequencyOfNumbers[int32(mpKeys[right])])
		}

		left += 1
		right -= 1
	}

	fmt.Println(frequencyOfNumbers)
	fmt.Println(frequencyOfNumbers)
	return result
}

func main() {
	//arr1 := []int32{1, 7, 2, 4, 5}
	//arr1 := []int32{1, 5, 7, 9, 11}
	//	arr3 := []int32{422346306, 940894801, 696810740, 862741861, 85835055, 313720373}
	arr4 := []int32{582740017, 954896345,  590538156,  298333230,  859747706,  155195851,  331503493,  799588305,  164222042,  563356693, 80522822 ,432354938, 652248359}
	//arr := []int32{278, 576, 496, 727, 410, 124, 338, 149, 209, 702, 282, 718, 771, 575, 436}
	fmt.Println(nonDivisibleSubset(int32(11), arr4))

}
