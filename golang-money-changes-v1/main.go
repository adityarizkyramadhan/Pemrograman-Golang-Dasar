package main

import "fmt"

type Product struct {
	Name  string
	Price int
	Tax   int
}

func MoneyChanges(amount int, products []Product) []int {
	result := make([]int, 0)
	for _, product := range products {
		total := product.Price + product.Tax
		amount -= total
	}
	fmt.Println("Kembalian", amount)
	for amount > 0 {
		if amount >= 1000 {
			amount -= 1000
			result = append(result, 1000)
		} else if amount >= 500 {
			amount -= 500
			result = append(result, 500)
		} else if amount >= 200 {
			amount -= 200
			result = append(result, 200)
		} else if amount >= 100 {
			amount -= 100
			result = append(result, 100)
		} else if amount >= 50 {
			amount -= 50
			result = append(result, 50)
		} else if amount >= 20 {
			amount -= 20
			result = append(result, 20)
		} else if amount >= 10 {
			amount -= 10
			result = append(result, 10)
		} else if amount >= 5 {
			amount -= 5
			result = append(result, 5)
		} else if amount >= 2 {
			amount -= 2
			result = append(result, 2)
		} else if amount >= 1 {
			amount -= 1
			result = append(result, 1)
		}
	}
	return result
}

func main() {
	products := []Product{
		Product{Name: "Buku", Price: 2000, Tax: 200},
		Product{Name: "Pensil", Price: 1000, Tax: 100},
	}
	amount := 10000
	result := MoneyChanges(amount, products)
	fmt.Println(result)
}
