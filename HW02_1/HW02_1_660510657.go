// Tharachon Warintaweewat
// 660510657
// HW02_1
// 204203 Sec 001

package main

import (
	"fmt"
	"log"
	"strings"
)

// set the maximum digit length. why 36 and not 35?
const MAX = 36

func main() {
	// why are we using string?
	var n1, n2 string

	// read in two string (can be multiple lines)
	_, err := fmt.Scan(&n1, &n2)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(addition(n1, n2))
}

func addition(n1, n2 string) string {
	// this is just a skeleton and will give out wrong result
	result := []byte(strings.Repeat("x", MAX))
	len1 := len(n1)
	len2 := len(n2)
	var max_len int

	if len1 >= len2 {
		max_len = len1 - 1
	} else {
		max_len = len2 - 1
	}

	// loop from the left most digit
	i, j, k := len1-1, len2-1, MAX-1

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

		if currDigit >= 10 {
			currDigit %= 10
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
