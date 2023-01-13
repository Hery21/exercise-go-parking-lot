package parking

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultFirstAvailable_SelectStyle(t *testing.T) {
	t.Run("Should get first available style as when chosen DefaultFirstAvailable", func(t *testing.T) {
		parkingLot1 := CreateParkingLot(1)
		parkingLot2 := CreateParkingLot(1)

		parkingAttendant := NewAttendant([]*Lot{parkingLot1, parkingLot2})
		parkingAttendant.changeParkingStyle(DefaultFirstAvailable{})

		assert.Equal(t, parkingAttendant.parkingStyle, DefaultFirstAvailable{})
	})
}

func TestMostCapacity_SelectStyle(t *testing.T) {
	t.Run("Should get most capacity style as when chosen MostCapacity", func(t *testing.T) {
		parkingLot1 := CreateParkingLot(1)
		parkingLot2 := CreateParkingLot(2)

		parkingAttendant := NewAttendant([]*Lot{parkingLot1, parkingLot2})
		parkingAttendant.changeParkingStyle(MostCapacity{})

		assert.Equal(t, parkingAttendant.parkingStyle, MostCapacity{})
	})
}

func TestMostFreeSpace_SelectStyle(t *testing.T) {
	t.Run("Should get most free space style as when chosen MostFreeSpace", func(t *testing.T) {
		parkingLot1 := CreateParkingLot(1)
		parkingLot2 := CreateParkingLot(1)

		parkingAttendant := NewAttendant([]*Lot{parkingLot1, parkingLot2})
		parkingAttendant.changeParkingStyle(MostFreeSpace{})

		assert.Equal(t, parkingAttendant.parkingStyle, MostFreeSpace{})
	})
}
