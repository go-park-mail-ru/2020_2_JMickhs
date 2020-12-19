package commentRepository

import (
	"errors"
	"testing"

	"github.com/lib/pq"

	"github.com/go-park-mail-ru/2020_2_JMickhs/main/configs"
	"github.com/spf13/viper"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/clientError"
	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/package/error"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/serverError"

	commModel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/comment/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestCommentRepository_GetComments(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	t.Run("GetComments", func(t *testing.T) {
		rowsComments := sqlmock.NewRows([]string{"user_id", "comm_id", "message", "rating", "concat", "username", "hotel_id", "time"}).AddRow(
			1, 1, "hello", 3, "src/kek.jpg", "kotik", 1, "22-02-2000")

		rowsPhotos := sqlmock.NewRows([]string{"photos"}).AddRow(
			"fds")

		commentsTest := commModel.FullCommentInfo{UserID: 1, CommID: 1, HotelID: 1, Message: "hello",
			Rating: 3, Avatar: "src/kek.jpg", Username: "kotik", Time: "22-02-2000", Photos: []string{"fds"}}
		query1 := GetCommentsPostgreRequest
		query2 := CheckPhotosExistPostgreRequest
		mock.ExpectQuery(query1).
			WithArgs("0", 1, "1", 1).
			WillReturnRows(rowsComments)

		mock.ExpectQuery(query2).
			WithArgs("1", 1, viper.GetString(configs.ConfigFields.S3Url)).
			WillReturnRows(rowsPhotos)

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewCommentRepository(sqlxDb, nil)

		comments, err := rep.GetComments("1", 1, "0", 1)
		assert.NoError(t, err)
		assert.Equal(t, commentsTest, comments[0])
	})
	t.Run("GetCommentsErr", func(t *testing.T) {

		rowsComments := sqlmock.NewRows([]string{"user_id", "comm_id", "message", "rating", "concat", "username", "hotel_id", "time"}).AddRow(
			1, 1, "hello", 3, "src/kek.jpg", "kotik", 1, "22-02-2000")

		query1 := GetCommentsPostgreRequest
		query2 := CheckPhotosExistPostgreRequest
		mock.ExpectQuery(query1).
			WithArgs("0", 1, "1", 1).
			WillReturnRows(rowsComments)

		mock.ExpectQuery(query2).
			WithArgs("1", 1, viper.GetString(configs.ConfigFields.S3Url)).
			WillReturnError(errors.New("fds"))

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewCommentRepository(sqlxDb, nil)

		_, err = rep.GetComments("1", 1, "0", 1)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})
	t.Run("GetCommentsErr2", func(t *testing.T) {
		query := GetCommentsPostgreRequest

		mock.ExpectQuery(query).
			WithArgs("0", 1, "1", 1).
			WillReturnError(errors.New("fdsfs"))

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewCommentRepository(sqlxDb, nil)

		_, err = rep.GetComments("1", 1, "0", 1)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), clientError.BadRequest)
	})
}

func TestCommentRepository_AddComment(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	t.Run("AddComments", func(t *testing.T) {
		rowsComments := sqlmock.NewRows([]string{"comm_id", "time"}).AddRow(
			3, "22-02-2000")

		commentsTest := commModel.Comment{CommID: 3, Time: "22-02-2000", UserID: 1, HotelID: 1, Message: "hello", Rate: 3,
			Photos: []string{"kek,kes"}}
		query := AddCommentsPostgreRequest

		mock.ExpectQuery(query).
			WithArgs(commentsTest.UserID, commentsTest.HotelID, commentsTest.Message, commentsTest.Rate, pq.Array(commentsTest.Photos)).
			WillReturnRows(rowsComments)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewCommentRepository(sqlxDb, nil)

		comment, err := rep.AddComment(commentsTest)
		assert.NoError(t, err)
		assert.Equal(t, commentsTest, comment)
	})
	t.Run("AddCommentsErr", func(t *testing.T) {

		commentsTest := commModel.Comment{CommID: 3, Time: "22-02-2000", UserID: 1, HotelID: 1, Message: "hello", Rate: 3,
			Photos: []string{"kek,kes"}}
		query := AddCommentsPostgreRequest

		mock.ExpectQuery(query).
			WithArgs(commentsTest.UserID, commentsTest.HotelID, commentsTest.Message, commentsTest.Rate, pq.Array(commentsTest.Photos)).
			WillReturnError(errors.New("fdsf"))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewCommentRepository(sqlxDb, nil)

		_, err := rep.AddComment(commentsTest)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), clientError.Locked)
	})
}

