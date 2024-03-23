package main

import "fmt"

func EmailInfo(email string) string {
	domain := ""
	tld := ""

	atIndex := -1
	firstDotIndex := -1
	for i := 0; i < len(email); i++ {
		if email[i] == '@' {
			atIndex = i
		} else if email[i] == '.' {
			firstDotIndex = i
			break
		}
	}
	domain = email[atIndex+1 : firstDotIndex]
	tld = email[firstDotIndex+1:]
	return fmt.Sprintf("Domain: %s dan TLD: %s", domain, tld)
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(EmailInfo("admin@yahoo.co.id"))
}
