package example

import (
	"testing"
)

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
