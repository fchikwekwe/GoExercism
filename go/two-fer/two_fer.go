// Package twofer returns a string that either replaces X with a name or
// replaces X with you.
package twofer

import "fmt"
// ShareWith checks to see if we should return a string with name or with you
func ShareWith(name string) string {
	if name != "" {
		return fmt.Sprintf("One for %v, one for me.", name)
	}
	return fmt.Sprintf("One for you, one for me.")
}
