package dao

import "cloudbees/model"

type UserStore interface {
	CreateUser(*model.User) error
	DeleteUser(*model.User) error
}
