package main

import (
	"fmt"
)

func PhoneNumberChecker(number string, result *string) {
	if len(number) < 10 || len(number) > 13 {
		*result = "invalid"
		return
	}
	// 0811 - 0815 / 62811 - 62815 => Telkomsel
	// 0816 - 0819 / 62816 - 62819 => Indosat
	// 0821 - 0823 / 62821 - 62823 => XL
	// 0827 - 0829 / 62827 - 62829 => Tri
	// 0852 - 0853 / 62852 - 62853 => AS
	// 0881 - 0888 / 62881 - 62888 => Smartfren
	// cek apakah kode negara atau 0
	if number[0:2] == "62" {
		number = "0" + number[2:]
	}
	// cek apakah telkomsel
	if number[0:4] == "0811" || number[0:4] == "0812" || number[0:4] == "0813" || number[0:4] == "0814" || number[0:4] == "0815" {
		*result = "Telkomsel"
		return
	}
	// cek apakah indosat
	if number[0:4] == "0816" || number[0:4] == "0817" || number[0:4] == "0818" || number[0:4] == "0819" {
		*result = "Indosat"
		return
	}
	// cek apakah xl
	if number[0:4] == "0821" || number[0:4] == "0822" || number[0:4] == "0823" {
		*result = "XL"
		return
	}
	// cek apakah tri
	if number[0:4] == "0827" || number[0:4] == "0828" || number[0:4] == "0829" {
		*result = "Tri"
		return
	}
	// cek apakah as
	if number[0:4] == "0852" || number[0:4] == "0853" {
		*result = "AS"
		return
	}
	// cek apakah smartfren
	if number[0:4] == "0881" || number[0:4] == "0882" || number[0:4] == "0883" || number[0:4] == "0884" || number[0:4] == "0885" || number[0:4] == "0886" || number[0:4] == "0887" || number[0:4] == "0888" {
		*result = "Smartfren"
		return
	}
	*result = "invalid"
}

func main() {
	// bisa digunakan untuk pengujian test case
	var number = "6281311111111"
	var result string

	PhoneNumberChecker(number, &result)
	fmt.Println(result)
}
