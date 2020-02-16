// Package lsproduct is a utility to calculate the largest
// product for contiguous substring of predefined length.
package lsproduct

import (
	"errors"
	"strconv"
	"unicode"
)

// LargestSeriesProduct is the function that the test suite
// expects to pass data to.
func LargestSeriesProduct(in string, n int) (int, error) {
	//var trimmed []string
	out := 0
	switch {
	case len(in) < n:
		return -1, errors.New("Span cannot be larger than string length")
	case len(in) == n:
		out = 1
		for i := 0; i < len(in); i++ {
			j, _ := strconv.Atoi(string(in[i]))
			out = out * j
		}
		return out, nil
	case in == "" || n == 0:
		return 1, nil
	case n < 0:
		return -1, errors.New("invalide span")
	}
	temp := 1
	out = 0
	for _, g := range in {
		if unicode.IsDigit(g) == false {
			return -1, errors.New("Non integer detected")
		}
	}
	for i := 0; i <= len(in)-n; i++ {
		temp = 1
		for j := 0; j < n; j++ {
			f, _ := strconv.Atoi(string(in[i+j]))
			temp = temp * f
		}
		if out < temp {
			out = temp
		}
	}
	return out, nil
}
