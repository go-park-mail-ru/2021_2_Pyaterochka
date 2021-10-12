// Code generated by MockGen. DO NOT EDIT.
// Source: patreon/internal/app/repository/user (interfaces: Repository)

// Package mock_repository is a generated GoMock package.
package mock_repository_user

import (
	"patreon/internal/app/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// UserRepository is a mock of Repository interface.
type UserRepository struct {
	ctrl     *gomock.Controller
	recorder *UserRepositoryMockRecorder
}

// UserRepositoryMockRecorder is the mock recorder for UserRepository.
type UserRepositoryMockRecorder struct {
	mock *UserRepository
}

// NewUserRepository creates a new mock instance.
func NewUserRepository(ctrl *gomock.Controller) *UserRepository {
	mock := &UserRepository{ctrl: ctrl}
	mock.recorder = &UserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *UserRepository) EXPECT() *UserRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *UserRepository) Create(arg0 *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *UserRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*UserRepository)(nil).Create), arg0)
}

// FindByID mocks base method.
func (m *UserRepository) FindByID(arg0 int64) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", arg0)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *UserRepositoryMockRecorder) FindByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*UserRepository)(nil).FindByID), arg0)
}

// FindByLogin mocks base method.
func (m *UserRepository) FindByLogin(arg0 string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByLogin", arg0)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByLogin indicates an expected call of FindByLogin.
func (mr *UserRepositoryMockRecorder) FindByLogin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByLogin", reflect.TypeOf((*UserRepository)(nil).FindByLogin), arg0)
}
