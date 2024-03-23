package main

import (
	"fmt"
	"strconv"
)

func BiggestPairNumber(numbers int) int {
	stringNumber := fmt.Sprintf("%d", numbers)
	var arrayNumber []int
	for _, number := range stringNumber {
		arrayNumber = append(arrayNumber, int(number)-48)
	}
	max := 0
	numberOne := 0
	numberTwo := 0
	for i := 0; i < len(arrayNumber)-1; i++ {
		sum := arrayNumber[i] + arrayNumber[i+1]
		if sum > max {
			max = sum
			numberOne = arrayNumber[i]
			numberTwo = arrayNumber[i+1]
		}
	}
	number, _ := strconv.Atoi(fmt.Sprintf("%d%d", numberOne, numberTwo))
	return number
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BiggestPairNumber(11223344))
}
