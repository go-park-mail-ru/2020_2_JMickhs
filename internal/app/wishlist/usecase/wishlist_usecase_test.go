package wishlistusecase

import (
	"errors"
	"testing"

	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/error"

	"github.com/bxcodec/faker/v3"
	hotelmodel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/hotels/models"
	wishlist_mock "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/wishlist/mocks"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/serverError"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHotelUseCase_GetHotels(t *testing.T) {

	testHotels := make([]hotelmodel.MiniHotel, 4)
	err := faker.FakeData(&testHotels)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when create fake data", err)
	}
	t.Run("WishlistGetWishlist", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockHotelRepo := wishlist_mock.NewMockRepository(ctrl)

		mockHotelRepo.EXPECT().GetWishlist(42).
			Return(testHotels, nil)

		u := NewWishlistUseCase(mockHotelRepo)

		hotels, err := u.GetWishlist(42)

		assert.NoError(t, err)
		assert.Equal(t, hotels, testHotels)
	})
	t.Run("HotelGetHotelsErr", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockHotelRepo := wishlist_mock.NewMockRepository(ctrl)

		mockHotelRepo.EXPECT().
			GetWishlist(42).
			Return(testHotels, customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		u := NewWishlistUseCase(mockHotelRepo)

		_, err := u.GetWishlist(42)

		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})
}
