package userRepository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/user/models"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUserByID(t *testing.T) {
	db,mock,err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id","username","email","password","avatar"}).AddRow(
		1,"kotik","kek@mail.ru","12345","src/kek.jpg")

	query := "SELECT id,username,email,password,avatar FROM users WHERE id=$1"

	mock.ExpectQuery(query).WithArgs("1").WillReturnRows(rows)
	sqlxDb:= sqlx.NewDb(db,"sqlmock")
	defer sqlxDb.Close()

	rep := NewPostgresUserRepository(sqlxDb)

	user,err := rep.GetUserByID(1)
	assert.NoError(t,err)
	assert.NotNil(t,user)
}


func TestGetUserByName(t *testing.T) {
	db,mock,err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id","username","email","password","avatar"}).AddRow(
		1,"kotik","kek@mail.ru","12345","src/kek.jpg")

	query := "select id,username,email,password,avatar FROM users WHERE username=$1"

	mock.ExpectQuery(query).WithArgs("kotik").WillReturnRows(rows)
	sqlxDb:= sqlx.NewDb(db,"sqlmock")
	defer sqlxDb.Close()

	rep := NewPostgresUserRepository(sqlxDb)

	user,err := rep.GetByUserName("kotik")
	assert.NoError(t,err)
	assert.NotNil(t,user)
}

func TestUpdateUser(t *testing.T) {
	db,mock,err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := "UPDATE users SET username=$2,email=$3 WHERE id=$1"

	mock.ExpectQuery(query).WithArgs(1,"kotik","kek@mail.ru").WillReturnRows()
	sqlxDb:= sqlx.NewDb(db,"sqlmock")
	defer sqlxDb.Close()

	rep := NewPostgresUserRepository(sqlxDb)
	user := models.User{ID: 1,Username: "kotik",Email: "kek@mail.ru"}
	err = rep.UpdateUser(user)
	assert.NoError(t,err)
	assert.NotNil(t,user)
}


func TestUpdatePassword(t *testing.T) {
	db,mock,err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := "UPDATE users SET  password=$2 WHERE id=$1"

	mock.ExpectQuery(query).WithArgs(1,"12345").WillReturnRows()
	sqlxDb:= sqlx.NewDb(db,"sqlmock")
	defer sqlxDb.Close()

	rep := NewPostgresUserRepository(sqlxDb)
	user := models.User{ID: 1,Password: "12345"}
	err = rep.UpdatePassword(user)
	assert.NoError(t,err)
	assert.NotNil(t,user)
}



func TestUpdateAvatar(t *testing.T) {
	db,mock,err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := "UPDATE users SET avatar=$2 WHERE id=$1"

	mock.ExpectQuery(query).WithArgs(1,"src/kek.jpg").WillReturnRows()
	sqlxDb:= sqlx.NewDb(db,"sqlmock")
	defer sqlxDb.Close()

	rep := NewPostgresUserRepository(sqlxDb)
	user := models.User{ID: 1,Avatar: "src/kek.jpg"}
	err = rep.UpdateAvatar(user)
	assert.NoError(t,err)
	assert.NotNil(t,user)
}

func TestUpdateAddUser(t *testing.T) {
	db,mock,err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	query := "INSERT INTO users VALUES (default, $1, $2,$3,$4) RETURNING id"

	mock.ExpectQuery(query).WithArgs( "kotik", "kek@mail.ru", "12345", "src/kek.jpg").WillReturnRows(rows)
	sqlxDb:= sqlx.NewDb(db,"sqlmock")
	defer sqlxDb.Close()

	rep := NewPostgresUserRepository(sqlxDb)
	user := models.User{ID: 1, Username:"kotik", Email:"kek@mail.ru", Password:"12345", Avatar: "src/kek.jpg"}

	user, err = rep.Add(user)
	assert.NoError(t,err)
	assert.NotNil(t,user)
}


