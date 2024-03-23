package main

import "fmt"

func GetTicketPrice(VIP, regular, student, day int) float32 {
	vipPrice := 30.0
	regularPrice := 20.0
	studentPrice := 10.0
	totalTicket := VIP + regular + student
	totalPrice := float64(VIP)*vipPrice + float64(regular)*regularPrice + float64(student)*studentPrice
	if totalPrice < 100 {
		return float32(totalPrice)
	}
	if day%2 != 0 {
		if totalTicket < 5 {
			return float32(totalPrice * 0.85)
		} else {
			return float32(totalPrice * 0.75)
		}
	} else {
		if totalTicket < 5 {
			return float32(totalPrice * 0.9)
		} else {
			return float32(totalPrice * 0.8)
		}
	}
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetTicketPrice(1, 1, 1, 20))
}
