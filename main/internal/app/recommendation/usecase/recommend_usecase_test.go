package reccomendUsecase

import (
	"testing"

	recommModels "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/recommendation/models"

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

func TestRecommendationsUseCase_GetBestRecommendations(t *testing.T) {
	t.Run("GetBestRecommendations", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRecommendRepo := recommend_mock.NewMockRepository(ctrl)

		u := NewRecommendationsUseCase(mockRecommendRepo)
		matrix := map[float64]map[float64]float64{
			float64(1): {float64(4): float64(5), float64(9): float64(4), float64(10): float64(5)},
			float64(2): {float64(3): float64(4), float64(7): float64(4), float64(1): float64(4)},
			float64(3): {float64(3): float64(4), float64(4): float64(4), float64(5): float64(4)},
		}
		res := u.GetBestRecommendations(3, matrix)
		resTest := []int64{10}
		assert.Equal(t, res, resTest)
	})
	t.Run("GetBestRecommendations", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRecommendRepo := recommend_mock.NewMockRepository(ctrl)

		u := NewRecommendationsUseCase(mockRecommendRepo)
		matrix := map[float64]map[float64]float64{
			float64(1): {float64(4): float64(5), float64(9): float64(4), float64(10): float64(5)},
			float64(2): {float64(3): float64(4), float64(7): float64(4), float64(1): float64(4)},
			float64(3): {float64(3): float64(4), float64(4): float64(4), float64(5): float64(4)},
		}
		res := u.GetBestRecommendations(2, matrix)
		resTest := []int64{4}
		assert.Equal(t, res, resTest)
	})
}

func TestRecommendationsUseCase_DotProduct(t *testing.T) {
	t.Run("DotProduct", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRecommendRepo := recommend_mock.NewMockRepository(ctrl)

		u := NewRecommendationsUseCase(mockRecommendRepo)
		vecA := map[float64]float64{
			float64(1): float64(2), float64(2): float64(3), float64(3): float64(4),
		}
		vecB := map[float64]float64{
			float64(1): float64(1), float64(2): float64(2), float64(3): float64(2),
		}
		res := u.DotProduct(vecA, vecB)
		resTest := 16.0
		assert.Equal(t, res, resTest)
	})
	t.Run("DotProduct", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRecommendRepo := recommend_mock.NewMockRepository(ctrl)

		u := NewRecommendationsUseCase(mockRecommendRepo)
		vecA := map[float64]float64{
			float64(1): float64(5), float64(4): float64(3), float64(3): float64(4),
		}
		vecB := map[float64]float64{
			float64(1): float64(1), float64(0): float64(2), float64(3): float64(2),
		}
		res := u.DotProduct(vecA, vecB)
		resTest := 13.0
		assert.Equal(t, res, resTest)
	})
}

func TestRecommendationsUseCase_DistCosine(t *testing.T) {
	t.Run("DistCosine", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRecommendRepo := recommend_mock.NewMockRepository(ctrl)

		u := NewRecommendationsUseCase(mockRecommendRepo)
		vecA := map[float64]float64{
			float64(1): float64(2), float64(2): float64(3), float64(3): float64(4),
		}
		vecB := map[float64]float64{
			float64(1): float64(1), float64(2): float64(2), float64(3): float64(2),
		}
		res := u.DistCosine(vecA, vecB)
		resTest := 0.9903751369442767
		assert.Equal(t, res, resTest)
	})
	t.Run("DistCosine", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRecommendRepo := recommend_mock.NewMockRepository(ctrl)

		u := NewRecommendationsUseCase(mockRecommendRepo)
		vecA := map[float64]float64{
			float64(1): float64(5), float64(4): float64(3), float64(3): float64(4),
		}
		vecB := map[float64]float64{
			float64(1): float64(1), float64(0): float64(2), float64(3): float64(2),
		}
		res := u.DistCosine(vecA, vecB)
		resTest := 0.6128258770283411
		assert.Equal(t, res, resTest)
	})
}

