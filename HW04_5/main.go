package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	var windowSize, bufferSize int
	beginText := false
	var inputString string

	for scanner.Scan() {
		line := scanner.Text()

		line = strings.TrimSpace(line)

		if !beginText {
			fmt.Sscanf(line, "%d %d", &windowSize, &bufferSize)
			beginText = true
			continue
		}
		inputString += scanner.Text() + "\n"

	}

	compressed := encodeLZ77(windowSize, bufferSize, inputString)

	for i, match := range compressed {

		fmt.Printf("(%d,%d,%d)", match.Offset, match.Length, match.Next)
		if i != len(compressed)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println()
}
