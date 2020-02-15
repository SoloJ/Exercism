package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	re := regexp.MustCompile("_|-")
	w := re.ReplaceAllLiteralString(s, " ")
	words := strings.Fields(w)
	x := len(words)
	answer := make([]string, x)
	for i := 0; i < x; i++ {
		answer[i] = strings.ToUpper(string(words[i][0]))
	}
	return strings.Join(answer, "")
}

//unc ReplaceAll(s, old, new string) string