func TestCommentRepository_DeleteComment(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	t.Run("DeleteComments", func(t *testing.T) {
		query := DeleteCommentsPostgreRequest

		mock.ExpectQuery(query).
			WithArgs("2").
			WillReturnRows()

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewCommentRepository(sqlxDb, nil)

		err := rep.DeleteComment(2)
		assert.NoError(t, err)
	})
	t.Run("DeleteCommentsErr", func(t *testing.T) {
		query := DeleteCommentsPostgreRequest

		mock.ExpectQuery(query).
			WithArgs("2").
			WillReturnError(errors.New(""))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewCommentRepository(sqlxDb, nil)

		err := rep.DeleteComment(2)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})
}

func TestCommentRepository_UpdateComment(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	t.Run("UpdateComments", func(t *testing.T) {
		rowsComments := sqlmock.NewRows([]string{"time"}).AddRow(
			"22-02-2000")
		rowsPhotos := sqlmock.NewRows([]string{"photos"}).AddRow(
			"kek")
		commentsTest := commModel.Comment{CommID: 3, Time: "22-02-2000", UserID: 1, HotelID: 1, Message: "hello", Rate: 3,
			Photos: []string{"kek"}}

		query1 := UpdateCommentsPostgreRequest
		query2 := GetOneCommentPhotosPostgreRequest

		mock.ExpectQuery(query1).
			WithArgs(commentsTest.CommID, commentsTest.Message, commentsTest.Rate, pq.Array(commentsTest.Photos)).
			WillReturnRows(rowsComments)

		mock.ExpectQuery(query2).
			WithArgs(commentsTest.CommID, viper.GetString(configs.ConfigFields.S3Url)).
			WillReturnRows(rowsPhotos)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewCommentRepository(sqlxDb, nil)

		err := rep.UpdateComment(&commentsTest)
		assert.NoError(t, err)
		assert.Equal(t, commentsTest.Time, "22-02-2000")
	})
	t.Run("UpdateCommentsErr", func(t *testing.T) {
		commentsTest := commModel.Comment{CommID: 3, Time: "22-02-2000", UserID: 1, HotelID: 1, Message: "hello", Rate: 3,
			Photos: []string{"kek,kes"}}
		rowsPhotos := sqlmock.NewRows([]string{"photos"}).AddRow(
			"kek")

		query := UpdateCommentsPostgreRequest
		query2 := GetOneCommentPhotosPostgreRequest

		mock.ExpectQuery(query).
			WithArgs(commentsTest.CommID, commentsTest.Message, commentsTest.Rate, commentsTest.Photos).
			WillReturnError(errors.New(""))

		mock.ExpectQuery(query2).
			WithArgs(commentsTest.CommID, viper.GetString(configs.ConfigFields.S3Url)).
			WillReturnRows(rowsPhotos)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewCommentRepository(sqlxDb, nil)

		err := rep.UpdateComment(&commentsTest)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})
	t.Run("UpdateCommentsErr2", func(t *testing.T) {
		commentsTest := commModel.Comment{CommID: 3, Time: "22-02-2000", UserID: 1, HotelID: 1, Message: "hello", Rate: 3,
			Photos: []string{"kek,kes"}}
		rowsComments := sqlmock.NewRows([]string{"time"}).AddRow(
			"22-02-2000")

		query := UpdateCommentsPostgreRequest
		query2 := GetOneCommentPhotosPostgreRequest

		mock.ExpectQuery(query).
			WithArgs(commentsTest.CommID, commentsTest.Message, commentsTest.Rate, commentsTest.Photos).
			WillReturnRows(rowsComments)

		mock.ExpectQuery(query2).
			WithArgs(commentsTest.CommID, viper.GetString(configs.ConfigFields.S3Url)).
			WillReturnError(errors.New(""))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewCommentRepository(sqlxDb, nil)

		err := rep.UpdateComment(&commentsTest)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})
}

