// Code generated by MockGen. DO NOT EDIT.
// Source: patreon/internal/app/csrf/usecase (interfaces: Usecase)

// Package mock_usecase_csrf is a generated GoMock package.
package mock_usecase_csrf

import (
	csrf_models "patreon/internal/app/csrf/csrf_models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// CsrfUsecase is a mock of Usecase interface.
type CsrfUsecase struct {
	ctrl     *gomock.Controller
	recorder *CsrfUsecaseMockRecorder
}

// CsrfUsecaseMockRecorder is the mock recorder for CsrfUsecase.
type CsrfUsecaseMockRecorder struct {
	mock *CsrfUsecase
}

// NewCsrfUsecase creates a new mock instance.
func NewCsrfUsecase(ctrl *gomock.Controller) *CsrfUsecase {
	mock := &CsrfUsecase{ctrl: ctrl}
	mock.recorder = &CsrfUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *CsrfUsecase) EXPECT() *CsrfUsecaseMockRecorder {
	return m.recorder
}

// Check mocks base method.
func (m *CsrfUsecase) Check(arg0 string, arg1 int64, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Check indicates an expected call of Check.
func (mr *CsrfUsecaseMockRecorder) Check(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*CsrfUsecase)(nil).Check), arg0, arg1, arg2)
}

// Create mocks base method.
func (m *CsrfUsecase) Create(arg0 string, arg1 int64) (csrf_models.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(csrf_models.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *CsrfUsecaseMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*CsrfUsecase)(nil).Create), arg0, arg1)
}
