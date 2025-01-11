package dao

import "cloudbees/model"

type TicketStore interface {
	CreateTicket(ticket *model.Ticket) error
	GetTicket(id int32) (*model.Ticket, error)
	GetTicketByUserEmail(email string) (*model.Ticket, error)
	UpdateTicket(id, seatNumber int32) (*model.Ticket, error)
	DeleteTicket(id int32) error
}
