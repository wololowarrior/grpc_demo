package inmem

import (
	"errors"
	"sync/atomic"

	"cloudbees/model"
)

type Tickets struct {
	ticket   map[int32]*model.Ticket
	ticketID atomic.Int32
}

func (t *Tickets) CreateTicket(ticket *model.Ticket) error {
	id := t.ticketID.Load()
	ticket.TicketID = id
	t.ticket[id] = ticket
	t.ticketID.Add(1)

	return nil
}

func (t *Tickets) GetTicket(id int32) (*model.Ticket, error) {
	ticket, ok := t.ticket[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return ticket, nil
}

func (t *Tickets) GetTicketByUserEmail(email string) (*model.Ticket, error) {
	for _, ticket := range t.ticket {
		if ticket.UserEmail == email {
			return ticket, nil
		}
	}
	return nil, errors.New("ticket not found")
}

func (t *Tickets) UpdateTicket(ticketID, seatNumber int32) (*model.Ticket, error) {
	ticket, ok := t.ticket[ticketID]
	if !ok {
		return nil, errors.New("not found")
	}
	ticket.SeatNumber = seatNumber
	return ticket, nil
}

func (t *Tickets) DeleteTicket(ticketID int32) error {
	if _, ok := t.ticket[ticketID]; !ok {
		return errors.New("not found")
	}
	delete(t.ticket, ticketID)
	return nil
}

func NewTickets() *Tickets {
	return &Tickets{
		ticket: make(map[int32]*model.Ticket),
	}
}
