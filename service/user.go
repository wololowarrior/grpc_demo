package service

import (
	context "context"

	"cloudbees/dao"
	cloudbeespb "cloudbees/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	store           dao.UserStore
	ticketDataStore dao.TicketStore
	seatDataStore   dao.SeatStore
}

func (u *UserService) DeleteUser(ctx context.Context, user *cloudbeespb.User) (*cloudbeespb.User, error) {
	tickets, err := u.ticketDataStore.GetTicketByUserEmail(user.GetEmail())
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	seat, err := u.seatDataStore.Get(tickets.TicketID)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	err = u.seatDataStore.Delete(tickets.TicketID, seat.Section)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	err = u.ticketDataStore.DeleteTicket(tickets.TicketID)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	err = u.store.DeleteUser(user.GetEmail())
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return user, nil
}

func NewUserService(store dao.UserStore, ticketDataStore dao.TicketStore, seatDataStore dao.SeatStore) *UserService {
	return &UserService{
		store:           store,
		ticketDataStore: ticketDataStore,
		seatDataStore:   seatDataStore,
	}
}
