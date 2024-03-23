package main

import (
	"fmt"
	"strings"
)

func FindSimilarData(input string, data ...string) string {
	var saveData []string
	for _, v := range data {
		if strings.Contains(v, input) {
			saveData = append(saveData, v)
		}
	}
	return strings.Join(saveData, ",")
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindSimilarData("iphone", "laptop", "iphone 13", "iphone 12", "iphone 12 pro"))
}
