// Package lsproduct is a utility to calculate the largest
// product for contiguous substring of predefined length.
package lsproduct

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

// LargestSeriesProduct is the function that the test suite
// expects to pass data to.
func LargestSeriesProduct(in string, n int) (int, error) {
	//var temp []string
	//var trimmed []string
	out := 1
	switch {
	case len(in) < n:
		return -1, errors.New("Span cannot be larger than string length")
	case len(in) == n:
		for i := 0; i < len(in); i++ {
			j, _ := strconv.Atoi(string(in[i]))
			out = out * j
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
		a, _ := strconv.Atoi(string(in[i]))
		b, _ := strconv.Atoi(string(in[i+1]))
		c, _ := strconv.Atoi(string(in[i+2]))
		if out < a*b*c {
			out = a * b * c
			fmt.Print(out, "\n")

		}
	}
	return out, nil
}
