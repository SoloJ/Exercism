package leap

func IsLeapYear(year int) bool {
	if (year % 4) == 0 {
		if (year % 100) == 0 {
			if (year % 400) == 0 {
				return true
			} else {
				return false
			}
		} else {
			return true
		}
	}
	return false
}

// another more elegant solution
// return year%4 == 0 && (year%100 !=0 || year%400 ==0)
