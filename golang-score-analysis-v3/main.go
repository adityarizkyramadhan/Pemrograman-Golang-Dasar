package main

import "fmt"

type School struct {
	Name    string
	Address string
	Grades  []int
}

func (s *School) AddGrade(grades ...int) {
	s.Grades = append(s.Grades, grades...)
}

func Analysis(s School) (float64, int, int) {
	if len(s.Grades) == 0 {
		return 0, 0, 0
	}
	// hitung rata-rata
	var sum int
	for _, grade := range s.Grades {
		sum += grade
	}
	avg := float64(sum) / float64(len(s.Grades))

	// cari nilai terkecil
	min := s.Grades[0]
	for _, grade := range s.Grades {
		if grade < min {
			min = grade
		}
	}

	// cari nilai terbesar
	max := s.Grades[0]
	for _, grade := range s.Grades {
		if grade > max {
			max = grade
		}
	}

	return avg, min, max
}

// gunakan untuk melakukan debugging
func main() {
	avg, min, max := Analysis(School{
		Name:    "Imam Assidiqi School",
		Address: "Jl. Imam Assidiqi",
		Grades:  []int{100, 90, 80, 70, 60},
	})

	fmt.Println(avg, min, max)
}
