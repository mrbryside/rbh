// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package appointmentmock is a generated GoMock package.
package appointmentmock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	appointment "github.com/mrbryside/rbh/domain/interview/appointment/domain/appointment"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRepository) Create(arg0 appointment.Aggregate) (appointment.Aggregate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(appointment.Aggregate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), arg0)
}

// GetAll mocks base method.
func (m *MockRepository) GetAll(page, pageSize uint) (appointment.Aggregates, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", page, pageSize)
	ret0, _ := ret[0].(appointment.Aggregates)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockRepositoryMockRecorder) GetAll(page, pageSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockRepository)(nil).GetAll), page, pageSize)
}

// GetById mocks base method.
func (m *MockRepository) GetById(appointmentId uint) (appointment.Aggregate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", appointmentId)
	ret0, _ := ret[0].(appointment.Aggregate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockRepositoryMockRecorder) GetById(appointmentId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockRepository)(nil).GetById), appointmentId)
}

// UpdateById mocks base method.
func (m *MockRepository) UpdateById(arg0 appointment.Aggregate) (appointment.Aggregate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateById", arg0)
	ret0, _ := ret[0].(appointment.Aggregate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateById indicates an expected call of UpdateById.
func (mr *MockRepositoryMockRecorder) UpdateById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateById", reflect.TypeOf((*MockRepository)(nil).UpdateById), arg0)
}