package hotelRepository

import (
	"errors"
	"fmt"
	"testing"

	"github.com/go-park-mail-ru/2020_2_JMickhs/configs"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/serverError"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/clientError"
	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/error"

	hotelmodel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/hotels/models"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/sqlrequests"

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

		query := sqlrequests.GetHotelByIDPostgreRequest
		mock.ExpectQuery(query).WithArgs("1", configs.S3Url).WillReturnRows(rowsHotel)

		query = sqlrequests.GetHotelsPhotosPostgreRequest
		mock.ExpectQuery(query).WithArgs("1", configs.S3Url).
			WillReturnError(errors.New(""))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresHotelRepository(sqlxDb)

		_, err := rep.GetHotelByID(1)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})
	t.Run("GetHotelByIDPhotosErr1", func(t *testing.T) {
		rowsImages := sqlmock.NewRows([]string{"photos"}).AddRow(
			"kek.jpeg")

		query := sqlrequests.GetHotelByIDPostgreRequest

		mock.ExpectQuery(query).WithArgs("2", configs.S3Url).
			WillReturnError(errors.New(""))

		query = sqlrequests.GetHotelsPhotosPostgreRequest

		mock.ExpectQuery(query).WithArgs("2", configs.S3Url).WillReturnRows(rowsImages)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresHotelRepository(sqlxDb)

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
			"src/kek.jpg", "Moscow", 3.5, []string{"kek.jpeg"}, 4}

		query := sqlrequests.GetHotelByIDPostgreRequest

		mock.ExpectQuery(query).WithArgs("3", configs.S3Url).WillReturnRows(rowsHotel)

		query = sqlrequests.GetHotelsPhotosPostgreRequest

		mock.ExpectQuery(query).WithArgs("3", configs.S3Url).WillReturnRows(rowsImages)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresHotelRepository(sqlxDb)

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

		hotelTest := hotelmodel.Hotel{1, "Villa", "top hotel in the world",
			"src/kek.jpg", "Moscow", 3.5, nil, 4}

		query := fmt.Sprint("SELECT hotel_id, name, description, location, concat($4::varchar,img), curr_rating , comm_count FROM hotels ",
			sqlrequests.SearchHotelsPostgreRequest, " ORDER BY curr_rating DESC LIMIT $3 OFFSET $2")

		mock.ExpectQuery(query).WithArgs("top", 0, configs.BaseItemsPerPage, configs.S3Url).WillReturnRows(rowsHotel)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresHotelRepository(sqlxDb)

		hotels, err := rep.FetchHotels("top", 0)
		assert.NoError(t, err)
		assert.Equal(t, hotels[0], hotelTest)
	})
	t.Run("FetchHotelsErr", func(t *testing.T) {
		query := fmt.Sprint("SELECT hotel_id, name, description, location, concat($4::varchar,img), curr_rating , comm_count FROM hotels ",
			sqlrequests.SearchHotelsPostgreRequest, " ORDER BY curr_rating DESC LIMIT $3 OFFSET $2")

		mock.ExpectQuery(query).WithArgs("top", 0, configs.BaseItemsPerPage, configs.S3Url).
			WillReturnError(customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresHotelRepository(sqlxDb)

		_, err := rep.FetchHotels("top", 0)
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
		rows := sqlmock.NewRows([]string{"rating"}).AddRow(
			"5")

		ratingTest := 5

		query := sqlrequests.CheckRateIfExistPostgreRequest

		mock.ExpectQuery(query).WithArgs(3, 5).WillReturnRows(rows)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresHotelRepository(sqlxDb)

		rating, err := rep.CheckRateExist(3, 5)
		assert.NoError(t, err)
		assert.Equal(t, rating, ratingTest)
	})
	t.Run("CheckRateExistErr", func(t *testing.T) {
		query := sqlrequests.CheckRateIfExistPostgreRequest

		mock.ExpectQuery(query).WithArgs(3, 5).
			WillReturnError(customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresHotelRepository(sqlxDb)

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

		query := fmt.Sprint("SELECT hotel_id, name, location, concat($4::varchar,img) FROM hotels ",
			sqlrequests.SearchHotelsPostgreRequest, " ORDER BY curr_rating DESC LIMIT $2")

		hotelTest := hotelmodel.HotelPreview{1, "Villa", "src/kek.jpg", "Moscow"}

		mock.ExpectQuery(query).WithArgs("top", configs.PreviewItemLimit, configs.S3Url).WillReturnRows(rows)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresHotelRepository(sqlxDb)

		hotels, err := rep.GetHotelsPreview("top")
		assert.NoError(t, err)
		assert.Equal(t, hotels[0], hotelTest)
	})
	t.Run("GetHotelsPreviewErr", func(t *testing.T) {
		query := fmt.Sprint("SELECT hotel_id, name, location, concat($4::varchar,img)  FROM hotels ",
			sqlrequests.SearchHotelsPostgreRequest, " ORDER BY curr_rating DESC LIMIT $2")

		mock.ExpectQuery(query).WithArgs("top", configs.PreviewItemLimit, configs.S3Url).
			WillReturnError(customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresHotelRepository(sqlxDb)

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

		query := sqlrequests.GetHotelsPostgreRequest

		hotelTest := hotelmodel.Hotel{1, "Villa", "top hotel in the world", "src/kek.jpg", "Moscow", 3.5,
			nil, 4}

		mock.ExpectQuery(query).WithArgs("4", configs.S3Url).WillReturnRows(rowsHotel)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresHotelRepository(sqlxDb)

		hotels, err := rep.GetHotels(4)
		assert.NoError(t, err)
		assert.Equal(t, hotels[0], hotelTest)
	})
	t.Run("GetHotelsErr", func(t *testing.T) {

		query := sqlrequests.GetHotelsPostgreRequest

		mock.ExpectQuery(query).WithArgs("4", configs.S3Url).
			WillReturnError(customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresHotelRepository(sqlxDb)

		_, err := rep.GetHotels(4)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})
}
