package main

import (
	"fmt"
)

func cutTheSticks(arr []int32) []int32 {
    m := make([]int32, 1001)
    a := []int32{}
    n := int32(len(arr))
    for _, elem := range arr {
        m[elem]++
    }
    for _, elem := range m {
        if elem > 0 {
            a = append(a, n)
            n -= elem
        }
    }

    return a
}

func main() {
	arr := []int32{5, 4, 4, 2, 2, 8}
	fmt.Println(cutTheSticks(arr))
}