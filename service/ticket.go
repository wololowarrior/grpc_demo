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
	// get/create user
	user := model.User{Email: request.GetUser().Email,
		FirstName: request.GetUser().FirstName,
		LastName:  request.GetUser().LastName,
	}
	err := t.userDataStore.CreateUser(&user)
	if err != nil {
		status.New(codes.Internal, "Some error occurred")
		return nil, err
	}
	// check seat availability
	// book ticket
	ticketDetails := request.GetTicket()
	ticket := model.Ticket{
		From:      ticketDetails.GetSource(),
		To:        ticketDetails.GetDestination(),
		Price:     ticketDetails.GetPrice(),
		UserEmail: request.GetUser().GetEmail(),
	}
	err = t.datastore.CreateTicket(&ticket)
	if err != nil {
		status.New(codes.Internal, "Some error occurred")
		return nil, err
	}
	// assign seat
	//seat, err := t.seatDataStore.Allocate(ticket.TicketID)
	//if err != nil {
	//	return nil, err
	//}
	// return ticket
	ticketDetails.Id = &ticket.TicketID
	//seatDetails := &cloudbeespb.Seat{
	//	Section: seat.Section,
	//	SeatNo:  seat.SeatNumber,
	//}
	resp := &cloudbeespb.TicketResponse{
		Ticket: ticketDetails,
		//Seat:   seatDetails,
		User: request.GetUser(),
	}
	return resp, nil
}

func (t *TicketService) GetReceipt(ctx context.Context, user *cloudbeespb.User) (*cloudbeespb.Receipt, error) {
	user1, err := t.userDataStore.GetUserByEmail(user.Email)
	if err != nil {
		if err.Error() == "user not found" {
			status.New(codes.NotFound, "User not found")
		}
		return nil, err
	}

	ticket, err := t.datastore.GetTicketByUserEmail(user.Email)
	if err != nil {
		if err.Error() == "ticket not found" {
			status.New(codes.NotFound, "Ticket not found")
		}
	}
	//seat, err := t.seatDataStore.Get(ticket.TicketID)

	ticketDetails := &cloudbeespb.TicketDetails{
		Id:          &ticket.TicketID,
		Source:      ticket.To,
		Destination: ticket.From,
		Price:       ticket.Price,
	}

	resp := &cloudbeespb.Receipt{
		Ticket: ticketDetails,
		User:   &cloudbeespb.User{Email: user1.Email, FirstName: user1.FirstName, LastName: user1.LastName},
	}
	return resp, nil
}

func NewTicketService(store dao.TicketStore, userDataStore dao.UserStore, seatDataStore dao.SeatStore) *TicketService {
	return &TicketService{datastore: store, userDataStore: userDataStore, seatDataStore: seatDataStore}
}
