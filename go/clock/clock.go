// Package clock is a utility that takes in hour and
// minute data and outputs a time based on requirements.
package clock

import (
	"fmt"
)

// Clock defines the struct that is the integer
// values of a clock.
type Clock struct {
	minutes int
}

// New takes in hour and min its and
// outputs the string clock reading they denote.
func New(h, m int) Clock {
	min := (h*60 + m) % (24 * 60)
	if min < 0 {
		min += 24 * 60
	}
	return Clock{min}
}

// Add impliments a function that adds minutes
// to the clock time and returns a string.
func (c Clock) Add(minutes int) Clock {
	c.minutes += minutes
	return New(0, c.minutes)
}

// Subtract impliments a function that subtracts minutes
// to the clock time and returns a string.
func (c Clock) Subtract(minutes int) Clock {
	c.minutes -= minutes
	return New(0, c.minutes)
}

// clockString takes in min and hour integers and
// converts them to a string expression of time.
func (c Clock) String() string {
	hour := (c.minutes / 60) % 24
	c.minutes = c.minutes % 60
	return fmt.Sprintf("%02d:%02d", hour, c.minutes)
}
