package reverse

// String is a function that returns the reverse of the input string
func String(word string) string {
	backwards := ""
	for _, rune := range word {
		backwards = string(rune) + backwards
	}
	return backwards
}
