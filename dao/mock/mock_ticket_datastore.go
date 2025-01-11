// Code generated by MockGen. DO NOT EDIT.
// Source: dao/ticket.go

// Package mock_dao is a generated GoMock package.
package mock_dao

import (
	model "cloudbees/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTicketStore is a mock of TicketStore interface.
type MockTicketStore struct {
	ctrl     *gomock.Controller
	recorder *MockTicketStoreMockRecorder
}

// MockTicketStoreMockRecorder is the mock recorder for MockTicketStore.
type MockTicketStoreMockRecorder struct {
	mock *MockTicketStore
}

// NewMockTicketStore creates a new mock instance.
func NewMockTicketStore(ctrl *gomock.Controller) *MockTicketStore {
	mock := &MockTicketStore{ctrl: ctrl}
	mock.recorder = &MockTicketStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTicketStore) EXPECT() *MockTicketStoreMockRecorder {
	return m.recorder
}

// CreateTicket mocks base method.
func (m *MockTicketStore) CreateTicket(ticket *model.Ticket) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTicket", ticket)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTicket indicates an expected call of CreateTicket.
func (mr *MockTicketStoreMockRecorder) CreateTicket(ticket interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTicket", reflect.TypeOf((*MockTicketStore)(nil).CreateTicket), ticket)
}

// DeleteTicket mocks base method.
func (m *MockTicketStore) DeleteTicket(id int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTicket", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTicket indicates an expected call of DeleteTicket.
func (mr *MockTicketStoreMockRecorder) DeleteTicket(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTicket", reflect.TypeOf((*MockTicketStore)(nil).DeleteTicket), id)
}

// GetTicket mocks base method.
func (m *MockTicketStore) GetTicket(id int32) (*model.Ticket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTicket", id)
	ret0, _ := ret[0].(*model.Ticket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTicket indicates an expected call of GetTicket.
func (mr *MockTicketStoreMockRecorder) GetTicket(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTicket", reflect.TypeOf((*MockTicketStore)(nil).GetTicket), id)
}

// GetTicketByUserEmail mocks base method.
func (m *MockTicketStore) GetTicketByUserEmail(email string) (*model.Ticket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTicketByUserEmail", email)
	ret0, _ := ret[0].(*model.Ticket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTicketByUserEmail indicates an expected call of GetTicketByUserEmail.
func (mr *MockTicketStoreMockRecorder) GetTicketByUserEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTicketByUserEmail", reflect.TypeOf((*MockTicketStore)(nil).GetTicketByUserEmail), email)
}

// UpdateTicket mocks base method.
func (m *MockTicketStore) UpdateTicket(id, seatNumber int32) (*model.Ticket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTicket", id, seatNumber)
	ret0, _ := ret[0].(*model.Ticket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTicket indicates an expected call of UpdateTicket.
func (mr *MockTicketStoreMockRecorder) UpdateTicket(id, seatNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTicket", reflect.TypeOf((*MockTicketStore)(nil).UpdateTicket), id, seatNumber)
}