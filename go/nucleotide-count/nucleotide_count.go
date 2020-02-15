package dna

import (
	"errors"
)

//Histogram is the map that the test expects as an answer
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

//Counts is the function that the tests expect to pass arguments to
func (d DNA) Counts() (Histogram, error) {
	var h Histogram
	h = map[rune]int{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}
	answer := []rune(d)
	switch {
	case len(answer) == 0:
		break
	default:
		for i := 0; i < len(answer); i++ {
			if val, ok := h[answer[i]]; ok {
				h[answer[i]] = val + 1
				continue
			}
			return h, errors.New("Invalid characters detected")
		}
	}
	return h, nil
}
