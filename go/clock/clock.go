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
	min := m / 60
	if m < 0 && m != -60 {
		min = m/60 - 1
	}
	hour := modulo(modulo(h, 24)+min, 24)
	minute := modulo(m, 60)
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
	min := (c.minutes - minutes) / 60
	if (c.minutes-minutes) < 60 && (c.minutes-minutes) != 0 && (c.minutes-minutes) != -60 {
		min = min - 1
	}
	c.hours = modulo(c.hours+min, 24)
	c.minutes = modulo(c.minutes-minutes, 60)
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

// modulo impliments standard modulo functionality in go
func modulo(in, n int) int {
	var out int
	switch {
	case in >= 0:
		out = in % n
	case in < 0:
		out = (in%n + n) % n
	}
	return out
}
