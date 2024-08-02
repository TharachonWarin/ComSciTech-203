// Tharachon Warintaweewat
// 660510657
// HW02_3
// 204203 Sec 001
package main

import (
	"math"
	"strconv"
	"strings"
)

func twosComplToInt(x string) int64 {

	var result int64 = 0

	list_x := strings.Split(x, "")

	len_str := len(x) - 1

	for i, v := range list_x {
		digit, _ := strconv.Atoi(v)
		if i == 0 {
			result += (-1) * int64(digit) * int64(math.Pow(float64(2), float64(len_str)))
		} else {
			result += int64(digit) * int64(math.Pow(float64(2), float64(len_str)-float64(i)))
		}
	}

	return result
}
