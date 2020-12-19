package hotelRepository

import (
	"errors"
	"testing"

	"github.com/go-park-mail-ru/2020_2_JMickhs/main/configs"
	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/package/error"
	"github.com/spf13/viper"

	hotelmodel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/hotels/models"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/clientError"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/serverError"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestGetHoteBytIDErr(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	t.Run("GetHotelByIDPhotosErr2", func(t *testing.T) {
		rowsHotel := sqlmock.NewRows([]string{"hotel_id", "name", "description", "img", "location", "curr_rating", "comm_count"}).AddRow(
			1, "hotel", "top hotel in the world", "src/kek.jpg", "Moscow", "3.5", "4")

		query := GetHotelByIDPostgreRequest
		mock.ExpectQuery(query).WithArgs("1", viper.GetString(configs.ConfigFields.S3Url)).WillReturnRows(rowsHotel)

		query = GetHotelsPhotosPostgreRequest
		mock.ExpectQuery(query).WithArgs("1", viper.GetString(configs.ConfigFields.S3Url)).
			WillReturnError(errors.New(""))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresHotelRepository(sqlxDb, nil)

		_, err := rep.GetHotelByID(1)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), clientError.Gone)
	})
	t.Run("GetHotelByIDPhotosErr1", func(t *testing.T) {
		rowsImages := sqlmock.NewRows([]string{"photos"}).AddRow(
			"kek.jpeg")

		query := GetHotelByIDPostgreRequest

		mock.ExpectQuery(query).WithArgs("2", viper.GetString(configs.ConfigFields.S3Url)).
			WillReturnError(errors.New(""))

		query = GetHotelsPhotosPostgreRequest

		mock.ExpectQuery(query).WithArgs("2", viper.GetString(configs.ConfigFields.S3Url)).WillReturnRows(rowsImages)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresHotelRepository(sqlxDb, nil)

		_, err := rep.GetHotelByID(2)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), clientError.Gone)
	})
}

func TestGetHoteBytID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	t.Run("GetHotelByID", func(t *testing.T) {
		rowsHotel := sqlmock.NewRows([]string{"hotel_id", "name", "description", "img", "location", "curr_rating", "comm_count", "latitude", "longitude"}).AddRow(
			3, "hotel", "top hotel in the world", "src/kek.jpg", "Moscow Russia", 3.5, 4, 55.6, 34.5)

		rowsImages := sqlmock.NewRows([]string{"photos"}).AddRow(
			"kek.jpeg")

		hotelTest := hotelmodel.Hotel{HotelID: 3, Name: "hotel", Description: "top hotel in the world",
			Image: "src/kek.jpg", Location: "Moscow Russia",
			Rating: 3.5, Photos: []string{"kek.jpeg"}, CommCount: 4, Latitude: 55.6, Longitude: 34.5}

		query := GetHotelByIDPostgreRequest

		mock.ExpectQuery(query).WithArgs("3", viper.GetString(configs.ConfigFields.S3Url)).WillReturnRows(rowsHotel)

		query = GetHotelsPhotosPostgreRequest

		mock.ExpectQuery(query).WithArgs("3", viper.GetString(configs.ConfigFields.S3Url)).WillReturnRows(rowsImages)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresHotelRepository(sqlxDb, nil)

		hotel, err := rep.GetHotelByID(3)
		assert.NoError(t, err)
		assert.Equal(t, hotel, hotelTest)
	})
}

func TestFetchHotels(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	t.Run("FetchHotels", func(t *testing.T) {
		rowsHotel := sqlmock.NewRows([]string{"hotel_id", "name", "concat", "location", "curr_rating", "comm_count"}).AddRow(
			1, "Villa", "src/kek.jpg", "Moscow Russia", "3.5", "4").AddRow(
			2, "Hostel", "src/kek.jpg", "China", "7", "3")

		hotelTest := hotelmodel.Hotel{HotelID: 1, Name: "Villa",
			Image: "src/kek.jpg", Location: "Moscow Russia", Rating: 3.5, CommCount: 4}

		baseQuery := "SELECT hotel_id, name, description, location, concat($4::varchar,img),country,city,curr_rating , " +
			"comm_count,strict_word_similarity($1,name) as t1,strict_word_similarity($1,location) as t2 "

		baseQuery += " FROM hotels " + SearchHotelsPostgreRequest
		baseQuery += " AND (curr_rating BETWEEN $5 AND $6 OR curr_rating BETWEEN $6 AND $5) "
		baseQuery += " AND comm_count >= $7"
		baseQuery += " ORDER BY curr_rating DESC,t1 DESC,t2 DESC "
		baseQuery += "LIMIT $3 OFFSET $2"
		queryLimit := "Select set_limit(0.18)"
		mock.ExpectExec(queryLimit).WillReturnResult(sqlmock.NewResult(0, 0))
		filter := hotelmodel.HotelFiltering{RatingFilterStartNumber: "0", RatingFilterEndNumber: "3", CommentsFilterStartNumber: "0"}
		mock.ExpectQuery(baseQuery).WithArgs("top", 0, viper.GetInt(configs.ConfigFields.BaseItemPerPage),
			viper.GetString(configs.ConfigFields.S3Url), filter.RatingFilterStartNumber, filter.RatingFilterEndNumber,
			filter.CommentsFilterStartNumber).WillReturnRows(rowsHotel)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresHotelRepository(sqlxDb, nil)

		hotels, err := rep.FetchHotels(filter, "top", 0)
		assert.NoError(t, err)
		assert.Equal(t, hotels[0], hotelTest)
	})
}

