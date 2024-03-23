package main

import "fmt"

// hello World => d_l_r_o_W o_l_l_e_H
func ReverseString(str string) string {
	result := ""
	for i := len(str) - 1; i >= 0; i-- {
		result += string(str[i])
		if string(str[i]) == " " {
			continue
		}
		if i == 0 {
			break
		}
		if string(str[i-1]) == " " {
			continue
		}
		result += "_"
	}
	return result
}

// gunakan untuk melakukan debug
func main() {
	result := ReverseString("Hello World")
	fmt.Println(result)
	fmt.Println(result == "d_l_r_o_W o_l_l_e_H")
}
