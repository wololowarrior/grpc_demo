package dao

import "cloudbees/model"

type SeatStore interface {
	Allocate() (*model.Seat, error)
	List(section string) (*[]model.Seat, error)
	Modify(old, new *model.Seat) error
}
