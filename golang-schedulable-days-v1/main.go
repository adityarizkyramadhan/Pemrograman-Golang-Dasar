package main

func SchedulableDays(date1 []int, date2 []int) []int {
	var sameDay []int
	for _, day1 := range date1 {
		for _, day2 := range date2 {
			if day1 == day2 {
				sameDay = append(sameDay, day1)
			}
		}
	}
	if len(sameDay) == 0 {
		return []int{}
	}
	return sameDay
}
