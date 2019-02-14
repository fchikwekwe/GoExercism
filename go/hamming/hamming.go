package hamming

import "fmt"

var a, b string = "GTCAGTCA", "GTAACTGT"

// Distance takes in two strings and returns their hamming distance as an int
func Distance(a, b string) (int, error) {

	hammingDistance := 0
	error := fmt.Errorf("cannot compare; strings are not the same length")
	if len(a) == len(b) {
		if len(a) == 0 {
			return hammingDistance, nil
		}
		for aIndex, aValue := range a {
			for bIndex, bValue := range b {
				if aIndex == bIndex && aValue != bValue {
					hammingDistance += 1
				}
			}
		}
	} else {
		return hammingDistance, error
	}
	// fmt.Printf("hamming: %v ", hammingDistance)
	return hammingDistance, nil
}
