package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Setiap data akan berisi nama, umur, alamat, tinggi badan dan status pernikahan dengan format "[name];[age];[address];[height];[is_married]"
func PopulationData(data []string) []map[string]interface{} {
	result := make([]map[string]interface{}, 0)
	if len(data) == 0 {
		return result
	}
	for _, person := range data {
		personData := make(map[string]interface{})
		splitPerson := strings.Split(person, ";")
		// Jika data tidak ada pada array maka skip
		if splitPerson[0] != "" {
			personData["name"] = splitPerson[0]
		}
		if splitPerson[1] != "" {
			age, _ := strconv.Atoi(splitPerson[1])
			personData["age"] = age
		}
		if splitPerson[2] != "" {
			personData["address"] = splitPerson[2]
		}
		if splitPerson[3] != "" {
			// convert string to float
			height, _ := strconv.ParseFloat(splitPerson[3], 64)
			personData["height"] = height
		}
		if splitPerson[4] != "" {
			// convert string to boolean
			isMarried, _ := strconv.ParseBool(splitPerson[4])
			personData["isMarried"] = isMarried
		}
		result = append(result, personData)
	}
	return result
}

func main() {
	testData := []string{
		"Budi;23;Jakarta;160.42;true",
		"Joko;30;Bandung;;true",
		"Susi;25;Bogor;;false",
		"Santi;27;Jakarta;160;",
		"Rudi;23;Jakarta;161.1;",
		"Rudi;23;Jakarta;166.5;false",
		"Rudi;23;Jakarta;;",
	}
	result := PopulationData(testData)
	fmt.Println(result)
}
