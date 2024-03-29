// Code generated by MockGen. DO NOT EDIT.
// Source: creator.go

// Package creatormock is a generated GoMock package.
package creatormock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	creator "github.com/mrbryside/rbh/domain/interview/appointment/domain/creator"
)

// MockCreatorServicer is a mock of CreatorServicer interface.
type MockCreatorServicer struct {
	ctrl     *gomock.Controller
	recorder *MockCreatorServicerMockRecorder
}

// MockCreatorServicerMockRecorder is the mock recorder for MockCreatorServicer.
type MockCreatorServicerMockRecorder struct {
	mock *MockCreatorServicer
}

// NewMockCreatorServicer creates a new mock instance.
func NewMockCreatorServicer(ctrl *gomock.Controller) *MockCreatorServicer {
	mock := &MockCreatorServicer{ctrl: ctrl}
	mock.recorder = &MockCreatorServicerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCreatorServicer) EXPECT() *MockCreatorServicerMockRecorder {
	return m.recorder
}

// GetById mocks base method.
func (m *MockCreatorServicer) GetById(id uint) (creator.Aggregate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(creator.Aggregate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockCreatorServicerMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockCreatorServicer)(nil).GetById), id)
}
