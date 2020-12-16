package recommendRepository

import (
	"errors"
	"testing"

	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/package/error"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/serverError"

	recommModels "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/recommendation/models"

	"github.com/lib/pq"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2020_2_JMickhs/main/configs"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestPostgreRecommendationRepository_GetHotelByIDs(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	hotels := []recommModels.HotelRecommend{
		{HotelID: 0, Name: "hotel", Image: "src/kek.jpg", Location: "Moscow", Rating: "3.5"},
		{HotelID: 1, Name: "hotel", Image: "src/kek.jpg", Location: "Moscow", Rating: "3.5"},
	}
	t.Run("GetHotelByIDs", func(t *testing.T) {
		rowsHotel := sqlmock.NewRows([]string{"hotel_id", "name", "concat", "location", "curr_rating"}).AddRow(
			0, "hotel", "src/kek.jpg", "Moscow", "3.5").AddRow(
			1, "hotel", "src/kek.jpg", "Moscow", "3.5")

		hotelIDs := []int64{0, 1}

		query := GetBestRecommendationsRequest
		mock.ExpectQuery(query).WithArgs(viper.GetString(configs.ConfigFields.S3Url),
			viper.GetInt(configs.ConfigFields.RecommendationCount), pq.Array(hotelIDs)).WillReturnRows(rowsHotel)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgreRecommendationRepository(sqlxDb, nil)

		testHotels, err := rep.GetHotelByIDs(hotelIDs)
		assert.NoError(t, err)
		assert.Equal(t, testHotels, hotels)
	})
	t.Run("GetHotelByIDsErr", func(t *testing.T) {

		hotelIDs := []int64{0, 1}

		query := GetBestRecommendationsRequest
		mock.ExpectQuery(query).
			WithArgs(viper.GetString(configs.ConfigFields.S3Url),
				viper.GetInt(configs.ConfigFields.RecommendationCount), pq.Array(hotelIDs)).
			WillReturnError(customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgreRecommendationRepository(sqlxDb, nil)

		_, err := rep.GetHotelByIDs(hotelIDs)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})
}

func TestPostgreRecommendationRepository_UpdateUserRecommendations(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	userID := 3
	hotelIDs := []int64{1, 2}
	t.Run("UpdateUserRecommendations", func(t *testing.T) {

		query := UpdateRecommendationsForUser
		mock.ExpectExec(query).
			WithArgs(userID, pq.Array(hotelIDs)).WillReturnResult(sqlmock.NewResult(0, 0))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgreRecommendationRepository(sqlxDb, nil)

		err := rep.UpdateUserRecommendations(userID, hotelIDs)
		assert.NoError(t, err)
	})
	t.Run("UpdateUserRecommendationsErr", func(t *testing.T) {

		query := UpdateRecommendationsForUser
		mock.ExpectExec(query).
			WithArgs(userID, pq.Array(hotelIDs)).
			WillReturnError(customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgreRecommendationRepository(sqlxDb, nil)

		err := rep.UpdateUserRecommendations(userID, hotelIDs)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})
}

func TestPostgreRecommendationRepository_GetUsersComments(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	hotelIDs := []int{0, 1}
	userID := 3
	t.Run("GetUserComments", func(t *testing.T) {
		rowsHotel := sqlmock.NewRows([]string{"hotel_id"}).AddRow(
			0).AddRow(
			1)

		query := GetUserCommentsRequest
		mock.ExpectQuery(query).
			WithArgs(userID).
			WillReturnRows(rowsHotel)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgreRecommendationRepository(sqlxDb, nil)

		testHotels, err := rep.GetUsersComments(userID)
		assert.NoError(t, err)
		assert.Equal(t, testHotels, hotelIDs)
	})
	t.Run("GetUserCommentsErr", func(t *testing.T) {
		query := GetUserCommentsRequest
		mock.ExpectQuery(query).
			WithArgs(userID).
			WillReturnError(customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgreRecommendationRepository(sqlxDb, nil)

		_, err := rep.GetUsersComments(userID)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)

	})
}

func TestPostgreRecommendationRepository_GetRecommendationRows(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	MatrixRows := []recommModels.RecommendMatrixRow{
		{UserID: 3, RatingID: 4, HotelID: 0},
		{UserID: 3, RatingID: 4, HotelID: 1},
	}
	hotelIDs := []int{0, 1}
	userID := 3
	t.Run("GetRecommendationRows", func(t *testing.T) {
		rowsHotel := sqlmock.NewRows([]string{"user_id", "hotel1", "rating1"}).AddRow(
			3, 0, 4).AddRow(
			3, 1, 4)

		query := GetRecommendationsMatrixRows
		mock.ExpectQuery(query).
			WithArgs(pq.Array(hotelIDs)).
			WillReturnRows(rowsHotel)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgreRecommendationRepository(sqlxDb, nil)

		testHotels, err := rep.GetRecommendationRows(userID, hotelIDs)
		assert.NoError(t, err)
		assert.Equal(t, MatrixRows, testHotels)
	})
	t.Run("GetRecommendationRowsErr", func(t *testing.T) {

		query := GetRecommendationsMatrixRows
		mock.ExpectQuery(query).
			WithArgs(pq.Array(hotelIDs)).
			WillReturnError(customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgreRecommendationRepository(sqlxDb, nil)

		_, err := rep.GetRecommendationRows(userID, hotelIDs)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)

	})
}
