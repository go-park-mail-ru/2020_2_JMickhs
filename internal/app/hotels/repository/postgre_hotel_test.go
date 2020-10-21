package hotelRepository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetHoteBytID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "name", "description", "img"}).AddRow(
		1, "hotel", "top hotel in the world", "src/kek.jpg")

	query := "SELECT id,name,description,img FROM hotels WHERE id=$1"

	mock.ExpectQuery(query).WithArgs("1").WillReturnRows(rows)
	sqlxDb := sqlx.NewDb(db, "sqlmock")
	defer sqlxDb.Close()

	rep := NewPostgresHotelRepository(sqlxDb)

	hotel, err := rep.GetHotelByID(1)
	assert.NoError(t, err)
	assert.NotNil(t, hotel)
	assert.Equal(t, hotel.Name, "hotel")
}

func TestGetHotels(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "name", "description", "img"}).AddRow(
		1, "hotel", "top hotel in the world", "src/kek.jpg")

	query := "SELECT id,name,description,img FROM hotels"

	mock.ExpectQuery(query).WillReturnRows(rows)
	sqlxDb := sqlx.NewDb(db, "sqlmock")
	defer sqlxDb.Close()

	rep := NewPostgresHotelRepository(sqlxDb)

	hotel, err := rep.GetHotels()
	assert.NoError(t, err)
	assert.NotNil(t, hotel)

	assert.Equal(t, hotel[0].Name, "hotel")
}
