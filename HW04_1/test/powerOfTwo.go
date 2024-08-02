// Tharachon Warintaweewat
// 660510657
// HW04_1
// 204203 Sec 001

package main

import (
	"strconv"
	"strings"
)

func powerOfTwo(n uint64) bool {
	bin := strconv.FormatUint(n, 2)
	bin_dec := strconv.FormatUint(n-1, 2)
	max_bitlen := is_max(len(bin), len(bin_dec))
	bin = expand(bin, max_bitlen)
	bin_dec = expand(bin_dec, max_bitlen)
	return bitwise(bin, bin_dec)
}

func expand(x string, bitlen uint8) string {
	if len(x) < int(bitlen) {
		x = strings.Repeat("0", int(bitlen)-len(x)) + x
	}
	return x
}

func is_max(n1, n2 int) uint8 {
	if n1 > n2 {
		return uint8(n1)
	}
	return uint8(n2)
}

func bitwise(x1, x2 string) bool {
	// var result bool
	var n1, n2 int
	for i := 0; i < len(x1); i++ {
		n1, _ = strconv.Atoi(string(x1[i]))
		n2, _ = strconv.Atoi(string(x2[i]))
		if (n1 & n2) == 1 {
			return false
		}
		// result = strconv.(x1[0]) && int(x2[0])
	}
	return true
}
