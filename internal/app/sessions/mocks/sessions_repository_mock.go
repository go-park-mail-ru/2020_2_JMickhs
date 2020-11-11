// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package sessions_mock is a generated GoMock package.
package sessions_mock

import (
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

// AddToken mocks base method
func (m *MockRepository) AddToken(token string, ID int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToken", token, ID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddToken indicates an expected call of AddToken
func (mr *MockRepositoryMockRecorder) AddToken(token, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToken", reflect.TypeOf((*MockRepository)(nil).AddToken), token, ID)
}

// GetIDByToken mocks base method
func (m *MockRepository) GetIDByToken(token string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIDByToken", token)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIDByToken indicates an expected call of GetIDByToken
func (mr *MockRepositoryMockRecorder) GetIDByToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIDByToken", reflect.TypeOf((*MockRepository)(nil).GetIDByToken), token)
}

// DeleteSession mocks base method
func (m *MockRepository) DeleteSession(token string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSession", token)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSession indicates an expected call of DeleteSession
func (mr *MockRepositoryMockRecorder) DeleteSession(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockRepository)(nil).DeleteSession), token)
}
