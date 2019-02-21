// Package bob checks what you say to bob and then gives a response.
package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Hey returns a reply from lackadaisical bob
// depending on what yoy say (or don't say) to him.
func Hey(remark string) string {
	remark = strings.Trim(remark, " ")
	remark = strings.Trim(remark, "\n")
	remark = strings.Trim(remark, "\t")

	if len(remark) < 1 {
		return "Fine. Be that way!"
	}

	if strings.Contains(remark, "?") {
		for _, rune := range remark {
			fmt.Println(remark, rune)
			if unicode.IsLower(rune) {
				// Question with some lowercase letters
				return "Sure."
			}
		}
		// All letters are capital and question
		return "Calm down, I know what I'm doing!"
	}

	for _, rune := range remark {
		if unicode.IsLower(rune) {
			return "Whatever."
		}
	}
	return "Whoa, chill out!"
}

func main() {
	fmt.Println(Hey("\n"))
}
