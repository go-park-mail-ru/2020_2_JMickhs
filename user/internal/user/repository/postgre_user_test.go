package userRepository

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/package/error"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/clientError"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/serverError"
	"github.com/go-park-mail-ru/2020_2_JMickhs/user/internal/user/models"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestPostgresUserRepository_GetUserByID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	t.Run("GetUserByID", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "username", "email", "password", "avatar"}).AddRow(
			1, "kotik", "kek@mail.ru", "12345", "src/kek.jpg")

		testUser := models.User{1, "kotik", "kek@mail.ru", "12345", "src/kek.jpg"}

		query := GetUserByIDPostgreRequest

		mock.ExpectQuery(query).
			WithArgs("1", "").
			WillReturnRows(rows)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresUserRepository(sqlxDb, nil)

		user, err := rep.GetUserByID(1)
		assert.NoError(t, err)
		assert.Equal(t, testUser, user)
	})

	t.Run("GetUserByIDerr", func(t *testing.T) {
		query := GetUserByIDPostgreRequest

		mock.ExpectQuery(query).
			WithArgs("1", "").
			WillReturnError(customerror.NewCustomError(errors.New(""), clientError.Gone, 1))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresUserRepository(sqlxDb, nil)

		_, err := rep.GetUserByID(1)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), clientError.Gone)
	})
}

func TestPostgresUserRepository_GetByUserName(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	t.Run("GetUserByName", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "username", "email", "password", "avatar"}).AddRow(
			1, "kotik", "kek@mail.ru", "12345", "src/kek.jpg")

		testUser := models.User{1, "kotik", "kek@mail.ru", "12345", "src/kek.jpg"}

		query := GetUserByNamePostgreRequest

		mock.ExpectQuery(query).
			WithArgs("kotik", "").
			WillReturnRows(rows)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresUserRepository(sqlxDb, nil)

		user, err := rep.GetByUserName("kotik")
		assert.NoError(t, err)
		assert.Equal(t, testUser, user)
	})

	t.Run("GetUserByNameErr", func(t *testing.T) {
		query := GetUserByNamePostgreRequest

		mock.ExpectQuery(query).
			WithArgs("kotik", "").
			WillReturnError(customerror.NewCustomError(errors.New(""), clientError.Unauthorizied, 1))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresUserRepository(sqlxDb, nil)

		_, err := rep.GetByUserName("kotik")
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), clientError.Unauthorizied)
	})
}

func TestPostgresUserRepository_UpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	t.Run("UpdateUser", func(t *testing.T) {
		query := UpdateUserPostgreRequest

		mock.ExpectQuery(query).
			WithArgs(1, "kotik", "kek@mail.ru").
			WillReturnRows()

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresUserRepository(sqlxDb, nil)
		user := models.User{ID: 1, Username: "kotik", Email: "kek@mail.ru"}

		err = rep.UpdateUser(user)
		assert.NoError(t, err)
	})
	t.Run("UpdateUserErr", func(t *testing.T) {
		query := UpdateUserPostgreRequest

		mock.ExpectQuery(query).
			WithArgs(1, "kotik", "kek@mail.ru").
			WillReturnError(customerror.NewCustomError(errors.New(""), clientError.Conflict, 1))
		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresUserRepository(sqlxDb, nil)
		user := models.User{ID: 1, Username: "kotik", Email: "kek@mail.ru"}

		err = rep.UpdateUser(user)
		assert.Equal(t, customerror.ParseCode(err), clientError.Conflict)
	})
}

func TestPostgresUserRepository_UpdatePassword(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	t.Run("UpdatePassword", func(t *testing.T) {
		query := UpdateUserPasswordPostgreRequest

		mock.ExpectQuery(query).
			WithArgs(1, "12345").
			WillReturnRows()

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresUserRepository(sqlxDb, nil)
		user := models.User{ID: 1, Password: "12345"}
		err = rep.UpdatePassword(user)
		assert.NoError(t, err)
	})
	t.Run("UpdatePasswordErr", func(t *testing.T) {
		query := UpdateUserPasswordPostgreRequest

		mock.ExpectQuery(query).
			WithArgs(1, "12345").
			WillReturnError(customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresUserRepository(sqlxDb, nil)
		user := models.User{ID: 1, Password: "12345"}
		err = rep.UpdatePassword(user)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})
}

func TestPostgresUserRepository_UpdateAvatar(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	t.Run("UpdateAvatar", func(t *testing.T) {
		query := UpdateUserAvatarPostgreRequest

		mock.ExpectQuery(query).
			WithArgs(1, "src/kek.jpg").
			WillReturnRows()

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresUserRepository(sqlxDb, nil)
		user := models.User{ID: 1, Avatar: "src/kek.jpg"}
		err = rep.UpdateAvatar(user)
		assert.NoError(t, err)
		assert.NotNil(t, user)
	})
	t.Run("UpdateAvatarErr", func(t *testing.T) {
		query := UpdateUserAvatarPostgreRequest

		mock.ExpectQuery(query).
			WithArgs(1, "src/kek.jpg").
			WillReturnError(customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))
		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresUserRepository(sqlxDb, nil)
		user := models.User{ID: 1, Avatar: "src/kek.jpg"}
		err = rep.UpdateAvatar(user)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})

}

func TestPostgresUserRepository_Add(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	t.Run("AddUser", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"user_id"}).AddRow(1)
		query := AddUserPostgreRequest

		mock.ExpectQuery(query).
			WithArgs("kotik", "kek@mail.ru", "12345", "src/kek.jpg").
			WillReturnRows(rows)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresUserRepository(sqlxDb, nil)
		user := models.User{Username: "kotik", Email: "kek@mail.ru", Password: "12345", Avatar: "src/kek.jpg"}

		user, err = rep.Add(user)
		assert.NoError(t, err)
		assert.Equal(t, 1, user.ID)
	})
	t.Run("AddUserErr", func(t *testing.T) {
		query := AddUserPostgreRequest

		mock.ExpectQuery(query).
			WithArgs("kotik", "kek@mail.ru", "12345", "src/kek.jpg").
			WillReturnError(customerror.NewCustomError(errors.New(""), clientError.Conflict, 1))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgresUserRepository(sqlxDb, nil)
		user := models.User{Username: "kotik", Email: "kek@mail.ru", Password: "12345", Avatar: "src/kek.jpg"}

		_, err = rep.Add(user)
		assert.Equal(t, customerror.ParseCode(err), clientError.Conflict)
	})
}

func TestPostgresUserRepository_CompareHashAndPassword(t *testing.T) {
	db, _, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	t.Run("CompareHashAndPassword", func(t *testing.T) {
		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		destinationPassword := "scxzsdfsdxcfewdsfcx"
		inputPassword := "kek123"

		rep := NewPostgresUserRepository(sqlxDb, nil)
		err := rep.CompareHashAndPassword(destinationPassword, inputPassword)
		assert.Error(t, err)
	})
	t.Run("CompareHashAndPasswordNoError", func(t *testing.T) {
		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		inputPassword := "keks123456"

		rep := NewPostgresUserRepository(sqlxDb, nil)
		hash, _ := rep.GenerateHashFromPassword(inputPassword)
		err := rep.CompareHashAndPassword(string(hash), inputPassword)
		assert.NoError(t, err)
	})
}
