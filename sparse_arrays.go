package main

import (
	"fmt"
)

// My solution
func matchingStrings(strings []string, queries []string) []int32 {

	var (
		result []int32
	)

	for _, query := range queries {
		var (
			count int32
		)

		for _, str := range strings {
			if query == str {
				count += 1
			}
		}

		result = append(result, count)
	}

	return result
}

// Solution from discussion
func matchingStrings(inputStrings []string, queries []string) []int32 {
    // Write your code here
    res:= make([]int32, len(queries))
    stringCounts := make(map[string]int32)
    
    for _, str := range inputStrings {
        stringCounts[str]++
    }
	fmt.Println("Input string ", stringCounts)
    for i, query := range queries {
        res[i]=stringCounts[query]
    }
    
    return res
}


func main() {
	strings := []string{"ab", "ab", "abc"}
	queries := []string{"ab", "abc", "bc"}

	result := matchingStrings(strings, queries)
	fmt.Println("res ; ", result)
}