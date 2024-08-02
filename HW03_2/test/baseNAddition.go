k// 660510657
// HW03_2
// 204203 sec 001
package main

import (
	// "fmt"
	"math"
	"strings"
)

func baseNAddition(r1, r2 string, n int) string {
	decPointPos1 := strings.Split(r1, ".")
	decPointPos2 := strings.Split(r2, ".")

	if len(decPointPos1) == 1 {
		decPointPos1 = append(decPointPos1, "")
	}
	if len(decPointPos2) == 1 {
		decPointPos2 = append(decPointPos2, "")
	}
	var preresult1, preresult2, result1, result2 string
	var init = 0
	for i := len(decPointPos1) - 1; i >= 0; i-- {
		preresult1, preresult2 = editbit(decPointPos1[i], decPointPos2[i], i)
		if i == 1 {
			result2, init = addition(preresult1, preresult2, n, init, i)
		} else {
			result1, init = addition(preresult1, preresult2, n, init, i)
		}
	}
	// fmt.Println(result1 result2)
	if decPointPos1[1] == "" && decPointPos2[1] == "" {
		return result1
	}

	return result1 + "." + result2
}

func addition(n1, n2 string, base, init, cat int) (string, int) {
	len1 := len(n1)
	// this is just a skeleton and will give out wrong result
	result := []byte(strings.Repeat("0", len1+1))

	// loop from the left most digit
	i := len1 - 1
	j := len1 + 1

	var currDigit int
	nextDigit := 0 + init

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

		if currDigit >= base {
			currDigit %= base
			nextDigit = 1
		}

		// convert back to byte (one char is called byte)
		result[i+1] = byte(currDigit + int('0'))
	}
	if nextDigit != 0 {
		result[i+1] = byte(nextDigit + int('0'))
		if cat == 0 {
			j--
		}
	}
	return string(result[j:]), nextDigit
}

func editbit(n1, n2 string, cat int) (string, string) {
	diff := math.Abs(float64(len(n1) - len(n2)))
	if cat == 0 {
		if len(n2) > len(n1) {
			n1 = strings.Repeat("0", int(math.Abs(diff))) + n1
		} else if len(n1) > len(n2) {
			n2 = strings.Repeat("0", int(math.Abs(diff))) + n2
		}
	} else if cat == 1 {
		if len(n2) > len(n1) {
			n1 = n1 + strings.Repeat("0", int(math.Abs(diff)))
		} else if len(n1) > len(n2) {
			n2 = n2 + strings.Repeat("0", int(math.Abs(diff)))
		}
	}
	return n1, n2
}
