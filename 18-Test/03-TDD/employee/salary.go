package employee

import "math"

// base salary -> 10000000
// extra -> base / 30 / 7 * 1.4 * hours
// karaneh -> 5000000
// insurance -> b + e * 0.03
// tax -> b + e * 0.09

func ExtraSalaryCalculator(baseSalary float64, extraHours float64) float64 {
	return math.Ceil(baseSalary / 30 / 7 * 1.4 * extraHours)
}

func InsuranceCalculator(baseSalary float64, extraSalary float64) float64 {
	return math.Ceil((baseSalary + extraSalary) * 0.03)
}

func TaxCalculator(baseSalary float64, extraSalary float64) float64 {
	return math.Ceil((baseSalary + extraSalary) * 0.09)
}

func FinalSalaryCalculator(baseSalary float64, extraHours float64, karaneh float64) float64 {
	extraAmount := ExtraSalaryCalculator(baseSalary, extraHours)
	insuranceAmount := InsuranceCalculator(baseSalary, extraAmount)
	taxAmount := TaxCalculator(baseSalary, extraAmount)

	return baseSalary + extraAmount + karaneh - insuranceAmount - taxAmount
}
