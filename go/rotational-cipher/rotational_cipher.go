// Package rotationalcipher impliments a utility
// that provides a simple rotational cipher
// based on and input string and rotation index.
package rotationalcipher

import (
	"unicode"
)

// RotationalCipher is the function that the tests expect to
// pass values to.
func RotationalCipher(inString string, inShift int) string {
	var out string
	for _, val := range inString {
		switch {
		case unicode.IsUpper(val) && (val+rune(inShift)) > 'Z':
			out += string((val + rune(inShift)) - rune(26))
		case (val + rune(inShift)) > 'z':
			out += string((val + rune(inShift)) - rune(26))
		case inShift >= 26:
			out += string(val + rune(inShift) - rune(((inShift / 26) * 26)))
		case unicode.IsLetter(val):
			out += string(val + rune(inShift))
		default:
			out += string(val)
		}
	}
	return out
}
