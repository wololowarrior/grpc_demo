package service

import (
	"context"

	"cloudbees/dao"
	"cloudbees/model"
	cloudbeespb "cloudbees/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TicketService struct {
	datastore     dao.TicketStore
	userDataStore dao.UserStore
	seatDataStore dao.SeatStore
}

func (t *TicketService) Purchase(ctx context.Context, request *cloudbeespb.TicketRequest) (*cloudbeespb.TicketResponse, error) {
	user := model.User{Email: request.GetUser().Email,
		FirstName: request.GetUser().FirstName,
		LastName:  request.GetUser().LastName,
	}
	err := t.userDataStore.CreateUser(&user)
	if err != nil {
		return nil, status.Error(codes.Internal, "Some error occurred")
	}

	ticketDetails := request.GetTicket()
	ticket := model.Ticket{
		From:      ticketDetails.GetSource(),
		To:        ticketDetails.GetDestination(),
		Price:     ticketDetails.GetPrice(),
		UserEmail: request.GetUser().GetEmail(),
	}
	err = t.datastore.CreateTicket(&ticket)
	if err != nil {
		return nil, status.Error(codes.Internal, "Some error occurred")
	}

	ticketDetails.Id = &ticket.TicketID

	resp := &cloudbeespb.TicketResponse{
		Ticket: ticketDetails,
		User:   request.GetUser(),
	}
	return resp, nil
}

func (t *TicketService) GetReceipt(ctx context.Context, request *cloudbeespb.GetReceiptRequest) (*cloudbeespb.TicketResponse, error) {
	ticket, err := t.datastore.GetTicket(request.GetTicketId())
	if err != nil {
		if err.Error() == "not found" {
			return nil, status.Error(codes.NotFound, "Ticket not found")
		}
	}

	user1, err := t.userDataStore.GetUserByEmail(ticket.UserEmail)
	if err != nil {
		if err.Error() == "user not found" {
			return nil, status.Error(codes.NotFound, "User not found")
		}
		return nil, err
	}
	seat, err := t.seatDataStore.Get(ticket.TicketID)
	var seatDetails *cloudbeespb.Seat
	if err == nil {
		seatDetails = &cloudbeespb.Seat{
			Section: seat.Section,
			SeatNo:  seat.SeatNumber,
		}
	}
	ticketDetails := &cloudbeespb.TicketDetails{
		Id:          &ticket.TicketID,
		Source:      ticket.To,
		Destination: ticket.From,
		Price:       ticket.Price,
	}

	resp := &cloudbeespb.TicketResponse{
		Ticket: ticketDetails,
		User:   &cloudbeespb.User{Email: user1.Email, FirstName: user1.FirstName, LastName: user1.LastName},
		Seat:   seatDetails,
	}
	return resp, nil
}

func NewTicketService(store dao.TicketStore, userDataStore dao.UserStore, seatDataStore dao.SeatStore) *TicketService {
	return &TicketService{datastore: store, userDataStore: userDataStore, seatDataStore: seatDataStore}
}
