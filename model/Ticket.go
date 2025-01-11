package model

type Ticket struct {
	TicketID   int32
	From       string
	To         string
	Price      float32
	UserEmail  string
	SeatNumber int32
}
