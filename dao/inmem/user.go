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

func (u *UserStore) DeleteUser(email string) error {
	if _, ok := u.Store[email]; !ok {
		return errors.New("not found")
	}
	delete(u.Store, email)
	return nil
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
