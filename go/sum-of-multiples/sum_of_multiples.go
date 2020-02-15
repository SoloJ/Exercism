package summultiples

//SumMultiples is the function that the tests will pass arguments to
func SumMultiples(in int, divisors ...int) int {
	var out int
	temp := make([]int, in)
	if in == 0 {
		return 0
	}
	for _, val := range divisors {
		for i := val; i < in; i += val {
			temp[i] = i
		}

	}
	for _, x := range temp {
		out = out + x
	}
	return out

}

/*
func sumFinder(in int, div1 int) int {
	if in == 0 || div1 == 0 {
		return 0
	}
	var out int
	for i := 1; i <= ((in - 1) / div1); i++ {
		out = out + (div1 * i)
	}
	fmt.Print(out, " ")
	return out
}

func lcmFinder(in []int) int {
	if len(in) == 1 {
		return 0
	}
	counter := 0
	var b int
	for i := 1; counter != (len(in)); i++ {
		b = i * in[0]
		for _, val := range in {
			if b%val == 0 {
				counter++
			} else {
				counter = 0
				break
			}
			if counter == len(in) {
				return b
			}
		}
	}
	return 0
}
*/
