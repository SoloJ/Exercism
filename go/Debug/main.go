package main

import "fmt"

func main() {
	in := 10
	temp := make([]bool, in+1)
	out := make([]int, 0)
	if in == 2 {
		out = append(out, 2)
	}
	for i := 2; i < in+1; i++ {
		if !temp[i] {
			out = append(out, i)
			fmt.Print(i)
		}
		for p := i; p < in+1; p += i {
			temp[p] = true
		}
	}
	fmt.Print(out)
}
