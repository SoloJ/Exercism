package summultiples

// Source: exercism/problem-specifications
// Commit: bd2d4d9 sum-of-multiples: the factor 0 does not affect the sum of multiples of other factors
// Problem Specifications Version: 1.5.0

var varTests = []struct {
	divisors []int
	limit    int
	sum      int
}{
	{[]int{3, 5}, 1, 0},                            // no multiples within limit
	{[]int{3, 5}, 4, 3},                            // one factor has multiples within limit
	{[]int{3}, 7, 9},                               // more than one multiple within limit
	{[]int{3, 5}, 10, 23},                          // more than one factor with multiples within limit
	{[]int{3, 5}, 100, 2318},                       // each multiple is only counted once
	{[]int{3, 5}, 1000, 233168},                    // a much larger limit
	{[]int{7, 13, 17}, 20, 51},                     // three factors
	{[]int{4, 6}, 15, 30},                          // factors not relatively prime
	{[]int{5, 6, 8}, 150, 4419},                    // some pairs of factors relatively prime and some not
	{[]int{5, 25}, 51, 275},                        // one factor is a multiple of another
	{[]int{43, 47}, 1000000000, 22018801927016573}, // much larger factors
	{[]int{1}, 100, 4950},                          // all numbers are multiples of 1
	// the factor 0 does not affect the sum of multiples of other factors
}
