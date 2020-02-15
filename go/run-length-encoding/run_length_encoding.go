// Package encode impliments a simple utilty
// to run length encode strings of text.
package encode

import (
	"strconv"
	"strings"
	"unicode"
)

// RunLengthEncode accepts a string and return a
// run length encoded verion.
func RunLengthEncode(in string) string {
	count := 1
	var previousLetter rune
	var out string
	for key, val := range in {
		switch {
		case key == 0:
			previousLetter = val
			continue
		case val == previousLetter && len(in) != key+1:
			count++
		case val == previousLetter && len(in) == key+1:
			count++
			out += strconv.Itoa(count) + string(previousLetter)
			break
		case count > 1 && len(in) == key+1:
			out += strconv.Itoa(count) + string(previousLetter) + string(val)
			count = 1
		case count > 1:
			out += strconv.Itoa(count) + string(previousLetter)
			count = 1
		case count == 1 && len(in) == key+1:
			out += string(previousLetter) + string(val)
			count = 1
		default:
			out += string(previousLetter)
		}
		previousLetter = val
	}
	return out
}

// RunLengthDecode accepts a run length encoded
// string and returns a unencoded version.
func RunLengthDecode(in string) string {
	var out string
	var temp int
	for _, val := range in {
		switch {
		case unicode.IsDigit(val) && temp == 0:
			temp = int(val - '0')
			continue
		case unicode.IsDigit(val) && temp != 0:
			temp += temp*10 + int(val-'0') - temp
		case temp == 0:
			out += string(val)
			continue
		}
		switch {
		case temp != 0 && unicode.IsLetter(val):
			out += strings.Repeat(string(val), temp)
			temp = 0
		case temp != 0 && unicode.IsSpace(val):
			out += strings.Repeat(string(val), temp)
			temp = 0
		}
	}
	return out
}
