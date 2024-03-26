package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Define a struct to hold order data
type Order struct {
	firstName string
	lastName  string
	price     float32
	location  string
}

// Global variable to store the locations and their allowed delivery days
var locations = map[string][]string{
	"JKT": {"senin", "selasa", "rabu", "kamis", "jumat", "sabtu"},
	"BDG": {"rabu", "kamis", "sabtu"},
	"BKS": {"selasa", "kamis", "jumat"},
	"DPK": {"senin", "selasa"},
}

// Function to calculate the total delivery cost for a given day
func DeliveryOrder(data []string, day string) map[string]float32 {
	result := make(map[string]float32)

	// Parse the data and calculate the total cost for each order
	for _, item := range data {
		orderInfo := strings.Split(item, ":")
		firstName := orderInfo[0]
		lastName := orderInfo[1]
		price, _ := strconv.ParseFloat(orderInfo[2], 32)
		location := orderInfo[3]

		// Check if the order is allowed to be delivered on the given day
		if isDeliveryDay(location, day) {
			key := fmt.Sprintf("%s-%s", firstName, lastName)
			if _, found := result[key]; !found {
				result[key] = 0
			}
			result[key] += float32(price)
		}
	}

	// Apply the admin fee based on the day
	adminFee := getAdminFee(day)
	for key, value := range result {
		result[key] = value + (value * adminFee)
	}

	return result
}

// Function to check if the delivery to a given location is allowed on the given day
func isDeliveryDay(location, day string) bool {
	allowedDays, found := locations[location]
	if !found {
		return false
	}
	for _, allowedDay := range allowedDays {
		if allowedDay == day {
			return true
		}
	}
	return false
}

// Function to get the admin fee based on the given day
func getAdminFee(day string) float32 {
	if day == "senin" || day == "rabu" || day == "jumat" {
		return 0.1 // 10%
	}
	return 0.05 // 5%
}

func main() {
	data := []string{
		"Anggi:Anggraini:20000:DPK",
		"Andi:Sukirman:15000:JKT",
	}

	day := "selasa"

	deliveryData := DeliveryOrder(data, day)

	fmt.Println(deliveryData)
}
