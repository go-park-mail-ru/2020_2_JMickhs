// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	models "github.com/go-park-mail-ru/2020_2_JMickhs/internal/comment/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// GetComments mocks base method
func (m *MockRepository) GetComments(hotelID, StartID int) ([]models.FullCommentInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComments", hotelID, StartID)
	ret0, _ := ret[0].([]models.FullCommentInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComments indicates an expected call of GetComments
func (mr *MockRepositoryMockRecorder) GetComments(hotelID, StartID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComments", reflect.TypeOf((*MockRepository)(nil).GetComments), hotelID, StartID)
}

// AddComment mocks base method
func (m *MockRepository) AddComment(comment models.Comment) (models.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddComment", comment)
	ret0, _ := ret[0].(models.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddComment indicates an expected call of AddComment
func (mr *MockRepositoryMockRecorder) AddComment(comment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddComment", reflect.TypeOf((*MockRepository)(nil).AddComment), comment)
}

// DeleteComment mocks base method
func (m *MockRepository) DeleteComment(ID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteComment", ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteComment indicates an expected call of DeleteComment
func (mr *MockRepositoryMockRecorder) DeleteComment(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteComment", reflect.TypeOf((*MockRepository)(nil).DeleteComment), ID)
}

// UpdateComment mocks base method
func (m *MockRepository) UpdateComment(comment models.Comment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateComment", comment)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateComment indicates an expected call of UpdateComment
func (mr *MockRepositoryMockRecorder) UpdateComment(comment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateComment", reflect.TypeOf((*MockRepository)(nil).UpdateComment), comment)
}
