package dao

import "cloudbees/model"

type UserStore interface {
	CreateUser(*model.User) error
	DeleteUser(email string) error
	GetUserByEmail(string) (*model.User, error)
}
