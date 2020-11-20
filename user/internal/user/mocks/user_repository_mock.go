// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package user_mock is a generated GoMock package.
package user_mock

import (
	multipart "mime/multipart"
	reflect "reflect"

	models "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/user/models"
	gomock "github.com/golang/mock/gomock"
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

// GetByUserName mocks base method
func (m *MockRepository) GetByUserName(name string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUserName", name)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUserName indicates an expected call of GetByUserName
func (mr *MockRepositoryMockRecorder) GetByUserName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUserName", reflect.TypeOf((*MockRepository)(nil).GetByUserName), name)
}

// Add mocks base method
func (m *MockRepository) Add(user models.User) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", user)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add
func (mr *MockRepositoryMockRecorder) Add(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockRepository)(nil).Add), user)
}

// GetUserByID mocks base method
func (m *MockRepository) GetUserByID(ID int) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", ID)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID
func (mr *MockRepositoryMockRecorder) GetUserByID(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockRepository)(nil).GetUserByID), ID)
}

// UpdateUser mocks base method
func (m *MockRepository) UpdateUser(user models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser
func (mr *MockRepositoryMockRecorder) UpdateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockRepository)(nil).UpdateUser), user)
}

// UpdateAvatar mocks base method
func (m *MockRepository) UpdateAvatar(user models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAvatar", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAvatar indicates an expected call of UpdateAvatar
func (mr *MockRepositoryMockRecorder) UpdateAvatar(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAvatar", reflect.TypeOf((*MockRepository)(nil).UpdateAvatar), user)
}

// UpdatePassword mocks base method
func (m *MockRepository) UpdatePassword(user models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePassword indicates an expected call of UpdatePassword
func (mr *MockRepositoryMockRecorder) UpdatePassword(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*MockRepository)(nil).UpdatePassword), user)
}

// GenerateHashFromPassword mocks base method
func (m *MockRepository) GenerateHashFromPassword(password string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateHashFromPassword", password)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateHashFromPassword indicates an expected call of GenerateHashFromPassword
func (mr *MockRepositoryMockRecorder) GenerateHashFromPassword(password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateHashFromPassword", reflect.TypeOf((*MockRepository)(nil).GenerateHashFromPassword), password)
}

// CompareHashAndPassword mocks base method
func (m *MockRepository) CompareHashAndPassword(hashedPassword, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompareHashAndPassword", hashedPassword, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// CompareHashAndPassword indicates an expected call of CompareHashAndPassword
func (mr *MockRepositoryMockRecorder) CompareHashAndPassword(hashedPassword, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompareHashAndPassword", reflect.TypeOf((*MockRepository)(nil).CompareHashAndPassword), hashedPassword, password)
}

// DeleteAvatarInStore mocks base method
func (m *MockRepository) DeleteAvatarInStore(user models.User, filename string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAvatarInStore", user, filename)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAvatarInStore indicates an expected call of DeleteAvatarInStore
func (mr *MockRepositoryMockRecorder) DeleteAvatarInStore(user, filename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAvatarInStore", reflect.TypeOf((*MockRepository)(nil).DeleteAvatarInStore), user, filename)
}

// UpdateAvatarInStore mocks base method
func (m *MockRepository) UpdateAvatarInStore(file multipart.File, user *models.User, fileType string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAvatarInStore", file, user, fileType)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAvatarInStore indicates an expected call of UpdateAvatarInStore
func (mr *MockRepositoryMockRecorder) UpdateAvatarInStore(file, user, fileType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAvatarInStore", reflect.TypeOf((*MockRepository)(nil).UpdateAvatarInStore), file, user, fileType)
}