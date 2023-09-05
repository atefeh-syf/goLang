package example

import (
	"testing"
)

type testCase struct {
	RoomType   string
	Nights     int
	Persons    int
	FinalPrice int
}

var testCases = []testCase{
	{"standard", 2, 2, 610400},
	{"double", 2, 1, 479600},
	{"suite", 2, 2, 1526000},
}

func TestCalculateRoomPrice(t *testing.T) {
	// Arrange
	expectedRoomPrice := 610400
	nights := 2
	persons := 2

	// Act
	actualRoomPrice := CalculateRoomPrice("standard", nights, persons)

	// Assert
	if actualRoomPrice != expectedRoomPrice {
		t.Errorf("Expected %v, got %v", expectedRoomPrice, actualRoomPrice)
	}
}

func TestCalculateRoomPrice1(t *testing.T) {

	for _, item := range testCases {
		got := CalculateRoomPrice(item.RoomType, item.Nights, item.Persons)
		if got != item.FinalPrice {
			t.Errorf("Expected %v, got %v", item.FinalPrice, got)
		}
	}
}

func FuzzCalculateRoomPrice1(f *testing.F) {
	f.Fuzz(func(t *testing.T, roomType string, nights int, personCount int) {
		CalculateRoomPrice(roomType, nights, personCount)
	})
}
