package main

import (
	"fmt"
	"strings"
)

func CountingLetter(text string) int {
	// unreadable letters = R, S, T, Z
	count := 0
	for _, letter := range text {
		if strings.Contains("RSTZrstz", string(letter)) {
			count++
		}
	}
	return count

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingLetter("Semangat"))
}
