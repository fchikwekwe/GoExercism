package isogram

import (
	"fmt"
	"strings"
)

// IsIsogram is a function that takes in a word as a string and outputs a bool
// letting you know whether the word is an isogram.
func IsIsogram(word string) bool {
	// Make all letters in word the same case
	uWord := strings.ToLower(word)
	fmt.Println("Word", uWord)

	// Initialize map variable
	m := make(map[byte]int)

	// Iterate over letters in the word
	for i := range word {
		if _, ok := m[uWord[i]]; ok {
			// Exclude hyphens and spaces from map
			if uWord[i] != 45 && uWord[i] != 32 {
				m[uWord[i]]++
			}
		} else {
			m[uWord[i]] = 1
		}
	}
	for _, v := range m {
		if v != 1 {
			return false
		}
	}
	return true
}
