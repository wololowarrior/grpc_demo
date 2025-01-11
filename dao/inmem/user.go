package inmem

import (
	"errors"

	"cloudbees/model"
)

type UserStore struct {
	Store map[string]*model.User
}

func (u *UserStore) CreateUser(user *model.User) error {
	_, ok := u.Store[user.Email]
	if !ok {
		u.Store[user.Email] = user
	}
	return nil
}

func (u *UserStore) DeleteUser(user *model.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserStore) GetUserByEmail(email string) (*model.User, error) {
	user, ok := u.Store[email]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func NewUserStore() *UserStore {
	return &UserStore{
		Store: make(map[string]*model.User),
	}
}