func TestFetchHotelsErr(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	t.Run("FetchHotelsErr", func(t *testing.T) {
		query := "SELECT hotel_id, name, description, location, concat($4::varchar,img), curr_rating , comm_count FROM hotels " +
			SearchHotelsPostgreRequest + " ORDER BY curr_rating DESC LIMIT $3 OFFSET $2"

		filter := hotelmodel.HotelFiltering{}
		queryLimit := "Select set_limit(0.18)"
		mock.ExpectExec(queryLimit).WillReturnResult(sqlmock.NewResult(0, 0))
		mock.ExpectQuery(query).WithArgs("top", 0, viper.GetString(configs.ConfigFields.BaseItemPerPage),
			viper.GetString(configs.ConfigFields.S3Url)).
			WillReturnError(customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresHotelRepository(sqlxDb, nil)

		_, err := rep.FetchHotels(filter, "top", 0)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})
}
func TestCheckRateExist(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	t.Run("CheckRateExist", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"message", "time", "hotel_id", "user_id",
			"comm_id", "rating"}).AddRow("kekw", "22-02-2000", "3", "1",
			"10", "5")

		rowsPhotos := sqlmock.NewRows([]string{"photos"}).AddRow("fds")

		ratingTest := 5

		query1 := CheckRateIfExistPostgreRequest
		query2 := CheckPhotosExistPostgreRequest

		mock.ExpectQuery(query1).
			WithArgs(5, 3).
			WillReturnRows(rows)

		mock.ExpectQuery(query2).
			WithArgs(5, 3, viper.GetString(configs.ConfigFields.S3Url)).
			WillReturnRows(rowsPhotos)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresHotelRepository(sqlxDb, nil)

		rating, err := rep.CheckRateExist(3, 5)
		assert.NoError(t, err)
		assert.Equal(t, rating.Rating, float64(ratingTest))
	})
	t.Run("CheckRateExistErr1", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"message", "time", "hotel_id", "user_id",
			"comm_id", "rating"}).AddRow("kekw", "22-02-2000", "3", "1",
			"10", "5")

		query1 := CheckRateIfExistPostgreRequest
		query2 := CheckPhotosExistPostgreRequest

		mock.ExpectQuery(query1).
			WithArgs(5, 3).
			WillReturnRows(rows)

		mock.ExpectQuery(query2).
			WithArgs(5, 3, viper.GetString(configs.ConfigFields.S3Url)).
			WillReturnError(customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresHotelRepository(sqlxDb, nil)

		_, err := rep.CheckRateExist(3, 5)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})
	t.Run("CheckRateExistErr2", func(t *testing.T) {
		query := CheckRateIfExistPostgreRequest

		mock.ExpectQuery(query).
			WithArgs(3, 5).
			WillReturnError(customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresHotelRepository(sqlxDb, nil)

		_, err := rep.CheckRateExist(3, 5)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})
}

func TestGetHotelsPreview(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	t.Run("GetHotelsPreview", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"hotel_id", "name", "concat", "location"}).AddRow(
			1, "Villa", "src/kek.jpg", "Moscow").AddRow(
			2, "Hostel", "src/kek.jpg", "China")

		query := "SELECT hotel_id, name, location, concat($3::varchar,img) FROM hotels " +
			SearchHotelsPostgreRequest + " ORDER BY curr_rating DESC LIMIT $2"

		hotelTest := hotelmodel.HotelPreview{HotelID: 1, Name: "Villa", Image: "src/kek.jpg", Location: "Moscow"}

		mock.ExpectQuery(query).WithArgs("top", viper.GetInt(configs.ConfigFields.PreviewItemLimit),
			viper.GetString(configs.ConfigFields.S3Url)).WillReturnRows(rows)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresHotelRepository(sqlxDb, nil)

		hotels, err := rep.GetHotelsPreview("top")
		assert.NoError(t, err)
		assert.Equal(t, hotels[0], hotelTest)
	})
	t.Run("GetHotelsPreviewErr", func(t *testing.T) {
		query := "SELECT hotel_id, name, location, concat($3::varchar,img) FROM hotels " +
			SearchHotelsPostgreRequest + " ORDER BY curr_rating DESC LIMIT $2"

		mock.ExpectQuery(query).WithArgs("top", viper.GetString(configs.ConfigFields.PreviewItemLimit), viper.GetString(configs.ConfigFields.S3Url)).
			WillReturnError(customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresHotelRepository(sqlxDb, nil)

		_, err := rep.GetHotelsPreview("top")
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})
}
