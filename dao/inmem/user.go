package inmem

import "cloudbees/model"

type UserStore struct {
	Store map[string]model.User
}

func NewUserStore() *UserStore {
	return &UserStore{
		Store: make(map[string]model.User),
	}
}
