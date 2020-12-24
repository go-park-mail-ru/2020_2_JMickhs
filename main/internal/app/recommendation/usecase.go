//go:generate mockgen -source usecase.go -destination mocks/recommend_usecase_mock.go -package recommend_mock
package recommendation

import (
	recommModels "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/recommendation/models"
)

type Usecase interface {
	GetHotelsRecommendations(UserID int) ([]recommModels.HotelRecommend, error)
	BuildMatrix(UserID int, rows []recommModels.RecommendMatrixRow) map[float64]map[float64]float64
	GetBestRecommendations(UserID int, matrix map[float64]map[float64]float64) []int64
	AddInSearchHistory(UserID int, pattern string) error
	AddHistoryHotelsToCollaborative(UserID int, hotelIDs []int64) ([]recommModels.HotelRecommend, error)
	DistCosine(vecA map[float64]float64, vecB map[float64]float64) float64
	DotProduct(vecA map[float64]float64, vecB map[float64]float64) float64
}
