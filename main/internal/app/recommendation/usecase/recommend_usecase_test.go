package reccomendUsecase

import (
	"testing"

	recommend_mock "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/recommendation/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRecommendationsUseCase_AddInSearchHistory(t *testing.T) {
	t.Run("HotelGetByID", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRecommendRepo := recommend_mock.NewMockRepository(ctrl)

		mockRecommendRepo.EXPECT().
			AddInSearchHistory(3, "kekw").
			Return(nil)

		u := NewRecommendationsUseCase(mockRecommendRepo)

		err := u.AddInSearchHistory(3, "kekw")

		assert.NoError(t, err)
	})
}

func TestRecommendationsUseCase_GetHotelsRecommendations(t *testing.T) {
	t.Run("HotelGetByID", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRecommendRepo := recommend_mock.NewMockRepository(ctrl)

		mockRecommendRepo.EXPECT().
			AddInSearchHistory(3, "kekw").
			Return(nil)

		u := NewRecommendationsUseCase(mockRecommendRepo)

		err := u.AddInSearchHistory(3, "kekw")

		assert.NoError(t, err)
	})
}
