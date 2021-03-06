// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package hotels_mock is a generated GoMock package.
package hotels_mock

import (
	commModel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/comment/models"
	hotelmodel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/hotels/models"
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

// GetHotelByID mocks base method
func (m *MockRepository) GetHotelByID(ID int) (hotelmodel.Hotel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHotelByID", ID)
	ret0, _ := ret[0].(hotelmodel.Hotel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHotelByID indicates an expected call of GetHotelByID
func (mr *MockRepositoryMockRecorder) GetHotelByID(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHotelByID", reflect.TypeOf((*MockRepository)(nil).GetHotelByID), ID)
}

// FetchHotels mocks base method
func (m *MockRepository) FetchHotels(filter hotelmodel.HotelFiltering, pattern string, offset int) ([]hotelmodel.Hotel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchHotels", filter, pattern, offset)
	ret0, _ := ret[0].([]hotelmodel.Hotel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchHotels indicates an expected call of FetchHotels
func (mr *MockRepositoryMockRecorder) FetchHotels(filter, pattern, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchHotels", reflect.TypeOf((*MockRepository)(nil).FetchHotels), filter, pattern, offset)
}

// BuildQueryForCommentsPercent mocks base method
func (m *MockRepository) BuildQueryForCommentsPercent(filter *hotelmodel.HotelFiltering, param string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildQueryForCommentsPercent", filter, param)
	ret0, _ := ret[0].(string)
	return ret0
}

// BuildQueryForCommentsPercent indicates an expected call of BuildQueryForCommentsPercent
func (mr *MockRepositoryMockRecorder) BuildQueryForCommentsPercent(filter, param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildQueryForCommentsPercent", reflect.TypeOf((*MockRepository)(nil).BuildQueryForCommentsPercent), filter, param)
}

// BuildQueryToFetchHotel mocks base method
func (m *MockRepository) BuildQueryToFetchHotel(filter *hotelmodel.HotelFiltering) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildQueryToFetchHotel", filter)
	ret0, _ := ret[0].(string)
	return ret0
}

// BuildQueryToFetchHotel indicates an expected call of BuildQueryToFetchHotel
func (mr *MockRepositoryMockRecorder) BuildQueryToFetchHotel(filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildQueryToFetchHotel", reflect.TypeOf((*MockRepository)(nil).BuildQueryToFetchHotel), filter)
}

// CheckRateExist mocks base method
func (m *MockRepository) CheckRateExist(UserID, HotelID int) (commModel.FullCommentInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckRateExist", UserID, HotelID)
	ret0, _ := ret[0].(commModel.FullCommentInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckRateExist indicates an expected call of CheckRateExist
func (mr *MockRepositoryMockRecorder) CheckRateExist(UserID, HotelID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckRateExist", reflect.TypeOf((*MockRepository)(nil).CheckRateExist), UserID, HotelID)
}

// GetHotelsPreview mocks base method
func (m *MockRepository) GetHotelsPreview(pattern string) ([]hotelmodel.HotelPreview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHotelsPreview", pattern)
	ret0, _ := ret[0].([]hotelmodel.HotelPreview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHotelsPreview indicates an expected call of GetHotelsPreview
func (mr *MockRepositoryMockRecorder) GetHotelsPreview(pattern interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHotelsPreview", reflect.TypeOf((*MockRepository)(nil).GetHotelsPreview), pattern)
}

// GetHotelsByRadius mocks base method
func (m *MockRepository) GetHotelsByRadius(latitude, longitude, radius string) ([]hotelmodel.Hotel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHotelsByRadius", latitude, longitude, radius)
	ret0, _ := ret[0].([]hotelmodel.Hotel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHotelsByRadius indicates an expected call of GetHotelsByRadius
func (mr *MockRepositoryMockRecorder) GetHotelsByRadius(latitude, longitude, radius interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHotelsByRadius", reflect.TypeOf((*MockRepository)(nil).GetHotelsByRadius), latitude, longitude, radius)
}

// GetMiniHotelByID mocks base method
func (m *MockRepository) GetMiniHotelByID(HotelID int) (hotelmodel.MiniHotel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMiniHotelByID", HotelID)
	ret0, _ := ret[0].(hotelmodel.MiniHotel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMiniHotelByID indicates an expected call of GetMiniHotelByID
func (mr *MockRepositoryMockRecorder) GetMiniHotelByID(HotelID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMiniHotelByID", reflect.TypeOf((*MockRepository)(nil).GetMiniHotelByID), HotelID)
}
