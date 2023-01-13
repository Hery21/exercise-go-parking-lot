package parking

type Attendant struct {
	assignedLot  []*Lot
	availableLot []*Lot
	parkingStyle style
}

func (a *Attendant) NotifyLotAvailable(l *Lot) {
	a.availableLot = append(a.availableLot, l)
}

func (a *Attendant) NotifyLotFull(l *Lot) {
	updatedLots := make([]*Lot, 0)

	for _, lot := range a.availableLot {
		if lot != l {
			updatedLots = append(updatedLots, lot)
		}
	}
	a.availableLot = updatedLots
}

func NewAttendant(lots []*Lot) *Attendant {
	availableLotInit := make([]*Lot, len(lots))
	copy(availableLotInit, lots)

	attend := &Attendant{lots, availableLotInit, DefaultFirstAvailable{}}

	for _, lot := range lots {
		lot.SubscribeFullNotifier(attend)
		lot.SubscribeAvailableNotifier(attend)
	}
	return attend
}

func (a *Attendant) Park(car *Car) (*Ticket, error) {
	if len(a.availableLot) == 0 {
		return nil, &ErrorParking{}
	}

	return a.parkingStyle.StyleChange(a.availableLot).Park(car)
}

func (a *Attendant) Unpark(ticket *Ticket) (*Car, error) {
	for _, val := range a.assignedLot {
		if _, ok := val.lot[ticket]; ok {
			return val.Unpark(ticket)
		}
	}
	return nil, &ErrorUnparking{}
}

func (a *Attendant) changeParkingStyle(s style) {
	a.parkingStyle = s
}
