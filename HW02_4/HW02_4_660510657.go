// Tharachon Warintaweewat
// 660510657
// HW02_4
// 204203 Sec 001
package main

import (
	"math"
	"strings"
)

func additiveInverse(x string) (string, int64) {
	list_in := strings.Split(x, "")

	result := []byte(strings.Repeat("x", len(x)))

	var i = len(x) - 1
	var count = 0
	var sum = 0
	for ; i >= 0; i = i - 1 {
		if list_in[i] == "1" && count == 0 {
			result[i] = '1'
			count += 1
		} else if list_in[i] == "0" && count == 0 {
			result[i] = '0'
		} else if list_in[i] == "1" {
			result[i] = '0'
		} else {
			result[i] = '1'
		}

		if i == 0 {
			sum += int(-1) * int(result[i]-'0') * int(math.Pow(float64(2), float64(len(x)-1)))
		} else {
			sum += int(result[i]-'0') * int(math.Pow(float64(2), float64(len(x))-float64(i)-1))
		}
	}

	return string(result), int64(sum)
}
