package main

import (
	"fmt"
	"sort"
)

func covertInt32ToInt(a []int32) (result []int) {

	for _, val := range a {
		result = append(result, int(val))
	}
	return result
}

func ReturnIndex(arr []int, num int32) int32 {

	for i := 0; i < len(arr); i++ {

		if arr[0] < int(num) {

			return 1
		} else if arr[i] < int(num) {

			return int32(i + 1)
		} else if arr[i] == int(num) {

			return int32(i+1)
		}
	}

	return int32(len(arr)) + 1
}

func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	var (
		rank   = make(map[int32]int32)
		arr    = []int{}
		result = []int32{}
	)

	for _, val := range ranked {
		rank[val] += 1
	}

	for key := range rank {
		arr = append(arr, int(key))
	}

	sort.Sort(sort.Reverse(sort.IntSlice((arr))))
	for _, val := range player {
		index := ReturnIndex(arr, val)
		result = append(result, index)
	}

	return result
}

func main() {
	ranked := []int32{100, 90, 90, 80}
	player := []int32{70, 80, 105}

	fmt.Println(climbingLeaderboard(ranked, player))
}
