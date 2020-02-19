// Package clock is a utility that takes in hour and
// minute data and outputs a time based on requirements.
package clock

import (
	"fmt"
)

// Clock defines the struct that is the integer
// values of a clock.
type Clock struct {
	hours, minutes int
}

// New takes in hour and min its and
// outputs the string clock reading they denote.
func New(h, m int) Clock {
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
	return Clock{hour, minute}
}

// Add impliments a function that adds minutes
// to the clock time and returns a string.
func (c Clock) Add(minutes int) Clock {
	min := (minutes + c.minutes) / 60
	c.hours = (((c.hours) % 24) + min) % 24
	c.minutes = (minutes + c.minutes) % 60
	return c
}

// Subtract impliments a function that subtracts minutes
// to the clock time and returns a string.
func (c Clock) Subtract(minutes int) Clock {
	min := (minutes + c.minutes) / 60
	if (minutes + c.minutes) < 0 {
		min = (minutes+c.minutes)/60 - 1
	}
	switch {
	case ((((c.hours) % 24) + 24) + min) < 0:
		c.hours = ((((c.hours)%24)+24)+min)%24 + 24
	default:
		c.hours = ((((c.hours) % 24) + 24) + min) % 24
	}
	c.minutes = ((minutes + c.minutes) % 60) + 60

	return c
}

// clockString takes in min and hour integers and
// converts them to a string expression of time.
func (c Clock) String() string {
	var hourString string
	var minString string
	switch {
	case c.hours < 10:
		hourString = fmt.Sprintf("0%v:", c.hours)
	default:
		hourString = fmt.Sprintf("%v:", c.hours)
	}
	switch {
	case c.minutes < 10:
		minString = fmt.Sprintf("0%v", c.minutes)
	default:
		minString = fmt.Sprintf("%v", c.minutes)
	}
	return (hourString + minString)
}
