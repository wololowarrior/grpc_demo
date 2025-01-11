package service

import (
	"context"

	"cloudbees/dao"
	cloudbeespb "cloudbees/proto"
)

type TicketService struct {
	datastore dao.TicketStore
}

func (t *TicketService) Purchase(ctx context.Context, request *cloudbeespb.TicketRequest) (*cloudbeespb.TicketResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TicketService) GetReceipt(ctx context.Context, user *cloudbeespb.User) (*cloudbeespb.TicketResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewTicketService(store dao.TicketStore) *TicketService {
	return &TicketService{datastore: store}
}
