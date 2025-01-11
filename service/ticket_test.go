package service

import (
	"context"
	"errors"
	"testing"

	mock_dao "cloudbees/dao/mock"
	"cloudbees/model"
	cloudbeespb "cloudbees/proto"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestTicketService_Purchase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTicketStore := mock_dao.NewMockTicketStore(ctrl)
	mockUserStore := mock_dao.NewMockUserStore(ctrl)
	mockSeatStore := mock_dao.NewMockSeatStore(ctrl)

	service := NewTicketService(mockTicketStore, mockUserStore, mockSeatStore)

	user := &cloudbeespb.User{
		Email:     "test@example.com",
		FirstName: "John",
		LastName:  "Doe",
	}

	ticketDetails := &cloudbeespb.TicketDetails{
		Source:      "NYC",
		Destination: "LA",
		Price:       100,
	}

	request := &cloudbeespb.TicketRequest{
		User:   user,
		Ticket: ticketDetails,
	}

	mockUserStore.EXPECT().CreateUser(gomock.Any()).Return(nil)
	mockTicketStore.EXPECT().CreateTicket(gomock.Any()).Return(nil)

	resp, err := service.Purchase(context.Background(), request)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, user, resp.User)
	assert.Equal(t, ticketDetails, resp.Ticket)
}

func TestTicketService_GetReceipt(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTicketStore := mock_dao.NewMockTicketStore(ctrl)
	mockUserStore := mock_dao.NewMockUserStore(ctrl)
	mockSeatStore := mock_dao.NewMockSeatStore(ctrl)

	service := NewTicketService(mockTicketStore, mockUserStore, mockSeatStore)

	user := &cloudbeespb.User{
		Email: "test@example.com",
	}

	mockUser := &model.User{
		Email:     "test@example.com",
		FirstName: "John",
		LastName:  "Doe",
	}

	mockTicket := &model.Ticket{
		TicketID:  123,
		From:      "NYC",
		To:        "LA",
		Price:     100,
		UserEmail: "test@example.com",
	}

	mockSeat := &model.Seat{
		Section:    "A",
		SeatNumber: 10,
	}

	mockUserStore.EXPECT().GetUserByEmail(user.Email).Return(mockUser, nil)
	mockTicketStore.EXPECT().GetTicketByUserEmail(user.Email).Return(mockTicket, nil)
	mockSeatStore.EXPECT().Get(mockTicket.TicketID).Return(mockSeat, nil)

	resp, err := service.GetReceipt(context.Background(), user)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, user.Email, resp.User.Email)
	assert.Equal(t, mockTicket.From, resp.Ticket.Source)
	assert.Equal(t, mockTicket.To, resp.Ticket.Destination)
	assert.Equal(t, mockTicket.Price, resp.Ticket.Price)
	assert.Equal(t, mockSeat.Section, resp.Seat.Section)
	assert.Equal(t, mockSeat.SeatNumber, resp.Seat.SeatNo)
}

func TestTicketService_GetTicket_UserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTicketStore := mock_dao.NewMockTicketStore(ctrl)
	mockUserStore := mock_dao.NewMockUserStore(ctrl)
	mockSeatStore := mock_dao.NewMockSeatStore(ctrl)

	service := NewTicketService(mockTicketStore, mockUserStore, mockSeatStore)

	user := &cloudbeespb.User{
		Email: "test@example.com",
	}

	mockUserStore.EXPECT().GetUserByEmail(user.Email).Return(nil, errors.New("user not found"))
	resp, err := service.GetReceipt(context.Background(), user)
	assert.Nil(t, resp)
	assert.Error(t, err)

}
