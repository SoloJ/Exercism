package strain

type (
	//Ints does stuff
	Ints []int
	//Lists does stuff
	Lists [][]int
	//Strings does stuff
	Strings []string
)

//Keep is the Int function that tests passes arguments to
func (z Ints) Keep(input func(int) bool) Ints {
	answer := make(Ints, 0)
	if z == nil {
		return nil
	}
	for _, x := range z {
		if input(x) {
			answer = append(answer, x)
		}
	}
	return answer
}

//Discard is the Int function that tests passes arguments to
func (z Ints) Discard(input func(int) bool) Ints {
	answer := make(Ints, 0)
	if z == nil {
		return nil
	}
	for _, x := range z {
		if input(x) == false {
			answer = append(answer, x)
		}
	}
	return answer
}

//Keep is the Lists function that tests passes arguments to
func (z Lists) Keep(input func([]int) bool) Lists {
	answer := make(Lists, 0)
	if z == nil {
		return nil
	}
	for _, x := range z {
		if input(x) {
			answer = append(answer, x)
		}
	}
	return answer
}

//Keep is the Strings function that tests passes arguments to
func (z Strings) Keep(input func(string) bool) Strings {
	answer := make(Strings, 0)
	if z == nil {
		return nil
	}
	for _, x := range z {
		if input(x) {
			answer = append(answer, x)
		}
	}
	return answer
}
