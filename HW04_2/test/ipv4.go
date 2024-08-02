// Tharachon Warintaweewat
// 660510657
// HW04_2
// 204203 Sec 001

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ipv4Decode(ipUint uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", ipUint>>24, ipUint>>16<<24>>24, ipUint>>8<<24>>24, ipUint<<24>>24)
}

func ipv4Encode(ipString string) uint32 {
	var result uint32
	group := strings.Split(ipString, ".")
	sec1, _ := strconv.ParseUint(group[0], 10, 8)
	sec2, _ := strconv.ParseUint(group[1], 10, 8)
	sec3, _ := strconv.ParseUint(group[2], 10, 8)
	sec4, _ := strconv.ParseUint(group[3], 10, 8)
	result = uint32(sec1<<24 + sec2<<16 + sec3<<8 + sec4)
	return result
}
