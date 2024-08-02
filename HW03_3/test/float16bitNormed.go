// Tharachon Warintaweewat
// 660510657
// HW03_4
// 204203 Sec 001

package main

import (
	"fmt"
	"math"
	"strings"
)

const fracLen = 7
const expLen = 8

const BASE uint8 = 2
const DEBUG = false

// from HW02_2
const MAX_INT = 64
const DEC_PLACE = 160

func float16bitNormed(n float32) string {
	// expLen = 8, fracLen = 7127
	var bias = int(pow(2, expLen-1) - 1) // bias = 127

	if DEBUG {
		fmt.Println("Bias", bias)
	}
	var minNorm float64 = math.Pow(2, -126)             // dummy value
	var maxNorm float64 = 338953138925153547590470800371487866880  // dummy value

	sign := "0"
	if n < 0 {
		n = -n
		sign = "1"
	}

	if (float64(n) > maxNorm) || (float64(n) < minNorm) {
		if DEBUG {
			fmt.Println(n, "overflow")
		}
		return ""
	}

	e := int64(math.Floor(math.Log2(float64(n))))
	pre_exp := posIntToBaseB(e + int64(bias))
	exp := expand(pre_exp)

	// fmt.Println(int64(n))
	if n >= float32(math.Pow(2, 63)) {
		// fmt.Println("aa")
		n = n / float32(math.Pow(2, 127))
	}
	// fmt.Println(int64(n))
	var converted, frac string
	converted = floatToBaseB(float64(n))
	start := strings.Index(converted, "1") + 1
	// fmt.Println(e, converted)
	// fmt.Println("---")
	frac = converted[start : start+7]

	return sign + " " + exp + " " + frac
}

func pow(x, y int) float64 {
	return math.Pow(float64(x), float64(y))
}

//------------------------------------ HW02_2

func floatToBaseB(x float64) string {
	sign := ""

	if x < 0 { // turn negative numbers to positive
		x = -x
		sign = "-"
	}
	// split at the decimal point
	front := int64(x)
	// fmt.Println(front)
	back := x - float64(front)

	frontStr := posIntToBaseB(front)
	backStr := fractionToBaseB(back)
	// putting every part together
	converted := sign + frontStr+ "." + backStr

	return converted

}

func fractionToBaseB(x float64) string {
	// only need to implement this function
	result := []byte(strings.Repeat("x", DEC_PLACE))

	var currDigit byte

	for i := 0; i < DEC_PLACE; i++ {

		currDigit = byte(math.Floor(x*float64(2)) + float64('0'))
		if currDigit > '9' {
			currDigit = 'A' + currDigit - '9' - 1
		}
		x = math.Mod(x*float64(2), 1)

		result[i] = currDigit
	}
	return string(result)
}

func posIntToBaseB(x int64) string {
	// this function is working correctly
	if x == 0 {
		return "0"
	}

	result := []byte(strings.Repeat("x", MAX_INT))
	k := MAX_INT - 1
	var currDigit byte

	for x > 0 {
		// calculate and convert back to char
		currDigit = byte((x % int64(2)) + int64('0'))
		if currDigit > '9' {
			currDigit = 'A' + currDigit - '9' - 1
		}
		result[k] = currDigit
		x = x / int64(2)
		k--
	}
	// convert from byte array to string
	return string(result[k+1:])
}

func expand(x string) string {
	if len(x) < 8 {
		x = strings.Repeat("0", 8-len(x)) + x
	}
	return x
}
