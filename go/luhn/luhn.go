package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Valid is a function that takes in an int and returns a bool. This bool describes whether the number is valid per the Luhn formula.
func Valid(number string) bool {
	// Check to see if the string length is <= 1
	if len(number) <= 1 {
		return false
	}
	number = strings.Replace(number, " ", "", -1)
	numArr := strings.Split(number, "")

	var sum int

	for i := len(number); i > -1; i-- {
		num, err := strconv.Atoi(numArr[i])
		// Check to see if all values in string are valid digits
		if err != nil {
			return false
		}
		if i%2 == 0 {
			if num <= 4 {
				sum += num * 2
				fmt.Println("odd, small", numArr[i], sum)
			} else if num > 4 {
				sum += num*2 - 9
				fmt.Println("odd, large", numArr[i], sum)
			}
		} else {
			sum += num
			fmt.Println("even, _", numArr[i], sum)
		}
	}
	if sum%10 == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(Valid("0 91"))
}
