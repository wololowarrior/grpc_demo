package dao

import "cloudbees/model"

type TicketStore interface {
	CreateTicket(ticket model.Ticket) (model.Ticket, error)
	GetTicket(id string) (model.Ticket, error)
}
