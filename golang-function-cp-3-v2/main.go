package main

import (
	"fmt"
	"strings"
)

func FindShortestName(names string) string {
	arrayOfSpace := strings.Split(names, " ")
	arrayOfSemicolon := strings.Split(names, ";")
	arrayOfComma := strings.Split(names, ",")
	joinedArray := append(arrayOfSpace, arrayOfSemicolon...)
	joinedArray = append(joinedArray, arrayOfComma...)
	shortestName := joinedArray[0]
	for _, name := range joinedArray {
		if len(name) < len(shortestName) {
			shortestName = name
		} else if len(name) == len(shortestName) {
			if name < shortestName {
				shortestName = name
			}
		}
	}
	return shortestName
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindShortestName("Hanif Joko Tio Andi Budi Caca Hamdan")) // "Tio"
	fmt.Println(FindShortestName("Budi;Tia;Tio"))                         // "Tia"
}
