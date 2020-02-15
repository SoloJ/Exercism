// Package luhn impliments a function to validate
// if a string passes the Luhn algorithm.
package luhn

import (
	"unicode"
)

// Valid determins if a string passes the Luhn algorithm.
func Valid(in string) bool {
	var temp string
	stringSumOdd := 0
	stringSumEven := 0
	j := 0
	if len(in) < 2 {
		return false
	}
	for _, x := range in {
		switch {
		case unicode.IsSpace(x):
			continue
		case unicode.IsDigit(x):
			temp = string(x) + temp
			j++
		default:
			return false
		}
		switch {
		case j == 0:
			continue
		case j%2 == 0:
			stringSumOdd += doubleCheck(int(x - '0'))
			stringSumEven += int(x - '0')
		default:
			stringSumEven += doubleCheck(int(x - '0'))
			stringSumOdd += int(x - '0')
		}
	}
	switch {
	case len(temp) < 2:
		return false
	case len(temp)%2 == 0:
		if stringSumEven%2 == 0 {
			return true
		}
		return false
	default:
		if stringSumOdd%2 == 0 {
			return true
		}
		return false
	}
}

// doubleCheck checks if the number should have 9 subtracted from it.
func doubleCheck(in int) int {
	if 2*in > 9 {
		return (in*2 - 9)
	}
	return (2 * in)
}
