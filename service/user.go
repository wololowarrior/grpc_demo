package service

import "cloudbees/dao"

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
