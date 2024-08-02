// Tharachon Warintaweewat
// 660510657
// HW03_1
// 204203 Sec 001

package test

import (
	// "fmt"
	"math"
	"strconv"
	"strings"
)

func addNSubtract(n1, n2 string, bitLen uint8) (int64, int64) {
	// max := math.Max(float64(len(n1)), float64(len(n2)))
	n1 = editbit(n1, int(bitLen))
	n2 = editbit(n2, int(bitLen))
	// fmt.Println(n1,n2)
	bitresult := (addition(n1, n2))
	result1 := twosComplToInt(bitresult)
	// fmt.Println("*", bitresult)
	bitresult = (addition(n1, additiveInverse(n2)))
	result2 := twosComplToInt(bitresult)
	// fmt.Println("&", bitresult)
	return result1, result2
}

func additiveInverse(x string) string {
	list_in := strings.Split(x, "")

	result := []byte(strings.Repeat("0", len(x)))

	var i = len(x) - 1
	var count = 0
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

	}
	return string(result)
}

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

func addition(n1, n2 string) string {
	len1 := len(n1)
	// this is just a skeleton and will give out wrong result
	result := []byte(strings.Repeat("0", len1+1))

	// loop from the left most digit
	i := len1 - 1
	j := len1 + 1

	var currDigit int
	nextDigit := 0

	// loop from right most position
	for ; i >= 0; i, j = i-1, j-1 {
		// current digit
		currDigit = 0

		if nextDigit == 1 {
			currDigit += nextDigit
			nextDigit = 0
		}

		// add the value from the current digit of n1
		currDigit += int(n1[i]) - int('0')

		// add the value from the current digit of n2
		currDigit += int(n2[i]) - int('0')

		if currDigit >= 2 {
			currDigit %= 2
			nextDigit = 1
		}

		// convert back to byte (one char is called byte)
		result[i+1] = byte(currDigit + int('0'))
	}
	// fmt.Println(j)
	if nextDigit != 0 {
		result[i+1] = byte(nextDigit + int('0'))
		
	}
	// fmt.Println(j)
	// convert array of bytes to string
	// fmt.Println("full:",string(result[:]))
	return string(result[j:])
}

func editbit(n1 string, bitlen int) string {
	if bitlen == len(n1) {
		return n1
	}
	diff := math.Abs(float64(len(n1)) - float64(bitlen))
	if bitlen > len(n1) {
		if n1[0] == 49 {
			n1 = strings.Repeat("1", int(diff)) + n1
		} else {
			n1 = strings.Repeat("0", int(diff)) + n1
		}
	} else if bitlen < len(n1) {
		return n1[int(diff):]
	}
	return n1
}
