package main

import (
	"fmt"
)

func leftRotation(a []int, size int, rotation int) []int {
	var (
		newarray []int
	)

	for i := 0; i < rotation; i++ {
		newarray = a[1:size]
		fmt.Println("new", newarray)
		newarray = append(newarray, a[0])
		a = newarray
		fmt.Println("out", a)
	}
	return a
}

// This is my solution. But some test cases failed for performance
func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {
	result := []int32{}
	for i := 0; i < int(k); i++ {
		value := a[0 : len(a)-1]
		newAraray := append(a[len(a)-1:], value...)
		a = newAraray
	}

	for j := 0; j < len(queries); j ++ {
		result = append(result, a[queries[j]])
	}
	return result
}

// I found this solution from discussion tab.
func circularArrayRotation(a []int32,k int32, m []int32) []int32 {
    var n int32 = int32(len(a))
    k=k%n
    var result []int32
    for _, v := range m {
        if v < k {
            result = append(result, a[n-k+v])
        } else {
            result = append(result, a[v-k])
        }
    }

    return result
}

