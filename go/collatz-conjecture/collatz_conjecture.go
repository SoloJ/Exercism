package collatzconjecture

import (
	"errors"
)

//CollatzConjecture is the function that the test will pass arguments to
func CollatzConjecture(x int) (int, error) {
	steps := 0
	if x == 0 || x < 0 {
		return 0, errors.New("Invalid input value")
	}
	for x != 1 {
		switch {
		case x%2 == 0:
			steps++
			x = x / 2
		default:
			x = x*3 + 1
			steps++
		}
	}
	return steps, nil
}
