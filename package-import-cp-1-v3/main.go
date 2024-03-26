package main

import (
	"a21hc3NpZ25tZW50/internal"
	"fmt"
	"strings"
)

func AdvanceCalculator(calculate string) float32 {
	calculator := internal.NewCalculator(0.0)
	dataInput := strings.Split(calculate, " ")
	if len(dataInput) == 0 {
		return 0.0
	}

	var num float32
	var lastOperator string
	for _, token := range dataInput {
		switch token {
		case "+", "-":
			lastOperator = token
		case "*", "/":
			lastOperator = token
		default:
			num = parseFloat(token)
			switch lastOperator {
			case "":
				calculator.Add(num)
			case "+":
				calculator.Add(num)
			case "-":
				calculator.Subtract(num)
			case "*":
				calculator.Multiply(num)
			case "/":
				calculator.Divide(num)
			}
		}
	}
	return calculator.Result()
}

func parseFloat(s string) float32 {
	var result float32
	fmt.Sscanf(s, "%f", &result)
	return result
}

func main() {
	res := AdvanceCalculator("3 * 4 / 2 + 10 - 5")

	fmt.Println(res)
}
