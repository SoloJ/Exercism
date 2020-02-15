package listops

// IntList is an abstraction of a list of integers
type IntList []int

// PredFunc function to filter values
type PredFunc func(int) bool

// BinFunc function to combine values
type BinFunc func(int, int) int

// UnaryFunc function to transform values
type UnaryFunc func(int) int

//Append is the append function
func (ls IntList) Append(x IntList) IntList {

	answer := make([]int, 0) //initialize with size
	for i := 0; i < len(x); i++ {
		answer[i] = x[i]
	}
	for i := len(x); i < len(answer); i++ {
		answer[i] = x[i]
	}
	return answer
}

//Concat is the concat function
func Concat(x ...[]int) []int {
	answer := make([]int, 0)
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x[i]); j++ {
			answer[j] = x[i][j]
		}

	}
	return answer
}

/*
//Filter
func Filter(dna string) string {}

//Length
func Length(dna string) string {}

//Map
func Map(dna string) string {}

//Foldr
func Foldr(dna string) string {}

//Reverse
func Reverse(dna string) string {}
*/
