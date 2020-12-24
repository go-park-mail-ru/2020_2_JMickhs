// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package recommend_mock is a generated GoMock package.
package recommend_mock

import (
	recommModels "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/recommendation/models"
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

// GetHotelsRecommendations mocks base method
func (m *MockRepository) GetHotelsRecommendations(UserID int) ([]recommModels.HotelRecommend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHotelsRecommendations", UserID)
	ret0, _ := ret[0].([]recommModels.HotelRecommend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHotelsRecommendations indicates an expected call of GetHotelsRecommendations
func (mr *MockRepositoryMockRecorder) GetHotelsRecommendations(UserID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHotelsRecommendations", reflect.TypeOf((*MockRepository)(nil).GetHotelsRecommendations), UserID)
}

// GetRecommendationRows mocks base method
func (m *MockRepository) GetRecommendationRows(UserID int, hotelIDs []int) ([]recommModels.RecommendMatrixRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecommendationRows", UserID, hotelIDs)
	ret0, _ := ret[0].([]recommModels.RecommendMatrixRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecommendationRows indicates an expected call of GetRecommendationRows
func (mr *MockRepositoryMockRecorder) GetRecommendationRows(UserID, hotelIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecommendationRows", reflect.TypeOf((*MockRepository)(nil).GetRecommendationRows), UserID, hotelIDs)
}

// GetHotelByIDs mocks base method
func (m *MockRepository) GetHotelByIDs(hotelIDs []int64) ([]recommModels.HotelRecommend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHotelByIDs", hotelIDs)
	ret0, _ := ret[0].([]recommModels.HotelRecommend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHotelByIDs indicates an expected call of GetHotelByIDs
func (mr *MockRepositoryMockRecorder) GetHotelByIDs(hotelIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHotelByIDs", reflect.TypeOf((*MockRepository)(nil).GetHotelByIDs), hotelIDs)
}

// UpdateUserRecommendations mocks base method
func (m *MockRepository) UpdateUserRecommendations(userID int, hotelIDs []int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserRecommendations", userID, hotelIDs)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserRecommendations indicates an expected call of UpdateUserRecommendations
func (mr *MockRepositoryMockRecorder) UpdateUserRecommendations(userID, hotelIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserRecommendations", reflect.TypeOf((*MockRepository)(nil).UpdateUserRecommendations), userID, hotelIDs)
}

// GetUsersComments mocks base method
func (m *MockRepository) GetUsersComments(userID int) ([]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersComments", userID)
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsersComments indicates an expected call of GetUsersComments
func (mr *MockRepositoryMockRecorder) GetUsersComments(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersComments", reflect.TypeOf((*MockRepository)(nil).GetUsersComments), userID)
}

// CheckRecommendationExist mocks base method
func (m *MockRepository) CheckRecommendationExist(userID int) (recommModels.Recommendation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckRecommendationExist", userID)
	ret0, _ := ret[0].(recommModels.Recommendation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckRecommendationExist indicates an expected call of CheckRecommendationExist
func (mr *MockRepositoryMockRecorder) CheckRecommendationExist(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckRecommendationExist", reflect.TypeOf((*MockRepository)(nil).CheckRecommendationExist), userID)
}

// AddInSearchHistory mocks base method
func (m *MockRepository) AddInSearchHistory(UserID int, pattern string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddInSearchHistory", UserID, pattern)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddInSearchHistory indicates an expected call of AddInSearchHistory
func (mr *MockRepositoryMockRecorder) AddInSearchHistory(UserID, pattern interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddInSearchHistory", reflect.TypeOf((*MockRepository)(nil).AddInSearchHistory), UserID, pattern)
}

// GetHotelsFromHistory mocks base method
func (m *MockRepository) GetHotelsFromHistory(UserID int, hotelIDs []int64) ([]recommModels.HotelRecommend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHotelsFromHistory", UserID, hotelIDs)
	ret0, _ := ret[0].([]recommModels.HotelRecommend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHotelsFromHistory indicates an expected call of GetHotelsFromHistory
func (mr *MockRepositoryMockRecorder) GetHotelsFromHistory(UserID, hotelIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHotelsFromHistory", reflect.TypeOf((*MockRepository)(nil).GetHotelsFromHistory), UserID, hotelIDs)
}
