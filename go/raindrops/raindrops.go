package raindrops

import (
	"strconv"
)

//Convert is the function that the test file expects to pass arguments to
func Convert(input int) string {
	x := ""
	count := 0
	if input%3 == 0 {
		x += "Pling"
		count++
	}
	if input%5 == 0 {
		x += "Plang"
		count++
	}
	if input%7 == 0 {
		x += "Plong"
		count++
	}
	if count == 0 {
		return strconv.Itoa(input)
	}
	return x
}
