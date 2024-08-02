// Tharachon Warintaweewat
// 660510657
// HW04_4
// 204203 Sec 001

package main

type Triplet struct {
	Offset int
	Length int
	Next   byte
}

func decodeLZ77(matches []Triplet) string {
	decodedString := []byte{}
	var len_str int
	for _, v := range matches {
		len_str = len(decodedString)
		if v.Offset == 0 && v.Length == 0 {
			// println("aa")
			decodedString = append(decodedString, v.Next)
			continue
		}
		// println("bb"+string(byte(10))+"bb")
		for _, v := range decodedString[len_str-v.Offset:len_str-v.Offset+v.Length] {
			decodedString = append(decodedString, v)
		}
		if v.Next != 0 {
			decodedString = append(decodedString, v.Next)			
		}
	}
	return string(decodedString) + "  "
}
