package service

import (
	"context"

	"cloudbees/dao"
	cloudbeespb "cloudbees/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SeatService struct {
	datastore       dao.SeatStore
	ticketDataStore dao.TicketStore
	userDataStore   dao.UserStore
}

func (s *SeatService) List(ctx context.Context, request *cloudbeespb.ListSeatsRequest) (*cloudbeespb.ListSeatsResponse, error) {
	seats, err := s.datastore.List(request.GetSection())
	if err != nil {
		if err.Error() == "not found" {
			return nil, status.Error(codes.NotFound, "Seats not found")
		}
		return nil, err
	}
	var resp []*cloudbeespb.AllocatedSeat
	for _, seat := range seats {
		ticket, err := s.ticketDataStore.GetTicket(seat.TicketID)
		if err != nil {
			continue
		}
		user, err := s.userDataStore.GetUserByEmail(ticket.UserEmail)
		resp = append(resp, &cloudbeespb.AllocatedSeat{
			Seat: &cloudbeespb.Seat{
				SeatNo:  seat.SeatNumber,
				Section: seat.Section,
			},
			User: &cloudbeespb.User{
				Email:     user.Email,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			},
		})
	}

	return &cloudbeespb.ListSeatsResponse{AllocatedSeat: resp}, nil
}

func (s *SeatService) Modify(ctx context.Context, request *cloudbeespb.ModifySeatRequest) (*cloudbeespb.AllocatedSeat, error) {
	ticket, err := s.ticketDataStore.GetTicket(request.GetTicketId())
	if err != nil {
		if err.Error() == "not found" {
			return nil, status.Error(codes.NotFound, "Ticket not found")
		}
	}
	oldSeat, err := s.datastore.Get(ticket.TicketID)
	newSeat, err := s.datastore.Modify(oldSeat, request.GetSection())
	if err != nil {
		return nil, err
	}
	allocatedSeat := &cloudbeespb.AllocatedSeat{
		Seat: &cloudbeespb.Seat{
			SeatNo:  newSeat.SeatNumber,
			Section: newSeat.Section,
		},
	}
	return allocatedSeat, nil
}

func (s *SeatService) Allocate(ctx context.Context, request *cloudbeespb.AllocationRequest) (*cloudbeespb.AllocatedSeat, error) {
	ticket, err := s.ticketDataStore.GetTicket(request.GetTicketId())
	if err != nil {
		if err.Error() == "not found" {
			return nil, status.Error(codes.NotFound, "Ticket not found")
		}
	}
	if ticket.SeatNumber != 0 {
		return nil, status.Error(codes.AlreadyExists, "Seat already allocated")
	}
	seat, err := s.datastore.Allocate(request.GetTicketId(), request.GetSection())
	if err != nil {
		return nil, err
	}

	_, err = s.ticketDataStore.UpdateTicket(seat.TicketID, seat.SeatNumber)
	if err != nil {
		if err.Error() == "not found" {
			return nil, status.Error(codes.NotFound, "Ticket not found")
		}
	}

	allocatedSeat := &cloudbeespb.AllocatedSeat{
		Seat: &cloudbeespb.Seat{
			SeatNo:  seat.SeatNumber,
			Section: seat.Section,
		},
	}
	return allocatedSeat, nil
}

func NewSeatService(datastore dao.SeatStore, ticketDataStore dao.TicketStore, userDataStore dao.UserStore) *SeatService {
	return &SeatService{datastore: datastore, ticketDataStore: ticketDataStore, userDataStore: userDataStore}
}
