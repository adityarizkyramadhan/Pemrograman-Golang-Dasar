package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	Hour   int
	Minute int
}

func ChangeToStandartTime(time interface{}) string {
	switch t := time.(type) {
	case string:
		testSplit := strings.Split(t, ":")
		if len(testSplit) != 2 {
			return "Invalid input"
		}
		if len(testSplit[0]) != 2 || len(testSplit[1]) != 2 {
			return "Invalid input"
		}
		if t >= "00:00" && t <= "11:59" {
			return t + " AM"
		} else if t >= "12:00" && t <= "23:59" {
			if t == "12:00" {
				return "12:00 PM"
			}
			// Kurangi 12 jam dari waktu
			hours, _ := strconv.Atoi(t[:2])
			return fmt.Sprintf("%02d:%s PM", hours-12, t[3:])

		}
	case []int:
		if len(t) != 2 {
			return "Invalid input"
		}
		if t[0] < 0 || t[0] > 23 || t[1] < 0 || t[1] > 59 {
			return "Invalid input"
		}
		hour, minute := t[0], t[1]
		if hour >= 0 && hour <= 11 {
			return fmt.Sprintf("%02d:%02d AM", hour, minute)
		} else if hour >= 12 && hour <= 23 {
			if hour == 12 {
				return fmt.Sprintf("%02d:%02d PM", hour, minute)
			}
			return fmt.Sprintf("%02d:%02d PM", hour-12, minute)
		}
	case map[string]int:
		hour, minute := t["hour"], t["minute"]
		// jika salah satu key tidak ada maka return Invalid input
		_, okHour := t["hour"]
		_, okMinute := t["minute"]
		if !okHour || !okMinute {
			return "Invalid input"
		}
		if hour >= 0 && hour <= 11 {
			return fmt.Sprintf("%02d:%02d AM", hour, minute)
		} else if hour >= 12 && hour <= 23 {
			if hour == 12 {
				return fmt.Sprintf("%02d:%02d PM", hour, minute)
			}
			return fmt.Sprintf("%02d:%02d PM", hour-12, minute)
		}
	case Time:
		if t.Hour >= 0 && t.Hour <= 11 {
			return fmt.Sprintf("%02d:%02d AM", t.Hour, t.Minute)
		} else if t.Hour >= 12 && t.Hour <= 24 {
			if t.Hour == 12 {
				return fmt.Sprintf("%02d:%02d PM", t.Hour, t.Minute)
			}
			return fmt.Sprintf("%02d:%02d PM", t.Hour-12, t.Minute)
		}
	}
	return "Invalid input"
}

func main() {
	// fmt.Println(ChangeToStandartTime("16:00"))
	// fmt.Println(ChangeToStandartTime([]int{16, 0}))
	// fmt.Println(ChangeToStandartTime(map[string]int{"hour": 16, "minute": 0}))
	fmt.Println(ChangeToStandartTime("14:00"))
}
