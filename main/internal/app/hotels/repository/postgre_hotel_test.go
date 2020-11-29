package hotelRepository

import (
	"errors"
	"fmt"
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

		rep := NewPostgresHotelRepository(sqlxDb, nil, nil)

		_, err := rep.GetHotelByID(1)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
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

		rep := NewPostgresHotelRepository(sqlxDb, nil, nil)

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
		rowsHotel := sqlmock.NewRows([]string{"hotel_id", "name", "description", "img", "location", "curr_rating", "comm_count"}).AddRow(
			3, "hotel", "top hotel in the world", "src/kek.jpg", "Moscow", "3.5", "4")

		rowsImages := sqlmock.NewRows([]string{"photos"}).AddRow(
			"kek.jpeg")

		hotelTest := hotelmodel.Hotel{3, "hotel", "top hotel in the world",
			"src/kek.jpg", "Moscow Russia", "kek@mail.ru", "Russia", "Moscow",
			4, []string{"kek.jpeg"}, 4, 55.6, 34.5, ""}

		query := GetHotelByIDPostgreRequest

		mock.ExpectQuery(query).WithArgs("3", viper.GetString(configs.ConfigFields.S3Url)).WillReturnRows(rowsHotel)

		query = GetHotelsPhotosPostgreRequest

		mock.ExpectQuery(query).WithArgs("3", viper.GetString(configs.ConfigFields.S3Url)).WillReturnRows(rowsImages)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresHotelRepository(sqlxDb, nil, nil)

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
		rowsHotel := sqlmock.NewRows([]string{"hotel_id", "name", "description", "concat", "location", "curr_rating", "comm_count"}).AddRow(
			1, "Villa", "top hotel in the world", "src/kek.jpg", "Moscow", "3.5", "4").AddRow(
			2, "Hostel", "top hotel in the world", "src/kek.jpg", "China", "7", "3")

		hotelTest := hotelmodel.Hotel{3, "hotel", "top hotel in the world",
			"src/kek.jpg", "Moscow Russia", "kek@mail.ru", "Russia", "Moscow",
			4, []string{"kek.jpeg"}, 4, 55.6, 34.5, ""}

		query := fmt.Sprint("SELECT hotel_id, name, description, location, concat($4::varchar,img), curr_rating , comm_count FROM hotels ",
			SearchHotelsPostgreRequest, " ORDER BY curr_rating DESC LIMIT $3 OFFSET $2")

		filter := hotelmodel.HotelFiltering{}
		mock.ExpectQuery(query).WithArgs("top", 0, viper.GetString(configs.ConfigFields.BaseItemPerPage),
			viper.GetString(configs.ConfigFields.S3Url)).WillReturnRows(rowsHotel)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresHotelRepository(sqlxDb, nil, nil)

		hotels, err := rep.FetchHotels(filter, "top", 0)
		assert.NoError(t, err)
		assert.Equal(t, hotels[0], hotelTest)
	})
	t.Run("FetchHotelsErr", func(t *testing.T) {
		query := fmt.Sprint("SELECT hotel_id, name, description, location, concat($4::varchar,img), curr_rating , comm_count FROM hotels ",
			SearchHotelsPostgreRequest, " ORDER BY curr_rating DESC LIMIT $3 OFFSET $2")

		filter := hotelmodel.HotelFiltering{}
		mock.ExpectQuery(query).WithArgs("top", 0, viper.GetString(configs.ConfigFields.BaseItemPerPage),
			viper.GetString(configs.ConfigFields.S3Url)).
			WillReturnError(customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresHotelRepository(sqlxDb, nil, nil)

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
		rows := sqlmock.NewRows([]string{"message", "time", "hotel_id", "avatar", "user_id",
			"comm_id", "username", "rating"}).AddRow("kekw", "22-02-2000", "3", "src/kek.jpg", "1",
			"10", "kostik", "5")

		ratingTest := 5

		query := CheckRateIfExistPostgreRequest

		mock.ExpectQuery(query).WithArgs(3, 5).WillReturnRows(rows)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresHotelRepository(sqlxDb, nil, nil)

		rating, err := rep.CheckRateExist(3, 5)
		assert.NoError(t, err)
		assert.Equal(t, rating.Rating, float64(ratingTest))
	})
	t.Run("CheckRateExistErr", func(t *testing.T) {
		query := CheckRateIfExistPostgreRequest

		mock.ExpectQuery(query).WithArgs(3, 5).
			WillReturnError(customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresHotelRepository(sqlxDb, nil, nil)

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

		query := fmt.Sprint("SELECT hotel_id, name, location, concat($3::varchar,img) FROM hotels ",
			SearchHotelsPostgreRequest, " ORDER BY curr_rating DESC LIMIT $2")

		hotelTest := hotelmodel.HotelPreview{1, "Villa", "src/kek.jpg", "Moscow"}

		mock.ExpectQuery(query).WithArgs("top", viper.GetString(configs.ConfigFields.PreviewItemLimit),
			viper.GetString(configs.ConfigFields.S3Url)).WillReturnRows(rows)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresHotelRepository(sqlxDb, nil, nil)

		hotels, err := rep.GetHotelsPreview("top")
		assert.NoError(t, err)
		assert.Equal(t, hotels[0], hotelTest)
	})
	t.Run("GetHotelsPreviewErr", func(t *testing.T) {
		query := fmt.Sprint("SELECT hotel_id, name, location, concat($3::varchar,img) FROM hotels ",
			SearchHotelsPostgreRequest, " ORDER BY curr_rating DESC LIMIT $2")

		mock.ExpectQuery(query).WithArgs("top", viper.GetString(configs.ConfigFields.PreviewItemLimit), viper.GetString(configs.ConfigFields.S3Url)).
			WillReturnError(customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresHotelRepository(sqlxDb, nil, nil)

		_, err := rep.GetHotelsPreview("top")
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})
}

func TestGetHotels(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	t.Run("GetHotels", func(t *testing.T) {
		rowsHotel := sqlmock.NewRows([]string{"hotel_id", "name", "description", "concat", "location", "curr_rating", "comm_count"}).AddRow(
			1, "Villa", "top hotel in the world", "src/kek.jpg", "Moscow", "3.5", "4").AddRow(
			2, "Hostel", "top hotel in the world", "src/kek.jpg", "China", "7", "3")

		query := GetHotelsPostgreRequest

		hotelTest := hotelmodel.Hotel{3, "hotel", "top hotel in the world",
			"src/kek.jpg", "Moscow Russia", "kek@mail.ru", "Russia", "Moscow",
			4, []string{"kek.jpeg"}, 4, 55.6, 34.5, ""}

		mock.ExpectQuery(query).WithArgs("4", viper.GetString(configs.ConfigFields.S3Url)).WillReturnRows(rowsHotel)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresHotelRepository(sqlxDb, nil, nil)

		hotels, err := rep.GetHotels(4)
		assert.NoError(t, err)
		assert.Equal(t, hotels[0], hotelTest)
	})
	t.Run("GetHotelsErr", func(t *testing.T) {

		query := GetHotelsPostgreRequest

		mock.ExpectQuery(query).WithArgs("4", viper.GetString(configs.ConfigFields.S3Url)).
			WillReturnError(customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresHotelRepository(sqlxDb, nil, nil)

		_, err := rep.GetHotels(4)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})
}
