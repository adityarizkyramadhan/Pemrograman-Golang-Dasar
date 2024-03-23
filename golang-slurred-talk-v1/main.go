package main

import (
	"fmt"
	"strings"
)

func SlurredTalk(words *string) {
	for i := 0; i < len(*words); i++ {
		// rubah pada S R Z kecil jadi l kecil atau S R Z besar jadi L
		if string((*words)[i]) == "S" || string((*words)[i]) == "R" || string((*words)[i]) == "Z" {
			*words = strings.Replace(*words, string((*words)[i]), "L", -1)
		}
		if string((*words)[i]) == "s" || string((*words)[i]) == "r" || string((*words)[i]) == "z" {
			*words = strings.Replace(*words, string((*words)[i]), "l", -1)
		}
	}
}

func main() {
	var words string = "Saya Steven, saya suka menggoreng telur dan suka hewan zebra"
	SlurredTalk(&words)
	fmt.Println(words)
}
