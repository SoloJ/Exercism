package etl

import (
	"strings"
)

//Transform is the function that the tests expect to pass arguments to
func Transform(input map[int][]string) map[string]int {

	answer := make(map[string]int)

	for k, e := range input {
		for _, e2 := range e {
			answer[strings.ToLower(e2)] = k
		}
	}
	return answer
}
