package service

import (
	context "context"

	"cloudbees/dao"
	cloudbeespb "cloudbees/proto"
)

type UserService struct {
	store dao.UserStore
}

func (u UserService) DeleteUser(ctx context.Context, user *cloudbeespb.User) (*cloudbeespb.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserService(store dao.UserStore) *UserService {
	return &UserService{
		store: store,
	}
}