func TestRecommendationsUseCase_BuildMatrix(t *testing.T) {
	t.Run("BuildMatrix", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRecommendRepo := recommend_mock.NewMockRepository(ctrl)

		u := NewRecommendationsUseCase(mockRecommendRepo)
		matrix := map[float64]map[float64]float64{
			float64(3): {float64(1): float64(2)},
			float64(2): {float64(2): float64(2)},
			float64(1): {float64(3): float64(2)},
			float64(0): {float64(4): float64(2)},
		}
		rows := []recommModels.RecommendMatrixRow{
			{3, 2, 1},
			{2, 2, 2},
			{1, 2, 3},
			{0, 2, 4},
		}
		res := u.BuildMatrix(3, rows)
		assert.Equal(t, res, matrix)
	})
	t.Run("BuildMatrix", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRecommendRepo := recommend_mock.NewMockRepository(ctrl)

		u := NewRecommendationsUseCase(mockRecommendRepo)
		matrix := map[float64]map[float64]float64{
			float64(1): {float64(1): float64(2)},
			float64(2): {float64(2): float64(2)},
			float64(3): {float64(3): float64(2)},
			float64(4): {float64(4): float64(2)},
		}
		rows := []recommModels.RecommendMatrixRow{
			{1, 2, 1},
			{2, 2, 2},
			{3, 2, 3},
			{4, 2, 4},
		}
		res := u.BuildMatrix(3, rows)
		assert.Equal(t, res, matrix)
	})
}

func TestRecommendationsUseCase_AddHistoryHotelsToCollaborative(t *testing.T) {
	t.Run("AddHistory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRecommendRepo := recommend_mock.NewMockRepository(ctrl)

		u := NewRecommendationsUseCase(mockRecommendRepo)
		hotelIDs := []int64{1, 2}
		userID := 2
		var hotels1 []recommModels.HotelRecommend
		hotels2 := []recommModels.HotelRecommend{
			{HotelID: 1, Name: "kek", Image: "kekw.jpeg", Location: "moscow", Rating: "3"},
			{HotelID: 2, Name: "kek", Image: "kekw.jpeg", Location: "moscow", Rating: "3"},
		}
		mockRecommendRepo.EXPECT().
			GetHotelsFromHistory(userID, hotelIDs).
			Return(hotels2, nil)

		mockRecommendRepo.EXPECT().
			GetHotelByIDs(hotelIDs).
			Return(hotels1, nil)

		hotels, err := u.AddHistoryHotelsToCollaborative(userID, hotelIDs)
		assert.NoError(t, err)
		assert.Equal(t, hotels, hotels2)
	})
	t.Run("AddHistory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRecommendRepo := recommend_mock.NewMockRepository(ctrl)

		u := NewRecommendationsUseCase(mockRecommendRepo)
		hotelIDs := []int64{1, 2}
		userID := 2
		var hotels2 []recommModels.HotelRecommend
		hotels1 := []recommModels.HotelRecommend{
			{HotelID: 1, Name: "kek", Image: "kekw.jpeg", Location: "moscow", Rating: "3"},
			{HotelID: 2, Name: "kek", Image: "kekw.jpeg", Location: "moscow", Rating: "3"},
		}
		mockRecommendRepo.EXPECT().
			GetHotelsFromHistory(userID, hotelIDs).
			Return(hotels1, nil)

		mockRecommendRepo.EXPECT().
			GetHotelByIDs(hotelIDs).
			Return(hotels2, nil)

		hotels, err := u.AddHistoryHotelsToCollaborative(userID, hotelIDs)
		assert.NoError(t, err)
		assert.Equal(t, hotels, hotels1)
	})
	t.Run("AddHistory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRecommendRepo := recommend_mock.NewMockRepository(ctrl)

		u := NewRecommendationsUseCase(mockRecommendRepo)
		hotelIDs := []int64{1, 2}
		userID := 2
		hotels2 := []recommModels.HotelRecommend{
			{HotelID: 2, Name: "kek", Image: "kekw.jpeg", Location: "moscow", Rating: "3"},
		}
		hotels1 := []recommModels.HotelRecommend{
			{HotelID: 1, Name: "kek", Image: "kekw.jpeg", Location: "moscow", Rating: "3"},
		}
		hotelsTest := append(hotels2, hotels1...)
		mockRecommendRepo.EXPECT().
			GetHotelsFromHistory(userID, hotelIDs).
			Return(hotels1, nil)

		mockRecommendRepo.EXPECT().
			GetHotelByIDs(hotelIDs).
			Return(hotels2, nil)

		hotels, err := u.AddHistoryHotelsToCollaborative(userID, hotelIDs)
		assert.NoError(t, err)
		assert.Equal(t, hotels, hotelsTest)
	})
}
