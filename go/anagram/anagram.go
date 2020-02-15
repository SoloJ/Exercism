// Package anagram is a simple untility that takes a word as input
// and a list of candidate words and returns all valid anagrams from
// from the candidate list.
package anagram

import (
	"strings"
	"unicode"
)

// Detect is the function that the test expects to pass variables to
func Detect(subject string, candidates []string) []string {
	index := make([]int, len(candidates))
	out := make([]string, 0)
	for _, val := range subject {
		val = unicode.ToUpper(val)
		for keyCandidates := range candidates {
			switch {
			case strings.ContainsRune(strings.ToUpper(candidates[keyCandidates]), val) == false:
				index[keyCandidates] = 0
			default:
				index[keyCandidates] = 1
			}
		}
	}
	for key, val := range index {
		if val = 1 {
			out = append(out, candidates[key])
		}
	}
	return out

}
