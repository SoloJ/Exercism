package hamming

import (
	"errors"
)

//Distance is the function that the test expects to pass argumenets to
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("DNA Length Missmatch")
	}
	a1 := []rune(a)
	b1 := []rune(b)
	count := 0
	for x, u := range a1 {
		if u != b1[x] {
			count++
		}
	}
	return count, nil
}
