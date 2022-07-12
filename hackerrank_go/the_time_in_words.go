package main

import (
	"fmt"
)

func timeInWords(h int32, m int32) string {
	var (
		hour, minute, result string
	)

	oneToNine := map[int32]string{
		0: "o' clock",
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
	}

	tenToNineteen := map[int32]string{
		10: "ten",
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		14: "fourteen",
		15: "quarter",
		16: "sixteen",
		17: "seventeen",
		18: "eighteen",
		19: "nineteen",
	}

	twentyToSixty := map[int32]string{
		20: "twenty",
		30: "half",
		40: "forty",
		50: "fifty",
		60: "sixty",
	}

	if m == 0 && h < 10 {
		result = oneToNine[h] + " " + oneToNine[0]
		return result
	} else if m == 0 && h > 10 {
		result = tenToNineteen[h] + " " + oneToNine[0]
		return result
	}

	if h < 10 && m > 30 {
		hour = oneToNine[h+1]
	} else if h >= 10 && m > 30 {
		hour = tenToNineteen[h+1]
	}

	if h < 10 && m <= 30 {
		hour = oneToNine[h]
	} else if h > 10 && m <= 30 {
		hour = tenToNineteen[h]
	}

	if m < 30 && m > 0 {
		if m == 1 {
			minute = oneToNine[1] + " minute " + "past "
		} else if m > 9 && m < 19 && m != 15 {
			minute = tenToNineteen[m] + " minutes " + "past "
		} else if m > 19 && m < 30 {
			if m%10 != 0 {
				minute = twentyToSixty[20] + " " + oneToNine[m%10] + " minutes " + "past "
			} else {
				minute = twentyToSixty[20] + " minutes past "
			}

		} else {
			minute = tenToNineteen[m] + " past "
		}
	} else if m == 30 {
		minute = twentyToSixty[m] + " past " 
	} else if m > 30 && m < 59 {

		if (60 - m) == 15 {
			minute = tenToNineteen[15] + " to "
		} else {
			reminder := 60 - m
			if reminder == 1 {
				minute = oneToNine[1] + " minute to "
			}

			if reminder > 1 && reminder < 9 {
				minute = oneToNine[reminder] + " minutes to "
			}

			if reminder > 9 && reminder < 19 {
				minute = tenToNineteen[reminder] + " minutes to "
			}

			if reminder > 19 && reminder < 30 {
				if reminder%10 != 0 {
					minute = twentyToSixty[20] + " " + oneToNine[reminder%10] + " minutes to "
				} else {
					minute = twentyToSixty[20] + " minutes to "
				}
			}
		}
	}

	return minute + hour
}

func main() {

	result := timeInWords(10, 57)
	fmt.Println(result)
}
