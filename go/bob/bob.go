package bob

import (
	"strings"
)

// Hey is a function that does things...like makes snarky answers
func Hey(input string) string {
	remark := strings.TrimSpace(input)
	a := len(remark)
	switch {
	case remark == "":
		return "Fine. Be that way!"
	case strings.ToUpper(remark) == remark && remark[a-1] == byte('?') && strings.ToLower(remark) != remark:
		return "Calm down, I know what I'm doing!"
	case strings.ToUpper(remark) == remark && strings.ToLower(remark) != remark:
		return "Whoa, chill out!"
	case remark[a-1] == byte('?'):
		return "Sure."
	default:
		return "Whatever."
	}
}
