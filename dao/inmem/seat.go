package inmem

import (
	"errors"

	"cloudbees/model"
)

type SeatStore struct {
	Seating []*model.Seat
}

func (s *SeatStore) Allocate(ticketID int32, section string) (*model.Seat, error) {
	if section != "" {
		for _, seat := range s.Seating {
			if seat.TicketID == -1 && seat.Section == section {
				seat.TicketID = ticketID
				return seat, nil
			}
		}
	} else {
		for _, seat := range s.Seating {
			if seat.TicketID == -1 {
				seat.TicketID = ticketID
				return seat, nil
			}
		}
	}

	return nil, errors.New("all booked")
}

func (s *SeatStore) List(section string) ([]*model.Seat, error) {
	var seats []*model.Seat
	for _, seat := range s.Seating {
		if seat.Section == section && seat.TicketID != -1 {
			seats = append(seats, seat)
		}
	}
	if len(seats) == 0 {
		return nil, errors.New("not found")
	}
	return seats, nil
}

func (s *SeatStore) Modify(old *model.Seat, newSection string) (*model.Seat, error) {
	// assign new seat
	newSeat, err := s.Allocate(old.TicketID, newSection)
	if err != nil {
		return nil, err
	}
	// remove new seat
	err = s.Delete(old.TicketID, old.Section)
	if err != nil {
		return nil, err
	}
	return newSeat, nil
}

func (s *SeatStore) Delete(ticketID int32, section string) error {
	for _, seat := range s.Seating {
		if seat.TicketID == ticketID && seat.Section == section {
			seat.TicketID = -1
			return nil
		}
	}
	return errors.New("seat not found")
}

func (s *SeatStore) Get(ticketID int32) (*model.Seat, error) {
	for _, seat := range s.Seating {
		if seat.TicketID == ticketID {
			return seat, nil
		}
	}

	return nil, errors.New("not found")
}

func NewSeat() *SeatStore {
	var seatStore = SeatStore{
		Seating: make([]*model.Seat, 0),
	}
	for i := 1; i <= 5; i++ {
		seatStore.Seating = append(seatStore.Seating, &model.Seat{
			SeatNumber: int32(i),
			Section:    "A",
			TicketID:   -1,
		})
	}
	for i := 1; i <= 5; i++ {
		seatStore.Seating = append(seatStore.Seating, &model.Seat{
			SeatNumber: int32(i),
			Section:    "B",
			TicketID:   -1,
		})
	}
	return &seatStore
}
