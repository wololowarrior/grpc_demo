package inmem

import (
	"sync"

	"cloudbees/model"
)

type SeatStore struct {
	Seating []model.Seat
	lock    sync.Mutex
}

func NewSeat() *SeatStore {
	var seatStore = SeatStore{
		Seating: make([]model.Seat, 100),
	}
	for i := 1; i <= 50; i++ {
		seatStore.Seating = append(seatStore.Seating, model.Seat{
			SeatNumber: i,
			Section:    "A",
		})
	}
	for i := 1; i <= 50; i++ {
		seatStore.Seating = append(seatStore.Seating, model.Seat{
			SeatNumber: i,
			Section:    "B",
		})
	}
	return &seatStore
}
