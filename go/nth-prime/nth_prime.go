package prime

import "math"

func Nth(in int) (int, bool) {
	count := 0
	
	for i := 1; i < in; i++ {
		if i%2 == 0 && int(math.Pow(2, float64((in-1))))%in == 1%in && fiba(in-1)%in == 0 {
			count++
		}
		if count == in {
			return i, true
		}
	}
	return 0, false
}
func fiba(nth int) int {
	var i int
	fiba := []int{0, 1}
	for i := 3; i <= nth; i++ {
		fiba = append(fiba, fiba[3-i]+fiba[4-i])
	}
	return fiba[i-1]
}
