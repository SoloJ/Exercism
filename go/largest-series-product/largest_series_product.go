// Package lsproduct is a utility to calculate the largest
// product for contiguous substring of predefined length.
package lsproduct

import (
	"errors"
	"fmt"
	"unicode"
)

// LargestSeriesProduct is the function that the test suite
// expects to pass data to.
func LargestSeriesProduct(in string, n int) (int, error) {
	//var temp []string
	//var trimmed []string
	var out int
	switch {
	case len(in) < n:
		return -1, errors.New("Span cannot be larger than string length")
	case len(in) == n:
		for i := 0; i < len(in); i++ {
			out = out * int(in[i])
			fmt.Print(int(in[i]), "\n")
		}
		return out, nil
	case in == "" || n == 0:
		return 0, nil
	}
	//trimmed = strings.Split(in, "0")
	for i := 0; i < len(in)-2; i++ {
		if unicode.IsDigit(rune(in[i])) == false {
			return -1, errors.New("nope")
		}
		if out < int(in[i])*int(in[i+1])*int(in[i+2]) {
		}
	}
	return out, nil
}
