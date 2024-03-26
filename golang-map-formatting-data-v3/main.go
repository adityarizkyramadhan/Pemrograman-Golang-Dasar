package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO: answer here
// output: map[account: ["John Doe", "Jane Doe"]]
func ChangeOutput(data []string) map[string][]string {
	result := make(map[string][]string)

	for _, v := range data {
		splitData := strings.Split(v, "-")
		key := splitData[0]
		value := splitData[3]
		index, _ := strconv.Atoi(splitData[1])
		// buat key baru jika key belum ada
		if _, ok := result[key]; !ok {
			result[key] = make([]string, 0)
		}
		// jika array belum ada di index tersebut, maka buat array baru
		if len(result[key]) <= index {
			result[key] = append(result[key], value)
		} else {
			result[key][index] += " " + value
		}
	}

	return result
}

// bisa digunakan untuk melakukan debug
func main() {
	data := []string{"account-0-first-John", "account-0-last-Doe", "account-1-first-Jane", "account-1-last-Doe", "address-0-first-Jaksel", "address-0-last-Jakarta", "address-1-first-Bandung", "address-1-last-Jabar", "phone-0-first-081234567890", "phone-1-first-081234567891"}
	res := ChangeOutput(data)

	fmt.Println(res)
}
