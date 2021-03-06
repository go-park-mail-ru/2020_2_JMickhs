// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package sessions_mock is a generated GoMock package.
package sessions_mock

import (
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

// AddToken mocks base method
func (m *MockUsecase) AddToken(ID int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToken", ID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddToken indicates an expected call of AddToken
func (mr *MockUsecaseMockRecorder) AddToken(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToken", reflect.TypeOf((*MockUsecase)(nil).AddToken), ID)
}

// GetIDByToken mocks base method
func (m *MockUsecase) GetIDByToken(token string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIDByToken", token)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIDByToken indicates an expected call of GetIDByToken
func (mr *MockUsecaseMockRecorder) GetIDByToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIDByToken", reflect.TypeOf((*MockUsecase)(nil).GetIDByToken), token)
}

// DeleteSession mocks base method
func (m *MockUsecase) DeleteSession(token string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSession", token)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSession indicates an expected call of DeleteSession
func (mr *MockUsecaseMockRecorder) DeleteSession(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockUsecase)(nil).DeleteSession), token)
}
