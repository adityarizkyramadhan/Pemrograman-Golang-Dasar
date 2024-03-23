package main

import (
	"fmt"
	"strings"
)

var vowels = map[rune]bool{
	'a': true,
	'i': true,
	'u': true,
	'e': true,
	'o': true,
}

func CountVowelConsonant(str string) (int, int, bool) {
	str = strings.ToLower(str)
	var vowelCount, consonantCount int
	for _, char := range str {
		if char == ' ' {
			continue
		}
		if char < 'a' || char > 'z' {
			continue
		}
		if _, ok := vowels[char]; ok {
			vowelCount++
		} else {
			consonantCount++
		}
	}
	var isEmpty bool
	if vowelCount == 0 || consonantCount == 0 {
		isEmpty = true
	}
	return vowelCount, consonantCount, isEmpty
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountVowelConsonant("SEMANGAT PAGI, itu kata orang yang baru datang ke rumahku"))
}
