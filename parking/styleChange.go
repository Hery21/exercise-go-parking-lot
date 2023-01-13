package parking

type style interface {
	StyleChange([]*Lot) *Lot
}

type DefaultFirstAvailable struct{}

type MostCapacity struct{}

type MostFreeSpace struct{}

func (d DefaultFirstAvailable) StyleChange(lots []*Lot) *Lot {
	return lots[0]
}

func (d MostCapacity) StyleChange(lots []*Lot) *Lot {
	tempMostCapacity := 0
	var mostCapacityIndex int

	for i, val := range lots {
		if val.capacity > tempMostCapacity {
			tempMostCapacity = val.capacity
			mostCapacityIndex = i
		}
	}
	return lots[mostCapacityIndex]
}

func (d MostFreeSpace) StyleChange(lots []*Lot) *Lot {
	tempMostFreeSpace := 0
	var mostFreeSpaceIndex int

	for i, val := range lots {
		if val.capacity-len(val.lot) > tempMostFreeSpace {
			tempMostFreeSpace = val.capacity - len(val.lot)
			mostFreeSpaceIndex = i
		}
	}
	return lots[mostFreeSpaceIndex]
}
