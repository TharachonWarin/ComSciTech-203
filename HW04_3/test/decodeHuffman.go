// Tharachon Warintaweewat
// 660510657
// HW04_3
// 204203 Sec 001

package main

func decodeHuffman(encodedText string, codingTable map[string]string) string {
	var check, char, result string
	for _, v := range encodedText {
		check += string(v)
		char = codingTable[check]
		if len(char) != 0{
			// fmt.Println(check)
			check = ""
			result += char
		}
	}
	return result
}
