// Package clock is a utility that takes in hour and
// minute data and outputs a time based on requirements.
package main

import "fmt"

// CreateClock takes in hour and min its and
// outputs the string clock reading they denote.
func main() {
	h := 201
	m := 3001
	hour := h % 24
	var minute int
	switch {
	case m >= 0 && m < 60:
		minute = m
	case m < 0 && m > -60:
		switch {
		case hour-1 < 0:
			hour = (hour - 1) % 24
			minute = m % 60
		default:
			hour = hour - 1
			minute = m % 60
		}
	case m < 0:
		switch {
		case hour+m/60 < 0:
			hour = (hour + m/60) % 24
			minute = m % 60
		default:
			hour = hour + m/60
			minute = m % 60
		}
	case m > 0:
		switch {
		case hour+m/60 > 24:
			hour = (hour + m/60) % 24
			minute = m % 60
		default:
			hour = hour + m/60
			minute = m % 60
		}
	}
	fmt.Print(hour, minute)
}
