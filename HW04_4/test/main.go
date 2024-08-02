package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		line = strings.TrimSpace(line)

		tripletsStr := strings.Split(line, ", ")
		var triplets []Triplet
		for _, tStr := range tripletsStr {
			var t Triplet
			// fmt.Println(tStr)
			_, err := fmt.Sscanf(tStr, "(%d,%d,%d)", &t.Offset, &t.Length, &t.Next)
			if err != nil {
				if err != io.EOF {
					log.Fatal(err)
				}
				break
			}

			triplets = append(triplets, t)
		}
		// fmt.Print(triplets[0].Next)
		decoded := decodeLZ77(triplets)
		fmt.Print(decoded)

	}
}
