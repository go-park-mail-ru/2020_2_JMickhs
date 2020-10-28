// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package comment_mock is a generated GoMock package.
package comment_mock

import (
	commModel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/comment/models"
	paginationModel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/paginator/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUsecase is a mock of Usecase interface
type MockUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseMockRecorder
}

// MockUsecaseMockRecorder is the mock recorder for MockUsecase
type MockUsecaseMockRecorder struct {
	mock *MockUsecase
}

// NewMockUsecase creates a new mock instance
func NewMockUsecase(ctrl *gomock.Controller) *MockUsecase {
	mock := &MockUsecase{ctrl: ctrl}
	mock.recorder = &MockUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUsecase) EXPECT() *MockUsecaseMockRecorder {
	return m.recorder
}

// GetComments mocks base method
func (m *MockUsecase) GetComments(hotelID, page int) (paginationModel.PaginationModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComments", hotelID, page)
	ret0, _ := ret[0].(paginationModel.PaginationModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComments indicates an expected call of GetComments
func (mr *MockUsecaseMockRecorder) GetComments(hotelID, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComments", reflect.TypeOf((*MockUsecase)(nil).GetComments), hotelID, page)
}

// AddComment mocks base method
func (m *MockUsecase) AddComment(comment commModel.Comment) (commModel.NewRate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddComment", comment)
	ret0, _ := ret[0].(commModel.NewRate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddComment indicates an expected call of AddComment
func (mr *MockUsecaseMockRecorder) AddComment(comment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddComment", reflect.TypeOf((*MockUsecase)(nil).AddComment), comment)
}

// DeleteComment mocks base method
func (m *MockUsecase) DeleteComment(ID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteComment", ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteComment indicates an expected call of DeleteComment
func (mr *MockUsecaseMockRecorder) DeleteComment(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteComment", reflect.TypeOf((*MockUsecase)(nil).DeleteComment), ID)
}

// UpdateComment mocks base method
func (m *MockUsecase) UpdateComment(comment commModel.Comment) (commModel.NewRate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateComment", comment)
	ret0, _ := ret[0].(commModel.NewRate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateComment indicates an expected call of UpdateComment
func (mr *MockUsecaseMockRecorder) UpdateComment(comment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateComment", reflect.TypeOf((*MockUsecase)(nil).UpdateComment), comment)
}

// UpdateRating mocks base method
func (m *MockUsecase) UpdateRating(prevRate commModel.PrevRate) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRating", prevRate)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRating indicates an expected call of UpdateRating
func (mr *MockUsecaseMockRecorder) UpdateRating(prevRate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRating", reflect.TypeOf((*MockUsecase)(nil).UpdateRating), prevRate)
}

// AddRating mocks base method
func (m *MockUsecase) AddRating(comment commModel.Comment) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRating", comment)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddRating indicates an expected call of AddRating
func (mr *MockUsecaseMockRecorder) AddRating(comment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRating", reflect.TypeOf((*MockUsecase)(nil).AddRating), comment)
}
