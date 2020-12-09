package recommendRepository

import (
	"github.com/go-park-mail-ru/2020_2_JMickhs/main/configs"
	recommModels "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/recommendation/models"
	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/package/error"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/serverError"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/spf13/viper"
)

type PostgreRecommendationRepository struct {
	conn *sqlx.DB
}

func NewPostgreRecommendationRepository(conn *sqlx.DB) PostgreRecommendationRepository {
	return PostgreRecommendationRepository{conn}
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

func (p *PostgreRecommendationRepository) GetUsersComments(hotelID int) ([]int, error) {
	var userIDs []int
	err := p.conn.Select(&userIDs, GetUserCommentsRequest, hotelID)
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
