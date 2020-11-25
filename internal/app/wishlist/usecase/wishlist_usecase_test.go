package wishlistusecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bxcodec/faker/v3"
	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/error"

	mock_wishlist "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/wishlist/mocks"
	wishlistModel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/wishlist/models"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/serverError"
	"github.com/golang/mock/gomock"
)

func TestwishlistUseCase_GetWishlistMeta(t *testing.T) {
	wishlistFakeMeta := make([]wishlistModel.WishlisstHotel, 2)
	err := faker.FakeData(&wishlistFakeMeta)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when create fake data", err)
	}

	t.Run("GetWishlistMeta", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()
		mockWishlistRepository := mock_wishlist.NewMockRepository(controller)
		mockWishlistRepository.EXPECT().GetWishlistMeta(42).Return(wishlistFakeMeta, nil)
		wishlistUsecase := NewWishlistUseCase(mockWishlistRepository)
		wishlistMeta, err := wishlistUsecase.GetWishlistMeta(42)
		assert.NoError(t, err)
		assert.Equal(t, wishlistFakeMeta, wishlistMeta)
	})

	t.Run("GetWishlistMetaError", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()
		mockWishlistRepository := mock_wishlist.NewMockRepository(controller)
		mockWishlistRepository.EXPECT().GetWishlistMeta(42).Return(wishlistFakeMeta, customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))
		wishlistUsecase := NewWishlistUseCase(mockWishlistRepository)
		_, err := wishlistUsecase.GetWishlistMeta(42)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})

}
func TestwishlistUseCase_CreateWishlist(t *testing.T) {}
func TestwishlistUseCase_DeleteWishlist(t *testing.T) {}
func TestwishlistUseCase_AddHotel(t *testing.T)       {}
func TestwishlistUseCase_DeleteHotel(t *testing.T)    {}
