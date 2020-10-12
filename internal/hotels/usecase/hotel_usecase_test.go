package hotelUsecase

import (
	"errors"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/hotels/mocks"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/hotels/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestHotelUseCase_GetUserByID(t *testing.T) {
	mockHotelRepo := new(mocks.HotelsRepository)
	mockHotel := models.Hotel{Name: "bestHotel",Description: "in the world"}
	t.Run("HotelGetByID",func(t *testing.T) {
		mockHotelRepo := new(mocks.HotelsRepository)

		mockHotelRepo.On("GetHotelByID", mock.AnythingOfType("int")).Return(mockHotel, nil).Once()
		u := NewHotelUsecase(mockHotelRepo)

		hotel, err := u.GetHotelByID(mockHotel.HotelID)

		assert.NoError(t, err)
		assert.NotNil(t, hotel)
		mockHotelRepo.AssertExpectations(t)
	})
	t.Run("HotelGetByID-error",func(t *testing.T) {
		mockHotelRepoErr := new(mocks.HotelsRepository)
		mockHotelRepoErr.On("GetHotelByID",mock.AnythingOfType("int")).Return(mockHotel, errors.New("fdsw")).Once()

		uEr := NewHotelUsecase(mockHotelRepoErr)

		_, err := uEr.GetHotelByID(mockHotel.HotelID)
		assert.Error(t,err)
		mockHotelRepo.AssertExpectations(t)
	})
}

func TestHotelUseCase_Add(t *testing.T) {

	mockHotel := [](models.Hotel){{Name: "bestHotel",Description: "in the world"},
		{Name: "bestHotel",Description: "in the world"}}

	t.Run("HotelsGet",func(t *testing.T) {
		mockHotelRepo := new(mocks.HotelsRepository)
		mockHotelRepo.On("GetHotels").Return(mockHotel, nil).Once()
		u := NewHotelUsecase(mockHotelRepo)

		hotels, err := u.GetHotels()

		assert.NoError(t,err)
		assert.NotNil(t,hotels)

		mockHotelRepo.AssertExpectations(t)
	})

	t.Run("HotelsGet-error",func(t *testing.T) {
		mockHotelRepoErr := new(mocks.HotelsRepository)
		mockHotelRepoErr.On("GetHotels").Return(mockHotel, errors.New("fdsw")).Once()

		uEr := NewHotelUsecase(mockHotelRepoErr)

		_, err := uEr.GetHotels()
		assert.Error(t,err)
		mockHotelRepoErr.AssertExpectations(t)
	})
}
