package main

import "fmt"

func ExchangeCoin(amount int) []int {
	if amount == 0 {
		return []int{}
	}
	var coins = []int{1000, 500, 200, 100, 50, 20, 10, 5, 1}
	var result []int
	for _, coin := range coins {
		if amount == 0 {
			break
		}
		if amount >= coin {
			count := amount / coin
			amount = amount % coin
			for i := 0; i < count; i++ {
				result = append(result, coin)
			}
		}
	}
	return result
}

func main() {
	fmt.Println(ExchangeCoin(0))
}
