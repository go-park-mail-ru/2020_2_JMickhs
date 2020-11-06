package commentRepository

import (
	"errors"
	"testing"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/serverError"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/clientError"

	"github.com/go-park-mail-ru/2020_2_JMickhs/configs"
	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/error"

	commModel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/comment/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/s"
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
		rowsComments := sqlmock.NewRows([]string{"user_id", "comm_id", "message", "rating", "avatar", "username", "hotel_id", "time"}).AddRow(
			1, 1, "hello", 3, "src/kek.jpg", "kotik", 1, "22-02-2000").AddRow(
			3, 2, "hello", 3, "src/kek.jpg", "kotik", 1, "22-02-2000")

		commentsTest := commModel.FullCommentInfo{3, 2, 1, "hello",
			3, "src/kek.jpg", "kotik", "22-02-2000"}
		query := s.GetCommentsPostgreRequest

		mock.ExpectQuery(query).
			WithArgs("0", configs.BaseItemsPerPage, "1").
			WillReturnRows(rowsComments)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewCommentRepository(sqlxDb)

		comments, err := rep.GetComments(1, 0)
		assert.NoError(t, err)
		assert.Equal(t, commentsTest, comments[1])
	})
	t.Run("GetCommentsErr", func(t *testing.T) {

		query := s.GetCommentsPostgreRequest

		mock.ExpectQuery(query).
			WithArgs("0", configs.BaseItemsPerPage, "1").
			WillReturnError(errors.New("fdsfs"))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewCommentRepository(sqlxDb)

		_, err := rep.GetComments(1, 0)
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

		commentsTest := commModel.Comment{CommID: 3, Time: "22-02-2000", UserID: 1, HotelID: 1, Message: "hello", Rate: 3}
		query := s.AddCommentsPostgreRequest

		mock.ExpectQuery(query).
			WithArgs(commentsTest.UserID, commentsTest.HotelID, commentsTest.Message, commentsTest.Rate).
			WillReturnRows(rowsComments)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewCommentRepository(sqlxDb)

		comment, err := rep.AddComment(commentsTest)
		assert.NoError(t, err)
		assert.Equal(t, commentsTest, comment)
	})
	t.Run("AddCommentsErr", func(t *testing.T) {

		commentsTest := commModel.Comment{CommID: 3, Time: "22-02-2000", UserID: 1, HotelID: 1, Message: "hello", Rate: 3}
		query := s.AddCommentsPostgreRequest

		mock.ExpectQuery(query).
			WithArgs(commentsTest.UserID, commentsTest.HotelID, commentsTest.Message, commentsTest.Rate).
			WillReturnError(errors.New("fdsf"))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewCommentRepository(sqlxDb)

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
		query := s.DeleteCommentsPostgreRequest

		mock.ExpectQuery(query).
			WithArgs("2").
			WillReturnRows()

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewCommentRepository(sqlxDb)

		err := rep.DeleteComment(2)
		assert.NoError(t, err)
	})
	t.Run("DeleteCommentsErr", func(t *testing.T) {
		query := s.DeleteCommentsPostgreRequest

		mock.ExpectQuery(query).
			WithArgs("2").
			WillReturnError(errors.New(""))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewCommentRepository(sqlxDb)

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
		commentsTest := commModel.Comment{CommID: 3, Time: "22-02-2000", UserID: 1, HotelID: 1, Message: "hello", Rate: 3}

		query := s.UpdateCommentsPostgreRequest

		mock.ExpectQuery(query).
			WithArgs(commentsTest.CommID, commentsTest.Message, commentsTest.Rate).
			WillReturnRows(rowsComments)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewCommentRepository(sqlxDb)

		err := rep.UpdateComment(&commentsTest)
		assert.NoError(t, err)
		assert.Equal(t, commentsTest.Time, "22-02-2000")
	})
	t.Run("UpdateCommentsErr", func(t *testing.T) {
		commentsTest := commModel.Comment{CommID: 3, Time: "22-02-2000", UserID: 1, HotelID: 1, Message: "hello", Rate: 3}

		query := s.UpdateCommentsPostgreRequest

		mock.ExpectQuery(query).
			WithArgs(commentsTest.CommID, commentsTest.Message, commentsTest.Rate).
			WillReturnError(errors.New(""))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewCommentRepository(sqlxDb)

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
		query := s.UpdateHotelRatingPostgreRequest

		mock.ExpectQuery(query).
			WithArgs("4.5", "3").
			WillReturnRows()

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewCommentRepository(sqlxDb)

		err := rep.UpdateHotelRating(3, 4.5)
		assert.NoError(t, err)
	})
	t.Run("UpdateHotelRatingErr", func(t *testing.T) {
		query := s.UpdateHotelRatingPostgreRequest

		mock.ExpectQuery(query).
			WithArgs("4.5", "3").
			WillReturnError(errors.New(""))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewCommentRepository(sqlxDb)

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
		query := s.GetCommentsCountPostgreRequest

		mock.ExpectQuery(query).
			WithArgs(5).
			WillReturnRows(rowsComments)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewCommentRepository(sqlxDb)

		count, err := rep.GetCommentsCount(5)
		assert.NoError(t, err)
		assert.Equal(t, count, 34)
	})
	t.Run("GetCommentsCountErr", func(t *testing.T) {
		query := s.GetCommentsCountPostgreRequest

		mock.ExpectQuery(query).
			WithArgs(5).
			WillReturnError(errors.New("fsd"))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewCommentRepository(sqlxDb)

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

		query := s.GetCurrRatingPostgreRequest
		testInfo := commModel.RateInfo{32, 8.5}

		mock.ExpectQuery(query).
			WithArgs(5).
			WillReturnRows(rowsComments)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewCommentRepository(sqlxDb)

		count, err := rep.GetCurrentRating(5)
		assert.NoError(t, err)
		assert.Equal(t, count, testInfo)
	})
	t.Run("GetCurrentRatingErr", func(t *testing.T) {

		query := s.GetCurrRatingPostgreRequest
		mock.ExpectQuery(query).
			WithArgs(5).
			WillReturnError(errors.New(""))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewCommentRepository(sqlxDb)

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

		query := s.GetPrevRatingOnCommentPostgreRequest

		mock.ExpectQuery(query).
			WithArgs("3").
			WillReturnRows(rowsComments)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewCommentRepository(sqlxDb)

		rate, err := rep.CheckUser(&commentsTest)
		assert.NoError(t, err)
		assert.Equal(t, rate, 8)
	})
	t.Run("CheckUserErr", func(t *testing.T) {
		commentsTest := commModel.Comment{CommID: 3, Time: "22-02-2000", UserID: 1, HotelID: 1, Message: "hello", Rate: 3}

		query := s.GetPrevRatingOnCommentPostgreRequest

		mock.ExpectQuery(query).
			WithArgs("3").
			WillReturnError(errors.New(""))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewCommentRepository(sqlxDb)

		_, err := rep.CheckUser(&commentsTest)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
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

		query := s.GetPrevRatingOnCommentPostgreRequest

		mock.ExpectQuery(query).
			WithArgs("3").
			WillReturnRows(rowsComments)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewCommentRepository(sqlxDb)

		_, err := rep.CheckUser(&commentsTest)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), clientError.Locked)
	})
}
