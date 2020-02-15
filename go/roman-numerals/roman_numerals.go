package romannumerals

import "errors"

//ToRomanNumeral is the function that takes a int as input and outputs the roman numeral version
func ToRomanNumeral(in int) (string, error) {
	var out string
	numeralMap := [28]string{"M", "CM", "DCCC", "DCC", "DC", "D", "CD", "CCC", "CC", "C", "XC", "LXXX", "LXX", "LX", "L", "XL", "XXX", "XX", "X", "IX", "VIII", "VII", "VI", "V", "IV", "III", "II", "I"}
	mapArray := [28]int{1000, 900, 800, 700, 600, 500, 400, 300, 200, 100, 90, 80, 70, 60, 50, 40, 30, 20, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	if in <= 0 || in > 3000 {
		return "", errors.New("invalid Entry")
	}
	key := 0
	for in > 0 {
		if in >= mapArray[key] {
			out = out + numeralMap[key]
			in = in - mapArray[key]
		} else {
			key++
		}
	}
	return out, nil
}
