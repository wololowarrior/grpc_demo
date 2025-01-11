package inmem

import "cloudbees/model"

type Tickets struct {
	ticket []model.Ticket
}

func (t Tickets) CreateTicket(ticket model.Ticket) (model.Ticket, error) {
	//TODO implement me
	panic("implement me")
}

func (t Tickets) GetTicket(id string) (model.Ticket, error) {
	//TODO implement me
	panic("implement me")
}

func NewTickets() *Tickets {
	return &Tickets{
		ticket: make([]model.Ticket, 0),
	}
}
