package main

import (
	"fmt"
	"strings"
	"unicode"
)

func ReverseWord(str string) string {
	result := ""
	for i := len(str) - 1; i >= 0; i-- {
		result += string(str[i])
	}
	perword := strings.Split(result, " ")
	result = ""
	for i := len(perword) - 1; i >= 0; i-- {
		allCaps := true
		allLower := true
		for j := len(perword[i]) - 1; j >= 0; j-- {
			if !unicode.IsLetter(rune(perword[i][j])) {
				continue
			}
			if !unicode.IsUpper(rune(perword[i][j])) {
				allCaps = false
			}
			if !unicode.IsLower(rune(perword[i][j])) {
				allLower = false
			}
		}
		if allCaps {
			perword[i] = strings.ToUpper(perword[i])
		} else if allLower {
			perword[i] = strings.ToLower(perword[i])
		} else {
			perword[i] = strings.ToLower(perword[i])
			perword[i] = strings.Title(perword[i])
		}
		result += perword[i] + " "
	}
	return strings.TrimSpace(result)
}

func main() {
	fmt.Println(ReverseWord("Aku Sayang Ibu"))
}
