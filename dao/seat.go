package dao

import "cloudbees/model"

type SeatStore interface {
	Allocate(int32, string) (*model.Seat, error)
	List(section string) ([]*model.Seat, error)
	Modify(*model.Seat, string) (*model.Seat, error)
	Get(id int32) (*model.Seat, error)
	Delete(ticketID int32, section string) error
}
