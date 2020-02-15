package triangle

import (
	"math"
)

// Kind is the triangle type
type Kind string

//NaT is Not a Triangle
const (
	NaT = "NaT" //NaT is Not a Triangle
	Equ = "Equ" //Equ is equateral Triangle
	Iso = "Iso" //Iso is isosceles Triangle
	Sca = "Sca" //Sca is a scalene Triangle
)

// KindFromSides is the function the tests expects to call
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	switch {
	case (a == 0 || b == 0 || c == 0):
		k = NaT
	case (math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) || math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0)):
		k = NaT
	case (a+b < c || a+c < b || b+c < a):
		k = NaT
	case (a == b && a == c):
		k = Equ
	case (a == b || b == c || a == c):
		k = Iso
	case (a != b && b != c):
		k = Sca
	}
	return k
}
