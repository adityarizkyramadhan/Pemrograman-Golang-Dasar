package main

import "strconv"

func ReverseData(arr [5]int) [5]int {
	for i, j := 0, 4; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	for i := range arr {
		numberStr := strconv.Itoa(arr[i])
		var reversedNumberStr string
		for i := len(numberStr) - 1; i >= 0; i-- {
			reversedNumberStr += string(numberStr[i])
		}
		arr[i], _ = strconv.Atoi(reversedNumberStr)
	}
	return arr
}
