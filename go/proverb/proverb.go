package proverb

import (
	"fmt"
)

// Proverb is the function that the test passes arguments to
func Proverb(rhyme []string) []string {
	s := make([]string, 2)
	s[0] = "For want of a %s the %s was lost."
	s[1] = "And all for the want of a %s."

	x := len(rhyme)
	answer := make([]string, x)
	switch {
	case len(rhyme) == 1:
		answer[0] = fmt.Sprintf(s[1], rhyme[0])
	case len(rhyme) > 1:
		for j := 0; (j + 1) < x; j++ {
			answer[j] = fmt.Sprintf(s[0], rhyme[j], rhyme[j+1])
		}
		answer[x-1] = fmt.Sprintf(s[1], rhyme[0])
	}
	return answer
}
