package twofer

import "fmt"

//ShareWith is the function that the test file expects to pass arguments to
func ShareWith(name string) string {
	if len(name) == 0 {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
