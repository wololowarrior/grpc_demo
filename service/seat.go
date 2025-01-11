package service

import (
	"context"

	"cloudbees/dao"
	cloudbeespb "cloudbees/proto"
)

type SeatService struct {
	datastore dao.SeatStore
}

func (s *SeatService) List(ctx context.Context, request *cloudbeespb.ListSeatsRequest) (*cloudbeespb.ListSeatsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SeatService) Modify(ctx context.Context, request *cloudbeespb.ModifySeatRequest) (*cloudbeespb.AllocatedSeat, error) {
	//TODO implement me
	panic("implement me")
}

func NewSeatService(datastore dao.SeatStore) *SeatService {
	return &SeatService{datastore: datastore}
}