func TestCommentRepository_UpdateHotelRating(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	t.Run("UpdateHotelRating", func(t *testing.T) {
		query := UpdateHotelRatingPostgreRequest

		mock.ExpectQuery(query).
			WithArgs("4.5", "3").
			WillReturnRows()

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewCommentRepository(sqlxDb, nil)

		err := rep.UpdateHotelRating(3, 4.5)
		assert.NoError(t, err)
	})
	t.Run("UpdateHotelRatingErr", func(t *testing.T) {
		query := UpdateHotelRatingPostgreRequest

		mock.ExpectQuery(query).
			WithArgs("4.5", "3").
			WillReturnError(errors.New(""))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewCommentRepository(sqlxDb, nil)

		err := rep.UpdateHotelRating(3, 4.5)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), clientError.BadRequest)
	})
}

func TestGetCommentsCount(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	t.Run("GetCommentsCount", func(t *testing.T) {
		rowsComments := sqlmock.NewRows([]string{"comm_count"}).AddRow(
			34)
		query := GetCommentsCountPostgreRequest

		mock.ExpectQuery(query).
			WithArgs(5).
			WillReturnRows(rowsComments)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewCommentRepository(sqlxDb, nil)

		count, err := rep.GetCommentsCount(5)
		assert.NoError(t, err)
		assert.Equal(t, count, 34)
	})
	t.Run("GetCommentsCountErr", func(t *testing.T) {
		query := GetCommentsCountPostgreRequest

		mock.ExpectQuery(query).
			WithArgs(5).
			WillReturnError(errors.New("fsd"))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewCommentRepository(sqlxDb, nil)

		_, err := rep.GetCommentsCount(5)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})
}

func TestGetCurrentRating(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	t.Run("GetCurrentRating", func(t *testing.T) {
		rowsComments := sqlmock.NewRows([]string{"round", "comm_count"}).AddRow(
			8.5, 32)

		query := GetCurrRatingPostgreRequest
		testInfo := commModel.RateInfo{RatesCount: 32, CurrRating: 8.5}

		mock.ExpectQuery(query).
			WithArgs(5).
			WillReturnRows(rowsComments)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewCommentRepository(sqlxDb, nil)

		count, err := rep.GetCurrentRating(5)
		assert.NoError(t, err)
		assert.Equal(t, count, testInfo)
	})
	t.Run("GetCurrentRatingErr", func(t *testing.T) {

		query := GetCurrRatingPostgreRequest
		mock.ExpectQuery(query).
			WithArgs(5).
			WillReturnError(errors.New(""))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewCommentRepository(sqlxDb, nil)

		_, err := rep.GetCurrentRating(5)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})
}

func TestCommentRepository_CheckUser(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	t.Run("CheckUser", func(t *testing.T) {
		rowsComments := sqlmock.NewRows([]string{"rating", "user_id", "hotel_id"}).AddRow(
			8, 1, 1)
		commentsTest := commModel.Comment{CommID: 3, Time: "22-02-2000", UserID: 1, HotelID: 1, Message: "hello", Rate: 3}

		query := GetPrevRatingOnCommentPostgreRequest

		mock.ExpectQuery(query).
			WithArgs("3").
			WillReturnRows(rowsComments)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewCommentRepository(sqlxDb, nil)

		rate, err := rep.CheckUser(&commentsTest)
		assert.NoError(t, err)
		assert.Equal(t, rate, 8)
	})
	t.Run("CheckUserErr", func(t *testing.T) {
		commentsTest := commModel.Comment{CommID: 3, Time: "22-02-2000", UserID: 1, HotelID: 1, Message: "hello", Rate: 3}

		query := GetPrevRatingOnCommentPostgreRequest

		mock.ExpectQuery(query).
			WithArgs("3").
			WillReturnError(errors.New(""))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewCommentRepository(sqlxDb, nil)

		_, err := rep.CheckUser(&commentsTest)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), clientError.NotFound)
	})
}

func TestCommentRepository_CheckUserErr(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	t.Run("CheckUserErr2", func(t *testing.T) {
		rowsComments := sqlmock.NewRows([]string{"rating", "user_id", "hotel_id"}).AddRow(
			8, 1, 1)
		commentsTest := commModel.Comment{CommID: 3, Time: "22-02-2000", UserID: 2, HotelID: 1, Message: "hello", Rate: 3}

		query := GetPrevRatingOnCommentPostgreRequest

		mock.ExpectQuery(query).
			WithArgs("3").
			WillReturnRows(rowsComments)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewCommentRepository(sqlxDb, nil)

		_, err := rep.CheckUser(&commentsTest)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), clientError.Locked)
	})
}
