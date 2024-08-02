// Tharachon Warintaweewat
// 660510657
// HW01_2
// Problem A
// 204203 Sec 002

// THIS TEMPLATE IS OPTIONAL
// YOU ARE NOT REQUIRED TO USE IT

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	DEBUG = false
)

func checkPair(w1, w2 string) bool {
	var len_1 = len(w1)
	var len_2 = len(w2)
	var slice_1 = strings.Split(w1, "")
	var slice_2 = strings.Split(w2, "")
	var count int

	if math.Abs(float64(len_1)-float64(len_2)) >= 2 {
		return false
	}

	if len_1 == len_2 {
		if strings.Compare(w1, w2) == 0 {
			return false
		} else {
			for i := 0; i < int(len_1); i++ {
				if strings.Compare(slice_1[i], slice_2[i]) != 0 {
					count += 1
				}
			}
			return count < 2
		}
	} else if len_1 > len_2 {
		for i := 0; i < int(len_1); i++ {
			if count == 2 {
				return false
			}

			if i == int(len_2) && count == 0 {
				return true
			}

			if strings.Compare(slice_1[i], slice_2[i-count]) != 0 {
				count += 1
			}
		}
		return count < 2
	} else {
		for i := 0; i < int(len_2); i++ {
			if count == 2 {
				return false
			}

			if i == int(len_1) && count == 0 {
				return true
			}
			// fmt.Println(slice_1[i-count], slice_2[i])
			if strings.Compare(slice_1[i-count], slice_2[i]) != 0 {
				count += 1
			}
		}
		return count < 2
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	numCases := scanToInt(scanner.Text())

	for i := 0; i < numCases; i++ {
		scanner.Scan()
		numWords := scanToInt(scanner.Text())
		if DEBUG {
			fmt.Println(numWords)
		}
		wordArray := make([]string, numWords)
		for j := 0; j < numWords; j++ {
			scanner.Scan()
			wordArray[j] = scanner.Text()
			if DEBUG {
				fmt.Println(wordArray[j])
			}
		}

		j := 1

		for ; j < numWords; j++ {
			if !checkPair(wordArray[j-1], wordArray[j]) {
				fmt.Println("F")
				break
			}
		}

		if j == numWords {
			fmt.Println("T")
		}
	}
}

func scanToInt(s string) int {
	var n int
	fmt.Sscan(s, &n)
	return n
}

// var name string
// var name = "name"
// name := "name"
