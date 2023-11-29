// Code generated by MockGen. DO NOT EDIT.
// Source: ./authorization.go

// Package mockgen is a generated GoMock package.
package mockgen

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	claim "github.com/mrbryside/rbh/pkg/claim"
)

// MockJwtAuthorizer is a mock of JwtAuthorizer interface.
type MockJwtAuthorizer struct {
	ctrl     *gomock.Controller
	recorder *MockJwtAuthorizerMockRecorder
}

// MockJwtAuthorizerMockRecorder is the mock recorder for MockJwtAuthorizer.
type MockJwtAuthorizerMockRecorder struct {
	mock *MockJwtAuthorizer
}

// NewMockJwtAuthorizer creates a new mock instance.
func NewMockJwtAuthorizer(ctrl *gomock.Controller) *MockJwtAuthorizer {
	mock := &MockJwtAuthorizer{ctrl: ctrl}
	mock.recorder = &MockJwtAuthorizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJwtAuthorizer) EXPECT() *MockJwtAuthorizerMockRecorder {
	return m.recorder
}

// GenerateToken mocks base method.
func (m *MockJwtAuthorizer) GenerateToken(userID uint, role string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", userID, role)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockJwtAuthorizerMockRecorder) GenerateToken(userID, role interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockJwtAuthorizer)(nil).GenerateToken), userID, role)
}

// ParseToken mocks base method.
func (m *MockJwtAuthorizer) ParseToken(tokenString string) (*claim.CustomClaims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", tokenString)
	ret0, _ := ret[0].(*claim.CustomClaims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockJwtAuthorizerMockRecorder) ParseToken(tokenString interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockJwtAuthorizer)(nil).ParseToken), tokenString)
}