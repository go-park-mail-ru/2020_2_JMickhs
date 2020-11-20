package wishlistrepository

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	hotelmodel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/hotels/models"
	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/error"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/serverError"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestGetWishlist(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	t.Run("GetWishlist", func(t *testing.T) {
		rowsHotel := sqlmock.NewRows([]string{"hotel_id", "name", "description", "img", "location", "curr_rating"}).AddRow(
			1, "Hotel", "beautiful hotel", "static/img/hotelImg1.jpg", "Moscow", 0).AddRow(
			2, "bauman na baumanskoy", "bar bauman", "static/img/hotelImg2.jpg", "Moscow", 0)

		query := "SELECT h.hotel_id,h.name,h.description,h.img,h.location,h.curr_rating FROM hotels AS h LEFT JOIN wishlistshotels AS wh ON h.hotel_id = wh.hotel_id LEFT JOIN wishlists AS w ON w.wishlist_id = wh.wishlist_id WHERE wh.wishlist_id = $1"

		hotelTest := hotelmodel.MiniHotel{1, "Hotel", "beautiful hotel", "static/img/hotelImg1.jpg", "Moscow", 0}
		mock.ExpectQuery(query).WithArgs(42).WillReturnRows(rowsHotel)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgreWishlistRepository(sqlxDb)

		hotels, err := rep.GetWishlist(42)
		assert.NoError(t, err)
		assert.Equal(t, hotels[0], hotelTest)
	})
	t.Run("GetWishlistErr", func(t *testing.T) {
		query := "SELECT h.hotel_id,h.name,h.description,h.img,h.location,h.curr_rating FROM hotels AS h LEFT JOIN wishlistshotels AS wh ON h.hotel_id = wh.hotel_id LEFT JOIN wishlists AS w ON w.wishlist_id = wh.wishlist_id WHERE wh.wishlist_id = $1"

		mock.ExpectQuery(query).WithArgs(42).WillReturnError(customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgreWishlistRepository(sqlxDb)

		_, err := rep.GetWishlist(42)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})
}
