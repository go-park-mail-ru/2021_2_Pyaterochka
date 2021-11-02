// Code generated by MockGen. DO NOT EDIT.
// Source: patreon/internal/app/csrf/repository/jwt (interfaces: Repository)

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	models "patreon/internal/app/csrf/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// JwtRepository is a mock of Repository interface.
type JwtRepository struct {
	ctrl     *gomock.Controller
	recorder *JwtRepositoryMockRecorder
}

// JwtRepositoryMockRecorder is the mock recorder for JwtRepository.
type JwtRepositoryMockRecorder struct {
	mock *JwtRepository
}

// NewJwtRepository creates a new mock instance.
func NewJwtRepository(ctrl *gomock.Controller) *JwtRepository {
	mock := &JwtRepository{ctrl: ctrl}
	mock.recorder = &JwtRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *JwtRepository) EXPECT() *JwtRepositoryMockRecorder {
	return m.recorder
}

// Check mocks base method.
func (m *JwtRepository) Check(arg0 models.TokenSources, arg1 models.Token) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Check indicates an expected call of Check.
func (mr *JwtRepositoryMockRecorder) Check(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*JwtRepository)(nil).Check), arg0, arg1)
}

// Create mocks base method.
func (m *JwtRepository) Create(arg0 models.TokenSources) (models.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(models.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *JwtRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*JwtRepository)(nil).Create), arg0)
}
