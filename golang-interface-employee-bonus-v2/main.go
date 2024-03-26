package main

type Employee interface {
	GetBonus() float64
}

type Junior struct {
	Name         string
	BaseSalary   int
	WorkingMonth int
}

type Senior struct {
	Name            string
	BaseSalary      int
	WorkingMonth    int
	PerformanceRate float64
}

type Manager struct {
	Name             string
	BaseSalary       int
	WorkingMonth     int
	PerformanceRate  float64
	BonusManagerRate float64
}

func (j Junior) GetBonus() float64 {
	// Bonus untuk Junior adalah 1 x BaseSalary x prorata lama kerja
	return float64(j.BaseSalary) * float64(j.WorkingMonth) / 12
}

func (s Senior) GetBonus() float64 {
	// Bonus untuk Senior adalah 2 x BaseSalary x prorata lama kerja + (PerformanceRate x BaseSalary)
	return 2*float64(s.BaseSalary)*float64(s.WorkingMonth)/12 + s.PerformanceRate*float64(s.BaseSalary)
}

func (m Manager) GetBonus() float64 {
	// 2 x BaseSalary x prorata lama kerja + (PerformanceRate x BaseSalary) + (BonusManagerRate x BaseSalary)
	return 2*float64(m.BaseSalary)*float64(m.WorkingMonth)/12 + m.PerformanceRate*float64(m.BaseSalary) + m.BonusManagerRate*float64(m.BaseSalary)
}

func EmployeeBonus(employee Employee) float64 {
	return employee.GetBonus()
}

func TotalEmployeeBonus(employees []Employee) float64 {
	var total float64
	for _, employee := range employees {
		total += employee.GetBonus()
	}
	return total
}
