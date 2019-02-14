package main

import (
	"fmt"
	"strings"
)

// Valid is a function that takes in an int and returns a bool. This bool describes whether the number is valid per the Luhn formula.
func Valid(number string) bool {
	// Check to see if the string length is <= 1
	if len(number) <= 1 {
		return false
	}
	number = strings.Replace(number, " ", "", -1)
	fmt.Println(number)
	return true
}

func main() {
	Valid("0 91")
}
