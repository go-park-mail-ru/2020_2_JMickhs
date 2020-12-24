//go:generate mockgen -source repository.go -destination mocks/recommend_recommend_mock.go -package recommend_mock
package recommendation

import (
	recommModels "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/recommendation/models"
)

type Repository interface {
	GetHotelsRecommendations(UserID int) ([]recommModels.HotelRecommend, error)
	GetRecommendationRows(UserID int, hotelIDs []int) ([]recommModels.RecommendMatrixRow, error)
	GetHotelByIDs(hotelIDs []int64) ([]recommModels.HotelRecommend, error)
	UpdateUserRecommendations(userID int, hotelIDs []int64) error
	GetUsersComments(userID int) ([]int, error)
	CheckRecommendationExist(userID int) (recommModels.Recommendation, error)
	AddInSearchHistory(UserID int, pattern string) error
	GetHotelsFromHistory(UserID int, hotelIDs []int64) ([]recommModels.HotelRecommend, error)
}
