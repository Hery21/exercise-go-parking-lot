package parking_test

import (
	. " hery-ciaputra/exercise-go-parking-lot/parking"
	" hery-ciaputra/exercise-go-parking-lot/parking/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLot_Park(t *testing.T) {
	t.Run("Should return a parking ticket when park a car", func(t *testing.T) {
		car := new(Car)

		parkingLot := CreateParkingLot(1)
		ticket, _ := parkingLot.Park(car)

		assert.NotNil(t, ticket)
	})

	t.Run("should return error when lot capacity is full", func(t *testing.T) {
		car1 := new(Car)
		car2 := new(Car)
		parkingLot := CreateParkingLot(1)

		_, _ = parkingLot.Park(car1)
		ticket, err := parkingLot.Park(car2)

		assert.Nil(t, ticket)
		assert.ErrorContains(t, err, "No available position")
	})

	t.Run("should return error when try to park parked car", func(t *testing.T) {
		car := new(Car)
		parkingLot := CreateParkingLot(1)

		_, _ = parkingLot.Park(car)
		_, err := parkingLot.Park(car)

		assert.ErrorContains(t, err, "Car is already parked")
	})

	t.Run("Should call NotifyLotFull when all lot is full", func(t *testing.T) {
		car := new(Car)

		parkingLot := CreateParkingLot(1)

		mocking := &mocks.LotFullNotifier{}
		mocking.On("NotifyLotFull", parkingLot).Return()

		parkingLot.SubscribeFullNotifier(mocking)
		_, _ = parkingLot.Park(car)

		mocking.AssertNumberOfCalls(t, "NotifyLotFull", 1)
	})

	t.Run("Should not call NotifyLotFull for not a subscriberFull when all lot is full", func(t *testing.T) {
		car := new(Car)

		parkingLot := CreateParkingLot(1)

		notASubscriber := &mocks.LotFullNotifier{}
		notASubscriber.On("NotifyLotFull", parkingLot).Return()
		_, _ = parkingLot.Park(car)

		notASubscriber.AssertNumberOfCalls(t, "NotifyLotFull", 0)
	})
}

func TestLot_Unpark(t *testing.T) {
	t.Run("Should return Nil when unparking car with invalid ticket", func(t *testing.T) {
		var ticket *Ticket

		parkingLot := CreateParkingLot(1)

		actual, _ := parkingLot.Unpark(ticket)

		assert.Nil(t, actual)
	})

	t.Run("Should return a car when unparking car with valid ticket", func(t *testing.T) {
		var ticket *Ticket
		car := new(Car)

		expected := car
		parkingLot := CreateParkingLot(1)
		ticket, _ = parkingLot.Park(car)

		actual, _ := parkingLot.Unpark(ticket)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should return 2 correct car when unparking 2 car", func(t *testing.T) {
		car1 := new(Car)
		car2 := new(Car)

		parkingLot := CreateParkingLot(2)

		park1, _ := parkingLot.Park(car1)
		park2, _ := parkingLot.Park(car2)

		actualCar1, _ := parkingLot.Unpark(park1)
		actualCar2, _ := parkingLot.Unpark(park2)

		assert.Same(t, car1, actualCar1)
		assert.Same(t, car2, actualCar2)

	})

	t.Run("Should return error when ticket is wrong or no ticket", func(t *testing.T) {
		parkingLot := CreateParkingLot(1)

		car := new(Car)

		_, _ = parkingLot.Park(car)
		var wrongParkingTicket = new(Ticket)

		_, err := parkingLot.Unpark(wrongParkingTicket)

		assert.ErrorContains(t, err, "Unrecognized parking ticket")
	})

	t.Run("Should return error when ticket is used", func(t *testing.T) {
		parkingLot := CreateParkingLot(1)

		car := new(Car)

		parkTicket, _ := parkingLot.Park(car)

		_, _ = parkingLot.Unpark(parkTicket)
		_, err := parkingLot.Unpark(parkTicket)

		assert.ErrorContains(t, err, "Unrecognized parking ticket")
	})

	t.Run("Should call NotifyLotAvailable when lot is freed", func(t *testing.T) {
		car := new(Car)

		parkingLot := CreateParkingLot(1)

		mocking := &mocks.LotAvailableNotifier{}
		mocking.On("NotifyLotAvailable", parkingLot).Return()

		parkingLot.SubscribeAvailableNotifier(mocking)
		ticket, _ := parkingLot.Park(car)
		_, _ = parkingLot.Unpark(ticket)

		mocking.AssertNumberOfCalls(t, "NotifyLotAvailable", 1)
	})
}
