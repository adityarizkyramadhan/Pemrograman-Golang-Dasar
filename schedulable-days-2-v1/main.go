package main

func SchedulableDays(villager [][]int) []int {
	var schedule []int
	for i := 0; i < len(villager); i++ {
		isAllAvailable := false
		for j := 0; j < len(villager[i]); j++ {
			isAllAvailable = true
			for k := 0; k < len(villager); k++ {
				if i == k {
					continue
				}
				isAvailable := false
				for l := 0; l < len(villager[k]); l++ {
					if villager[i][j] == villager[k][l] {
						isAvailable = true
						break
					}
				}
				if !isAvailable {
					isAllAvailable = false
					break
				}
			}
			if isAllAvailable {
				schedule = append(schedule, villager[i][j])
			}
		}
	}
	if len(schedule) == 0 {
		return []int{}
	}
	return schedule
}
