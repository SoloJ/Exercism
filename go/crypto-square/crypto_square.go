// Package cryptosquare is a utility to encode
// strings of text.
package cryptosquare

import (
	"math"
	"unicode"
)

// Encode is the function that the test expects to
// pass strings to
func Encode(in string) string {
	return encrypt(rectangle(normalize(in)))
}
func normalize(in string) string {
	var out string
	for _, val := range in {
		if unicode.IsLetter(val) || unicode.IsDigit(val) {
			out = out + string(unicode.ToLower(val))
		}
	}
	return out
}
func rectangle(in string) ([]string, int, int) {
	stringLength := len(in)
	root := math.Sqrt(float64(len(in)))
	c := int(math.Ceil(root))
	var r int
	var blanks int
	switch {
	case root < float64(c):
		r = c - 1
	case root == float64(c):
		r = c
	}
	switch {
	case c*r-stringLength > 0:
		blanks = r*c - stringLength

	case c*r-stringLength < 0:
		blanks = 6 - (stringLength - r*c)
		r = c

	case c*r-stringLength == 0:
		blanks = 0
	}
	for i := 0; i < blanks; i++ {
		in = in + " "
	}
	counter := 0
	out := make([]string, 0)
	for i := 0; i < r; i++ {
		out = append(out, in[counter:counter+c])
		counter += c
	}
	return out, c, r
}
func encrypt(in []string, c int, r int) string {
	var out string
	for i := 0; i < c; i++ {
		if i != 0 {
			out = out + " "
		}
		for j := 0; j < r; j++ {
			out = out + string(in[j][i])
		}
	}
	return out
}
