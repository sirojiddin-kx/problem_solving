package main

import (
	"fmt"
	"math"
	"strings"
)

func encryption(s string) string {
	var encryptedStr string
	removeSpace := strings.Replace(s, " ", "", -1)
	rows, columns := math.Floor(math.Sqrt(float64(len(removeSpace)))), math.Ceil(math.Sqrt(float64(len(removeSpace))))

	if int(rows) * int(columns) < len(removeSpace) {
		rows = columns
	}

	for i := 0; i < int(columns); i+=1 {
		var subStr string
		var (
			step int = i
		)
		for j := 0; j < int(rows); j+=1 {

			if (i + 1) * (j + 1) <= len(removeSpace) && step <= len(removeSpace) {
				subStr += string(removeSpace[step])
			}
			step += int(columns)
		}
		encryptedStr += subStr + " "
	}

	return encryptedStr
}

func main() {
	c := "haveaniceday"
	//s := "if man was meant to stay on the ground god would have given us roots"
	//fmt.Println(len(s))
	fmt.Println(encryption(c))
}