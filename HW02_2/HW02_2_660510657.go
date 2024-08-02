// Tharachon Warintaweewat
// 660510657
// HW02_2
// 204203 Sec 001
package main

import (
	"math"
	"strings"
)

const MAX_INT = 64
const DEC_PLACE = 6

func floatToBaseB(x float64, b uint8) string {
	sign := ""

	if x < 0 { // turn negative numbers to positive
		x = -x
		sign = "-"
	}
	// split at the decimal point
	front := int64(x)
	back := x - float64(front)

	frontStr := posIntToBaseB(front, b)
	backStr := fractionToBaseB(back, b)
	// putting every part together
	converted := sign + frontStr + "." + backStr

	return converted

}

func fractionToBaseB(x float64, b uint8) string {
	// only need to implement this function
	result := []byte(strings.Repeat("x", DEC_PLACE))

	var currDigit byte

	for i := 0; i < DEC_PLACE; i++ {
		
		currDigit = byte(math.Floor(x*float64(b)) + float64('0'))
		if currDigit > '9' {
			currDigit = 'A' + currDigit - '9' - 1
		}
		x = math.Mod(x*float64(b), 1)

		result[i] = currDigit
	}
	return string(result)
}

func posIntToBaseB(x int64, b uint8) string {
	// this function is working correctly
	if x == 0 {
		return "0"
	}

	result := []byte(strings.Repeat("x", MAX_INT))
	k := MAX_INT - 1
	var currDigit byte

	for x > 0 {
		// calculate and convert back to char
		currDigit = byte((x % int64(b)) + int64('0'))
		if currDigit > '9' {
			currDigit = 'A' + currDigit - '9' - 1
		}
		result[k] = currDigit
		x = x / int64(b)
		k--
	}
	// convert from byte array to string
	return string(result[k+1:])
}
