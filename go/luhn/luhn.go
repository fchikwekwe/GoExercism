package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Valid is a function that takes in an int and returns a bool. This bool describes whether the number is valid per the Luhn formula.
func Valid(number string) bool {
	number = strings.Replace(number, " ", "", -1)

	// Check to see if the string length is <= 1
	if len(number) < 2 {
		return false
	}

	var sum int64
	isOdd := len(number)%2 != 0

	for i := len(number) - 1; i >= 0; i-- {
		num, err := strconv.ParseInt(number[i:i+1], 0, 64)
		// Check to see if all values in string are valid digits
		if err != nil {
			return false
		}

		if isOdd {
			if i%2 != 0 {
				num *= 2
				if num > 9 {
					num -= 9
				}
			}

		} else {
			if i%2 == 0 {
				num *= 2
				if num > 9 {
					num -= 9
				}
			}
		}
		sum += num
	}
	return (sum%10 == 0)
}

func main() {
	fmt.Println(Valid("059"))
}
