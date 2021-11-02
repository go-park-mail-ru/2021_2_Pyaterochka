// Code generated by MockGen. DO NOT EDIT.
// Source: patreon/internal/app/usecase/posts_data (interfaces: Usecase)

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	io "io"
	models "patreon/internal/app/models"
	repository_files "patreon/internal/app/repository/files"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// PostsDataUsecase is a mock of Usecase interface.
type PostsDataUsecase struct {
	ctrl     *gomock.Controller
	recorder *PostsDataUsecaseMockRecorder
}

// PostsDataUsecaseMockRecorder is the mock recorder for PostsDataUsecase.
type PostsDataUsecaseMockRecorder struct {
	mock *PostsDataUsecase
}

// NewPostsDataUsecase creates a new mock instance.
func NewPostsDataUsecase(ctrl *gomock.Controller) *PostsDataUsecase {
	mock := &PostsDataUsecase{ctrl: ctrl}
	mock.recorder = &PostsDataUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *PostsDataUsecase) EXPECT() *PostsDataUsecaseMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *PostsDataUsecase) Delete(arg0 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *PostsDataUsecaseMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*PostsDataUsecase)(nil).Delete), arg0)
}

// GetData mocks base method.
func (m *PostsDataUsecase) GetData(arg0 int64) (*models.PostData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetData", arg0)
	ret0, _ := ret[0].(*models.PostData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetData indicates an expected call of GetData.
func (mr *PostsDataUsecaseMockRecorder) GetData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetData", reflect.TypeOf((*PostsDataUsecase)(nil).GetData), arg0)
}

// LoadImage mocks base method.
func (m *PostsDataUsecase) LoadImage(arg0 io.Reader, arg1 repository_files.FileName, arg2 int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadImage", arg0, arg1, arg2)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadImage indicates an expected call of LoadImage.
func (mr *PostsDataUsecaseMockRecorder) LoadImage(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadImage", reflect.TypeOf((*PostsDataUsecase)(nil).LoadImage), arg0, arg1, arg2)
}

// LoadText mocks base method.
func (m *PostsDataUsecase) LoadText(arg0 *models.PostData) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadText", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadText indicates an expected call of LoadText.
func (mr *PostsDataUsecaseMockRecorder) LoadText(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadText", reflect.TypeOf((*PostsDataUsecase)(nil).LoadText), arg0)
}

// UpdateImage mocks base method.
func (m *PostsDataUsecase) UpdateImage(arg0 io.Reader, arg1 repository_files.FileName, arg2 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateImage", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateImage indicates an expected call of UpdateImage.
func (mr *PostsDataUsecaseMockRecorder) UpdateImage(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateImage", reflect.TypeOf((*PostsDataUsecase)(nil).UpdateImage), arg0, arg1, arg2)
}

// UpdateText mocks base method.
func (m *PostsDataUsecase) UpdateText(arg0 *models.PostData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateText", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateText indicates an expected call of UpdateText.
func (mr *PostsDataUsecaseMockRecorder) UpdateText(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateText", reflect.TypeOf((*PostsDataUsecase)(nil).UpdateText), arg0)
}
