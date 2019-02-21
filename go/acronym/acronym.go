package acronym

import (
	"fmt"
	"strings"
)

// Abbreviate takes a phrase and returns its acronym
func Abbreviate(s string) string {

	fmt.Println("s", s)
	// Replace hyphens with spaces and make it into a slice
	words := strings.Fields(strings.Replace(s, "-", " ", -1))
	fmt.Println("words", words)
	acronym := ""

	for _, word := range words {
		fmt.Println(acronym)
		// Iterate over the slice and get the first index of each word
		acronym += strings.ToUpper(string(word[0]))
	}
	return acronym
}
