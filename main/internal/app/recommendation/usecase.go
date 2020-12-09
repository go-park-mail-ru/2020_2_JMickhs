package recommendation

import (
	recommModels "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/recommendation/models"
)

type Usecase interface {
	GetHotelsRecommendations(UserID int) ([]recommModels.HotelRecommend, error)
	BuildMatrix(UserID int, rows []recommModels.RecommendMatrixRow) map[float64]map[float64]float64
	GetBestRecommendations(UserID int, matrix map[float64]map[float64]float64) []int64
}
