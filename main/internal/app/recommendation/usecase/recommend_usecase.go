package reccomendUsecase

import (
	"math"
	"sort"
	"time"

	"github.com/go-park-mail-ru/2020_2_JMickhs/main/configs"

	"github.com/spf13/viper"

	"github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/recommendation"
	recommModels "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/recommendation/models"
)

type RecommendationsUseCase struct {
	recommendRepo recommendation.Repository
}

func NewRecommendationsUseCase(r recommendation.Repository) *RecommendationsUseCase {
	return &RecommendationsUseCase{
		recommendRepo: r,
	}
}

func (p *RecommendationsUseCase) AddInSearchHistory(UserID int, pattern string) error {
	return p.recommendRepo.AddInSearchHistory(UserID, pattern)
}

func (p *RecommendationsUseCase) GetHotelsRecommendations(UserID int) ([]recommModels.HotelRecommend, error) {
	var hotels []recommModels.HotelRecommend
	if UserID != -1 {
		recommend, _ := p.recommendRepo.CheckRecommendationExist(UserID)
		if time.Now().Unix()-recommend.Time.Unix() < viper.GetInt64(configs.ConfigFields.UpdateRecommendationTick)*int64(time.Minute.Seconds()) {
			hotels, err := p.AddHistoryHotelsToCollaborative(UserID, recommend.HotelIDs)
			if err != nil {
				return hotels, err
			}
			return hotels, nil
		}
		hotelsIDs, err := p.recommendRepo.GetUsersComments(UserID)
		if err != nil {
			return hotels, err
		}
		if len(hotelsIDs) == 0 {
			hotels, err := p.recommendRepo.GetHotelsFromHistory(UserID, recommend.HotelIDs)
			if err != nil {
				return hotels, err
			}
			if len(hotels) != 0 {

				return hotels, nil
			}
		}
		rows, err := p.recommendRepo.GetRecommendationRows(UserID, hotelsIDs)
		if err != nil {
			return hotels, err
		}
		matrix := p.BuildMatrix(UserID, rows)
		hotelIDs := p.GetBestRecommendations(UserID, matrix)
		hotels, err := p.AddHistoryHotelsToCollaborative(UserID, recommend.HotelIDs)
		if len(hotels) == 0 {
			hotels, err := p.recommendRepo.GetHotelsRecommendations(UserID)
			if err != nil {
				return hotels, err
			}
			return hotels, nil
		}
		if err != nil {
			return hotels, err
		}
		err = p.recommendRepo.UpdateUserRecommendations(UserID, hotelIDs)
		if err != nil {
			return hotels, err
		}
		return hotels, err
	}
	hotels, err := p.recommendRepo.GetHotelsRecommendations(UserID)
	if err != nil {
		return hotels, err
	}
	return hotels, nil
}
func (p *RecommendationsUseCase) AddHistoryHotelsToCollaborative(UserID int, hotelIDs []int64) ([]recommModels.HotelRecommend, error) {
	hotels, err := p.recommendRepo.GetHotelByIDs(hotelIDs)
	if err != nil {
		return hotels, err
	}
	hotelsFromHistory, err := p.recommendRepo.GetHotelsFromHistory(UserID, hotelIDs)
	if err != nil {
		return hotels, err
	}
	if len(hotels) <= 2 {
		for i := 0; i < 2; i++ {
			if len(hotelsFromHistory) <= i {
				break
			}
			hotels = append(hotels, hotelsFromHistory[i])
		}
	} else {
		for i := 2; i < len(hotels); i++ {
			if len(hotelsFromHistory) <= i-2 {
				break
			}
			hotels[i] = hotelsFromHistory[i-2]
		}
	}
	return hotels, nil
}

func (p *RecommendationsUseCase) BuildMatrix(UserID int, rows []recommModels.RecommendMatrixRow) map[float64]map[float64]float64 {
	matr := map[float64]map[float64]float64{}
	for i := 0; i < len(rows); i++ {
		if matr[rows[i].UserID] == nil {
			matr[rows[i].UserID] = make(map[float64]float64)
		}
		matr[rows[i].UserID][rows[i].HotelID] = rows[i].RatingID
	}
	return matr
}

func (p *RecommendationsUseCase) distCosine(vecA map[float64]float64, vecB map[float64]float64) float64 {
	dot := p.dotProduct(vecA, vecB) / (math.Sqrt(p.dotProduct(vecA, vecA)) * math.Sqrt(p.dotProduct(vecB, vecB)))
	return dot
}

func (p *RecommendationsUseCase) dotProduct(vecA map[float64]float64, vecB map[float64]float64) float64 {
	var d float64
	for keyA, p1 := range vecA {
		for keyB, p2 := range vecB {
			if keyA == keyB {
				d += p1 * p2
			}
		}
	}
	return d
}

func (p *RecommendationsUseCase) GetBestRecommendations(UserID int, matrix map[float64]map[float64]float64) []int64 {
	var matches []recommModels.Match
	for key, value := range matrix {
		if int(key) != UserID {
			matches = append(matches, recommModels.Match{UserID: int(key), Coefficient: p.distCosine(matrix[float64(UserID)], value)})
		}
	}

	sort.Slice(matches, func(a int, b int) bool {
		return matches[a].Coefficient > matches[b].Coefficient
	})

	sim := map[float64]float64{}
	var simCoeffAll float64
	bestMatches := map[float64]float64{}

	for i := range matches {
		simCoeffAll += matches[i].Coefficient
	}

	for i := range matches {
		if matches[i].Coefficient > 0.0 {
			bestMatches[float64(matches[i].UserID)] = matches[i].Coefficient
		}
	}

	for relatedUser, coefficient := range bestMatches {
		for hotel := range matrix[relatedUser] {
			check := false
			for key := range matrix[float64(UserID)] {
				if key != hotel {
					check = true
				}
			}
			if !check {
				sim[hotel] = 0.0
				continue
			}
			sim[hotel] += matrix[relatedUser][hotel] * coefficient
		}
	}

	var bestProducts []recommModels.BestProduct
	for key := range sim {
		sim[key] /= simCoeffAll
		bestProducts = append(bestProducts, recommModels.BestProduct{HotelID: int(key), Coefficient: sim[key]})
	}

	sort.Slice(bestProducts, func(a int, b int) bool {
		return bestProducts[a].Coefficient > bestProducts[b].Coefficient
	})

	var hotelIDs []int64
	for i := 0; i < len(bestProducts); i++ {
		if i > viper.GetInt(configs.ConfigFields.RecommendationCount) {
			break
		}
		if bestProducts[i].Coefficient > 0 {
			hotelIDs = append(hotelIDs, int64(bestProducts[i].HotelID))
		}
	}
	return hotelIDs
}
