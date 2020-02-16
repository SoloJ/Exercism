/// Package lsproduct is a utility to calculate the largest
// product for contiguous substring of predefined length.
package main

import (
	"fmt"
	"strconv"
)

// LargestSeriesProduct is the function that the test suite
// expects to pass data to.
func main() {
	//var trimmed []string
	in := "0123456789"
	n := 2
	out := 1
	switch {
	case len(in) < n:
		fmt.Print("return -1")
	case len(in) == n:
		for i := 0; i < len(in); i++ {
			j, _ := strconv.Atoi(string(in[i]))
			out = out * j
		}
		fmt.Print("return -1")
	case in == "" || n == 0:
		fmt.Print(out)
	}
	temp := 1
	out = 1
	for i := 0; i <= len(in)-n; i++ {
		temp = 1
		for j := 0; j < n; j++ {
			f, _ := strconv.Atoi(string(in[i+j]))
			temp = temp * f
			fmt.Print("1 ", temp, " ", f, "\n")
		}
		if out < temp {
			out = temp
			fmt.Print("2 ", out, "\n")
		}
		fmt.Print("3 clear", out, "\n")
	}
	fmt.Print(out)
}
