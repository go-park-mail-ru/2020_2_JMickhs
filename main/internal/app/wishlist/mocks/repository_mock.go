// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_wishlist is a generated GoMock package.
package mock_wishlist

import (
	models "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/wishlist/models"
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

// GetWishlistMeta mocks base method
func (m *MockRepository) GetWishlistMeta(wishlistID int) ([]models.WishlisstHotel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWishlistMeta", wishlistID)
	ret0, _ := ret[0].([]models.WishlisstHotel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWishlistMeta indicates an expected call of GetWishlistMeta
func (mr *MockRepositoryMockRecorder) GetWishlistMeta(wishlistID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWishlistMeta", reflect.TypeOf((*MockRepository)(nil).GetWishlistMeta), wishlistID)
}

// CreateWishlist mocks base method
func (m *MockRepository) CreateWishlist(wishlist models.Wishlist) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWishlist", wishlist)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateWishlist indicates an expected call of CreateWishlist
func (mr *MockRepositoryMockRecorder) CreateWishlist(wishlist interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWishlist", reflect.TypeOf((*MockRepository)(nil).CreateWishlist), wishlist)
}

// DeleteWishlist mocks base method
func (m *MockRepository) DeleteWishlist(wishlistID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWishlist", wishlistID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWishlist indicates an expected call of DeleteWishlist
func (mr *MockRepositoryMockRecorder) DeleteWishlist(wishlistID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWishlist", reflect.TypeOf((*MockRepository)(nil).DeleteWishlist), wishlistID)
}

// AddHotel mocks base method
func (m *MockRepository) AddHotel(hotelID, wishlistID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddHotel", hotelID, wishlistID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddHotel indicates an expected call of AddHotel
func (mr *MockRepositoryMockRecorder) AddHotel(hotelID, wishlistID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddHotel", reflect.TypeOf((*MockRepository)(nil).AddHotel), hotelID, wishlistID)
}

// DeleteHotel mocks base method
func (m *MockRepository) DeleteHotel(hotelID, wishlistID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteHotel", hotelID, wishlistID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteHotel indicates an expected call of DeleteHotel
func (mr *MockRepositoryMockRecorder) DeleteHotel(hotelID, wishlistID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteHotel", reflect.TypeOf((*MockRepository)(nil).DeleteHotel), hotelID, wishlistID)
}
