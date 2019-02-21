// Package bob checks what you say to bob and then gives a response.
package bob

import (
	"strings"
)

// Hey returns a reply from lackadaisical bob
// depending on what yoy say (or don't say) to him.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	silence := remark == ""
	yelling := remark == strings.ToUpper(remark) && strings.ContainsAny(remark, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	question := false
	if len(remark) > 0 {
		question = remark[len(remark)-1:] == "?"
	}

	if question && yelling {
		return "Calm down, I know what I'm doing!"
	}
	if yelling {
		return "Whoa, chill out!"
	}
	if question {
		return "Sure."
	}
	if silence {
		return "Fine. Be that way!"
	}
	return "Whatever."
}
