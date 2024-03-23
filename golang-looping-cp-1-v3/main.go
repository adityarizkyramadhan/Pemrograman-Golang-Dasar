package main

import "fmt"

func CountingNumber(n int) float64 {
	total := 0.0
	for i := 1; i <= n; i++ {
		total += float64(i)
		if i != n {
			total += float64(i) + 0.5
		}
	}
	return total
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingNumber(10))
}
