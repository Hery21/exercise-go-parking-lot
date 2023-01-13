package parking

type LotFullNotifier interface {
	NotifyLotFull(l *Lot)
}

type LotAvailableNotifier interface {
	NotifyLotAvailable(l *Lot)
}

type Lot struct {
	lot                 map[*Ticket]*Car
	capacity            int
	subscriberFull      []*LotFullNotifier
	subscriberAvailable []*LotAvailableNotifier
}

type Car interface {
}

type Ticket interface {
}

type ErrorParking struct {
}

type ErrorUnparking struct {
}

type ErrorCarParked struct {
}

func (e *ErrorParking) Error() string {
	return "No available position"
}

func (e *ErrorUnparking) Error() string {
	return "Unrecognized parking ticket"
}

func (e *ErrorCarParked) Error() string {
	return "Car is already parked"
}

func CreateParkingLot(capacity int) *Lot {
	return &Lot{make(map[*Ticket]*Car), capacity, []*LotFullNotifier{}, []*LotAvailableNotifier{}}
}

func (l *Lot) SubscribeFullNotifier(subscriber LotFullNotifier) {
	l.subscriberFull = append(l.subscriberFull, &subscriber)
}

func (l *Lot) SubscribeAvailableNotifier(subscriber LotAvailableNotifier) {
	l.subscriberAvailable = append(l.subscriberAvailable, &subscriber)
}

func (l *Lot) HasFreeSpace() bool {
	return len(l.lot) < l.capacity
}

func (l *Lot) IsParked(car *Car) bool {
	for _, val := range l.lot {
		if val == car {
			return true
		}
	}
	return false
}

func (l *Lot) Park(car *Car) (*Ticket, error) {
	if l.IsParked(car) {
		return nil, &ErrorCarParked{}
	}

	if !l.HasFreeSpace() {
		return nil, &ErrorParking{}
	} else {
		var ticket = new(Ticket)

		l.lot[ticket] = car

		if !l.HasFreeSpace() {
			for _, subscriber := range l.subscriberFull {
				(*subscriber).NotifyLotFull(l)
			}
		}

		return ticket, nil
	}
}

func (l *Lot) Unpark(ticket *Ticket) (*Car, error) {
	if car, ok := l.lot[ticket]; ok {
		delete(l.lot, ticket)

		if l.HasFreeSpace() {
			for _, subscriber := range l.subscriberAvailable {
				(*subscriber).NotifyLotAvailable(l)
			}
		}

		return car, nil
	} else {
		return nil, &ErrorUnparking{}
	}
}
