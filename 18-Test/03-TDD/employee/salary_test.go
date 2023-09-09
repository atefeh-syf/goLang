package employee

import (
	"testing"
)

// base salary -> 10000000
// extra -> base / 30 / 7 * 1.4 * hours = 800000
// karaneh -> 5000000
// insurance -> b + e * 0.03 = 324000.0
// tax -> b + e * 0.09 = 972000

func TestExtraSalaryCalculator(t *testing.T) {
	//Arrange
	baseSalary := 10000000.0
	extraHours := 12.0
	want := 800000.0

	//Act
	got := ExtraSalaryCalculator(baseSalary, extraHours)

	//Assert
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestInsuranceCalculator(t *testing.T) {
	//Arrange
	baseSalary := 10000000.0
	extraSalary := 800000.0
	want := 324000.0

	//Act
	got := InsuranceCalculator(baseSalary, extraSalary)

	//Assert
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestTaxCalculator(t *testing.T) {
	//Arrange
	baseSalary := 10000000.0
	extraSalary := 800000.0
	want := 324000.0

	//Act
	got := InsuranceCalculator(baseSalary, extraSalary)

	//Assert
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFinalSalaryCalculator(t *testing.T) {
	//Arrange
	baseSalary := 10000000.0
	extraHours := 12.0
	karaneh := 5000000.0
	want := 14504000.0

	//Act
	got := FinalSalaryCalculator(baseSalary, extraHours, karaneh)

	//Assert
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
