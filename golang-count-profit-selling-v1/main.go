package main

import "fmt"

func CountProfit(data [][][2]int) []int {
	// cek jika data kosong
	if len(data) == 0 {
		return []int{}
	}
	result := make([]int, len(data[0]))
	for _, cabang := range data {
		for indexBulan, bulan := range cabang {
			result[indexBulan] += bulan[0] - bulan[1]
		}
	}
	if len(result) == 0 {
		return []int{}
	}
	return result
}

func main() {
	// [[[1000, 800],[700, 500]], [[1000, 800],[900, 200]]]
	arrayInput := [][][2]int{}
	result := CountProfit(arrayInput)
	fmt.Println(result)
}
