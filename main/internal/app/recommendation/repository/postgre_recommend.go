package recommendRepository

import (
	"context"
	"strconv"

	"github.com/go-park-mail-ru/2020_2_JMickhs/main/configs"
	recommModels "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/recommendation/models"
	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/package/error"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/serverError"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/spf13/viper"
)

type PostgreRecommendationRepository struct {
	conn         *sqlx.DB
	historyStore *redis.Client
}

func NewPostgreRecommendationRepository(conn *sqlx.DB, historyStore *redis.Client) PostgreRecommendationRepository {
	return PostgreRecommendationRepository{conn, historyStore}
}

func (p *PostgreRecommendationRepository) AddInSearchHistory(UserID int, pattern string) error {

	count, err := p.historyStore.Do(context.Background(), "LLEN", strconv.Itoa(UserID)).Int()
	if err != nil {
		return customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	if count > 10 {
		err := p.historyStore.Do(context.Background(), "RPOP", strconv.Itoa(UserID)).Err()
		if err != nil {
			return customerror.NewCustomError(err, serverError.ServerInternalError, 1)
		}
	}
	err = p.historyStore.Do(context.Background(), "LPUSH", strconv.Itoa(UserID), pattern).Err()
	if err != nil {
		return customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return nil
}

func (p *PostgreRecommendationRepository) GetHotelsFromHistory(userID int, hotelIDs []int64) ([]recommModels.HotelRecommend, error) {
	var hotels []recommModels.HotelRecommend

	patternsInterface, err := p.historyStore.Do(context.Background(), "LRANGE", strconv.Itoa(userID), 0, 10).Result()
	if err != nil {
		return hotels, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	listOfInterfaces := patternsInterface.([]interface{})
	var patterns string
	for num, pattern := range listOfInterfaces {
		patterns += pattern.(string)
		if num < len(listOfInterfaces)-1 {
			patterns += "|"
		}
	}
	if len(patterns) == 0 {
		return hotels, nil
	}

	err = p.conn.Select(&hotels, GetRecommendationFromSearchHistory, viper.GetString(configs.ConfigFields.S3Url),
		viper.GetInt(configs.ConfigFields.RecommendationCount), patterns, pq.Array(hotelIDs))
	if err != nil {
		return hotels, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return hotels, nil
}

func (p *PostgreRecommendationRepository) CheckRecommendationExist(userID int) (recommModels.Recommendation, error) {
	recommend := recommModels.Recommendation{}
	err := p.conn.QueryRow(GetRecommendationsForUser, userID).Scan(&recommend.UserID, pq.Array(&recommend.HotelIDs), &recommend.Time)
	if err != nil {
		return recommend, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return recommend, nil
}

func (p *PostgreRecommendationRepository) UpdateUserRecommendations(userID int, hotelIDs []int64) error {
	_, err := p.conn.Exec(UpdateRecommendationsForUser, userID, pq.Array(hotelIDs))
	if err != nil {
		return customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return nil
}

func (p *PostgreRecommendationRepository) GetUsersComments(userID int) ([]int, error) {
	var userIDs []int
	err := p.conn.Select(&userIDs, GetUserCommentsRequest, userID)
	if err != nil {
		return userIDs, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return userIDs, nil
}

func (p *PostgreRecommendationRepository) GetHotelsRecommendations(UserID int) ([]recommModels.HotelRecommend, error) {
	var hotels []recommModels.HotelRecommend
	var err error
	if UserID == -1 {
		err = p.conn.Select(&hotels, GetRecommendationsForNonUnouthoriziedRequest, viper.GetString(configs.ConfigFields.S3Url),
			viper.GetInt(configs.ConfigFields.RecommendationCount))
		if err != nil {
			return hotels, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
		}
		return hotels, nil
	}
	err = p.conn.Select(&hotels, GetRecommendationsForNonUnouthoriziedRequest, viper.GetString(configs.ConfigFields.S3Url),
		viper.GetInt(configs.ConfigFields.RecommendationCount))
	if err != nil {
		return hotels, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return hotels, nil
}

func (p *PostgreRecommendationRepository) GetRecommendationRows(UserID int, hotelIDs []int) ([]recommModels.RecommendMatrixRow, error) {
	var Rows []recommModels.RecommendMatrixRow
	err := p.conn.Select(&Rows, GetRecommendationsMatrixRows, pq.Array(hotelIDs))
	if err != nil {
		return Rows, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return Rows, nil
}

func (p *PostgreRecommendationRepository) GetHotelByIDs(hotelIDs []int64) ([]recommModels.HotelRecommend, error) {
	var hotels []recommModels.HotelRecommend

	err := p.conn.Select(&hotels, GetBestRecommendationsRequest, viper.GetString(configs.ConfigFields.S3Url),
		viper.GetInt(configs.ConfigFields.RecommendationCount), pq.Array(hotelIDs))
	if err != nil {
		return hotels, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return hotels, nil
}
