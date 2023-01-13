package parking

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAttendant(t *testing.T) {
	t.Run("Should create a new attendant", func(t *testing.T) {
		parkingLot := CreateParkingLot(1)
		lots := []*Lot{parkingLot}
		parkingAttendant := NewAttendant(lots)
		assert.NotNil(t, parkingAttendant)
	})
}

func TestAttendant_Park(t *testing.T) {
	t.Run("Should return a parking ticket when park a car", func(t *testing.T) {
		car := new(Car)

		parkingLot := CreateParkingLot(1)
		lots := []*Lot{parkingLot}
		parkingAttendant := NewAttendant(lots)
		ticket, err := parkingAttendant.Park(car)

		assert.NotNil(t, ticket)
		assert.Nil(t, err)
	})

	t.Run("Should return error when lot capacity is full", func(t *testing.T) {
		car1 := new(Car)
		car2 := new(Car)
		parkingLot := CreateParkingLot(1)

		lots := []*Lot{parkingLot}
		parkingAttendant := NewAttendant(lots)
		_, _ = parkingAttendant.Park(car1)
		ticket, err := parkingAttendant.Park(car2)

		expectedError := &ErrorParking{}

		assert.Nil(t, ticket)
		assert.Exactly(t, expectedError, err)
	})

	t.Run("Should return error when every lot is full given more than 1 lot", func(t *testing.T) {
		lot1 := CreateParkingLot(1)
		lot2 := CreateParkingLot(1)
		var lots = []*Lot{lot1, lot2}

		parkingAttendant := NewAttendant(lots)
		_, _ = parkingAttendant.Park(new(Car))
		_, _ = parkingAttendant.Park(new(Car))
		_, err := parkingAttendant.Park(new(Car))

		expectedError := &ErrorParking{}

		assert.Equal(t, expectedError, err)
	})

	t.Run("Should return error when try to park but car is parked", func(t *testing.T) {
		car := new(Car)

		parkingLot := CreateParkingLot(2)
		lots := []*Lot{parkingLot}
		parkingAttendant := NewAttendant(lots)

		_, _ = parkingAttendant.Park(car)
		_, err := parkingAttendant.Park(car)

		expectedError := &ErrorCarParked{}

		assert.Same(t, expectedError, err)
	})

	t.Run("Should park where most capacity is when given style MostCapacity", func(t *testing.T) {
		car := new(Car)

		parkingLot1 := CreateParkingLot(5)
		parkingLot2 := CreateParkingLot(2)

		lots := []*Lot{parkingLot1, parkingLot2}
		parkingAttendant := NewAttendant(lots)

		parkingAttendant.changeParkingStyle(MostCapacity{})

		_, _ = parkingAttendant.Park(car)

		assert.True(t, parkingLot1.IsParked(car))
		assert.False(t, parkingLot2.IsParked(car))
	})

	t.Run("Should park where most free space is when given style MostFreeSpace", func(t *testing.T) {
		car1 := new(Car)
		car2 := new(Car)
		car3 := new(Car)

		parkingLot1 := CreateParkingLot(1)
		parkingLot2 := CreateParkingLot(2)

		lots := []*Lot{parkingLot1, parkingLot2}
		parkingAttendant := NewAttendant(lots)

		parkingAttendant.changeParkingStyle(MostFreeSpace{})

		_, _ = parkingAttendant.Park(car1)
		_, _ = parkingAttendant.Park(car2)
		_, _ = parkingAttendant.Park(car3)

		assert.True(t, parkingLot2.IsParked(car1))
		assert.True(t, parkingLot1.IsParked(car2))
		assert.True(t, parkingLot2.IsParked(car3))
	})

	t.Run("Should park on first available lot when not given style", func(t *testing.T) {
		car1 := new(Car)
		car2 := new(Car)
		car3 := new(Car)

		parkingLot1 := CreateParkingLot(1)
		parkingLot2 := CreateParkingLot(2)

		lots := []*Lot{parkingLot1, parkingLot2}
		parkingAttendant := NewAttendant(lots)

		_, _ = parkingAttendant.Park(car1)
		_, _ = parkingAttendant.Park(car2)
		_, _ = parkingAttendant.Park(car3)

		assert.True(t, parkingLot1.IsParked(car1))
		assert.True(t, parkingLot2.IsParked(car2))
		assert.True(t, parkingLot2.IsParked(car3))
	})
}

func TestAttendant_Unpark(t *testing.T) {
	t.Run("Should return Nil when unparking car with invalid ticket", func(t *testing.T) {
		parkingLot := CreateParkingLot(2)
		lots := []*Lot{parkingLot}
		parkingAttendant := NewAttendant(lots)
		var ticket *Ticket

		actual, _ := parkingAttendant.Unpark(ticket)

		assert.Nil(t, actual)
	})

	t.Run("Should return a car when unparking car with valid ticket", func(t *testing.T) {
		car := new(Car)

		expected := car
		parkingLot := CreateParkingLot(1)
		lots := []*Lot{parkingLot}
		parkingAttendant := NewAttendant(lots)
		ticket, _ := parkingAttendant.Park(car)

		actual, _ := parkingAttendant.Unpark(ticket)

		assert.Same(t, expected, actual)
	})

	t.Run("Should return 2 correct car when unparking 2 car", func(t *testing.T) {
		car1 := new(Car)
		car2 := new(Car)

		parkingLot := CreateParkingLot(2)
		lots := []*Lot{parkingLot}
		parkingAttendant := NewAttendant(lots)

		ticket1, _ := parkingAttendant.Park(car1)
		ticket2, _ := parkingAttendant.Park(car2)

		actualCar1, _ := parkingAttendant.Unpark(ticket1)
		actualCar2, _ := parkingAttendant.Unpark(ticket2)

		assert.Same(t, car1, actualCar1)
		assert.Same(t, car2, actualCar2)

	})

	t.Run("Should return error when ticket is wrong or no ticket", func(t *testing.T) {
		parkingLot := CreateParkingLot(1)
		lots := []*Lot{parkingLot}
		parkingAttendant := NewAttendant(lots)

		car := new(Car)

		_, _ = parkingAttendant.Park(car)
		var wrongParkingTicket = new(Ticket)

		_, actualErr := parkingAttendant.Unpark(wrongParkingTicket)

		expectedErr := &ErrorUnparking{}

		assert.Exactly(t, expectedErr, actualErr)
	})

	t.Run("Should return error when ticket is used", func(t *testing.T) {
		parkingLot := CreateParkingLot(1)
		lots := []*Lot{parkingLot}
		parkingAttendant := NewAttendant(lots)

		car := new(Car)

		parkTicket, _ := parkingAttendant.Park(car)

		_, _ = parkingAttendant.Unpark(parkTicket)
		_, actualErr := parkingAttendant.Unpark(parkTicket)

		expectedErr := &ErrorUnparking{}

		assert.Exactly(t, expectedErr, actualErr)
	})
}
