// Tharachon Warintaweewat
// 660510657
// HW04_5
// 204203 Sec 001

package main

import (
	// "fmt"
	"strings"
)

type Triplet struct {
	Offset int
	Length int
	Next   byte
}

func encodeLZ77(windowSize int, bufferSize int, inputString string) []Triplet {
	var result []Triplet
	var index, w, b, t_w, t_b int
	var t Triplet
	// fmt.Println(inputString[20-8:20-8+2])
	for index < len(inputString) {
		if index == 0 {
			// println("////")
			t.Offset, t.Length, t.Next = 0, 0, inputString[index]
			result = append(result, t)
			index++
			continue
		}
		if windowSize < bufferSize {
			w, b = windowSize, windowSize
		} else {
			w, b = windowSize, bufferSize
		}
		if index < windowSize {
			w, b = index, index
		}
		if len(inputString)-index-1 < bufferSize {
			b = len(inputString) - index
		}

		t_w, t_b = w, b
		count, off, leng, c_off, c_len := 0, 0, 0, 0, 0
		for t_w > 0 {
			// fmt.Println("index")
			if t_w == 0 || t_b == 0 {
				break
			}
			check := strings.Contains(inputString[index-t_w:index-t_w+t_b], inputString[index:index+t_b])
			// fmt.Printf("%d: |%s|%s|          %v\noffset: %d ; lenght: %d\n", index, inputString[index-t_w:index-t_w+t_b], inputString[index:index+t_b], check, t_w, t_b)
			if check {
				c_off, c_len = t_w, t_b
				if count == 0 {
					off, leng = c_off, c_len
					count++
				} else if c_len > leng {
					off, leng = c_off, c_len
					count++
				} else if c_len == leng && c_off < off {
					off, leng = c_off, c_len
					count++
				}
			}
				if t_b > 1 {
					t_b--
				} else {
					t_w--
					if t_w < b {
						t_b = t_w
					}else{
						t_b = b
					}
				}
		}

		if count == 0 {
			t.Offset, t.Length, t.Next = 0, 0, inputString[index]
			result = append(result, t)
			index++
		} else {
			if index+leng == len(inputString) {
				t.Offset, t.Length, t.Next = off, leng, 0
			} else {
				t.Offset, t.Length, t.Next = off, leng, inputString[index+leng]
			}
			result = append(result, t)
			index += leng + 1
		}
	}
	return result
}
