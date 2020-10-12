package mocks

import (
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/hotels/models"
	"github.com/stretchr/testify/mock"
)

type HotelsUsecase struct {
	mock.Mock
}

func (_m *HotelsUsecase) GetHotels() ([]models.Hotel, error) {
	ret := _m.Called()

	var r0 []models.Hotel
	if rf, ok := ret.Get(0).(func() []models.Hotel); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).([]models.Hotel)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *HotelsUsecase) GetHotelByID(ID int) (models.Hotel, error) {
	ret := _m.Called(ID)

	var r0 models.Hotel
	if rf, ok := ret.Get(0).(func(ID int) models.Hotel); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(models.Hotel)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(ID int) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
