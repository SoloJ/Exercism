package isogram

import (
	"strings"
	"unicode"
)

//IsIsogram is the function that the test expect to pass arguments to
func IsIsogram(input string) bool {
	check := strings.ToUpper(input)
	for _, s := range check {
		if unicode.IsLetter(s) == false {
			continue
		}
		if strings.Count(check, string(s)) != 1 {
			return false
		}
	}
	return true
}
