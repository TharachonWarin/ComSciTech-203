// Tharachon Warintaweewat
// 660510657
// HW03_4
// 204203 Sec 001

package main

import (
	"strings"
)

func roundToEven(x string, bPlace uint8) string {
	x_dotn := strings.Replace(x, ".", "", 1)
	x_s := strings.Index(x, ".")
	if int(bPlace) > len(x)-x_s-1 {
		return x + strings.Repeat("0", int(bPlace)-len(x)+x_s+1)
	}
	x_b := x_dotn[x_s+int(bPlace)-1:]

	if is_half(x_b[1:]) {
		if x_b[0] == 48 {
			return roundDown(x_dotn, x_s, bPlace)
		} else {
			return roundUp(x_dotn, x_s, bPlace)
		}
	} else {
		if more_or_less(x_b[1:]) {
			return roundUp(x_dotn, x_s, bPlace)
		} else {
			return roundDown(x_dotn, x_s, bPlace)
		}
	}
}

func is_half(x string) bool {
	var count int
	for _, v := range x {
		if v == 49 {
			count++
		}
	}
	return count == 1 && x[0] == 49
}

func more_or_less(x string) bool {
	var count int
	for _, v := range x {
		if v == 49 {
			count++
		}
	}
	return count > 1 && x[0] == 49
}

func roundDown(x string, pos_point int, bPlace uint8) string {
	var result string
	for i := 0; i < pos_point+int(bPlace); i++ {
		if i == pos_point {
			result += "."
		}
		result += string(x[i])
	}
	return result
}

func roundUp(x string, pos_point int, bPlace uint8) string {
	var pre_result, result string
	pre_result = addition(string(x[:pos_point+int(bPlace)]), "1")
	if len(pre_result) > pos_point+int(bPlace) {
		pos_point++
	}
	// fmt.Println(result)
	for i := 0; i < pos_point+int(bPlace); i++ {
		if i == pos_point {
			result += "."
		}
		result += string(pre_result[i])
	}
	return result
}

func addition(n1, n2 string) string {
	// this is just a skeleton and will give out wrong result
	len1 := len(n1)
	result := []byte(strings.Repeat("x", len1+1))
	len2 := len(n2)
	var max_len int

	if len1 >= len2 {
		max_len = len1 - 1
	} else {
		max_len = len2 - 1
	}

	// loop from the left most digit
	i, j, k := len1-1, len2-1, len1

	// fmt.Println("before :", i, j, k)
	var currDigit int
	nextDigit := 0

	// loop from right most position
	for ; max_len >= 0; i, j, max_len, k = i-1, j-1, max_len-1, k-1 {

		// current digit
		currDigit = 0

		if nextDigit == 1 {
			currDigit += nextDigit
			nextDigit = 0
		}

		// add the value from the current digit of n1
		if i >= 0 {
			currDigit += int(n1[i]) - int('0')
		}

		// add the value from the current digit of n2
		if j >= 0 {
			currDigit += int(n2[j]) - int('0')
		}

		if currDigit >= 2 {
			currDigit %= 2
			nextDigit = 1
		}

		// convert back to byte (one char is called byte)
		result[k] = byte(currDigit + int('0'))
	}

	if nextDigit != 0 {

		result[k] = byte(nextDigit + int('0'))
		k -= 1
	}
	// convert array of bytes to string
	return string(result[k+1:])

}
