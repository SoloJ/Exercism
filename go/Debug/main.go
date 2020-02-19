package main

import "fmt"

func main() {
	h := -121
	m := -5810
	var hour int
	var minute int
	min := m / 60
	if m < 0 {
		min = m/60 - 1
	}

	switch {
	case h >= 0:
		hour = ((h % 24) + min) % 24
	case (((h % 24) + 24) + min) < 0:
		hour = (((h%24)+24)+min)%24 + 24
	default:
		hour = (((h % 24) + 24) + min) % 24
	}
	switch {
	case m >= 0:
		minute = m % 60
	default:
		minute = (m % 60) + 60
	}
	fmt.Print(clockString(hour, minute))
}

// clockString takes in min and hour integers and
// converts them to a string expression of time.
func clockString(hour, minute int) string {
	var hourString string
	var minString string
	switch {
	case hour < 10:
		hourString = fmt.Sprintf("0%v:", hour)
	default:
		hourString = fmt.Sprintf("0%v:", hour)
	}
	switch {
	case minute < 10:
		minString = fmt.Sprintf("0%v", minute)
	default:
		minString = fmt.Sprintf("%v", minute)
	}
	return (hourString + minString)
}
