package main

import (
	"fmt"
)

var listOfMonth = map[int]string{
	1:  "January",
	2:  "February",
	3:  "March",
	4:  "April",
	5:  "May",
	6:  "June",
	7:  "July",
	8:  "August",
	9:  "September",
	10: "October",
	11: "November",
	12: "December",
}

func DateFormat(day, month, year int) string {
	if day < 10 {
		return fmt.Sprintf("0%d-%s-%d", day, listOfMonth[month], year)
	}
	return fmt.Sprintf("%d-%s-%d", day, listOfMonth[month], year)
}

// gunakan untuk debug
func main() {
	fmt.Println(DateFormat(1, 1, 2012))
}
