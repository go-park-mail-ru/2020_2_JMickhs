package wishlistrepository

import (
	"errors"
	"testing"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/serverError"

	wishlistmodel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/wishlist/models"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/clientError"

	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/package/error"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestPostgreWishlistRepository_GetWishlistMeta(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	t.Run("GetWishlist", func(t *testing.T) {
		rowsHotel := sqlmock.NewRows([]string{"wishlist_id", "hotel_id"}).AddRow(
			1, 4).AddRow(
			1, 3)

		query := GetWishlistMeta

		hotelTest := wishlistmodel.WishlistHotel{1, 4}
		mock.ExpectQuery(query).
			WithArgs(1).
			WillReturnRows(rowsHotel)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgreWishlistRepository(sqlxDb)

		hotels, err := rep.GetWishlistMeta(1)
		assert.NoError(t, err)
		assert.Equal(t, hotels[0], hotelTest)
	})
	t.Run("GetWishlistErr", func(t *testing.T) {
		query := GetWishlistMeta

		mock.ExpectQuery(query).
			WithArgs(1).
			WillReturnError(customerror.NewCustomError(errors.New(""), clientError.BadRequest, 1))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgreWishlistRepository(sqlxDb)

		_, err := rep.GetWishlistMeta(1)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), clientError.BadRequest)
	})
}

func TestPostgreWishlistRepository_GetUserWishlists(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	t.Run("GetUserWishlists", func(t *testing.T) {
		rowsHotel := sqlmock.NewRows([]string{"wishlist_id", "name"}).AddRow(
			1, "kekw").AddRow(
			2, "kterw")

		query := GetUserWithListsPostgreRequest

		testWishlists := wishlistmodel.Wishlist{WishlistID: 1, Name: "kekw"}
		mock.ExpectQuery(query).
			WithArgs(1).
			WillReturnRows(rowsHotel)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgreWishlistRepository(sqlxDb)

		wishLists, err := rep.GetUserWishlists(1)
		assert.NoError(t, err)
		assert.Equal(t, testWishlists, wishLists.Wishlists[0])
	})
	t.Run("GetUserWishlistsErr", func(t *testing.T) {
		query := GetWishlistMeta

		mock.ExpectQuery(query).
			WithArgs(1).
			WillReturnError(customerror.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgreWishlistRepository(sqlxDb)

		_, err := rep.GetUserWishlists(1)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), serverError.ServerInternalError)
	})
}

func TestPostgreWishlistRepository_DeleteWishlist(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	t.Run("DeleteWishlist", func(t *testing.T) {
		query := DeleteWishlistPostgreRequest

		mock.ExpectQuery(query).
			WithArgs(1).
			WillReturnRows()

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgreWishlistRepository(sqlxDb)

		err := rep.DeleteWishlist(1)
		assert.NoError(t, err)
	})
	t.Run("DeleteWishlistErr", func(t *testing.T) {
		query := DeleteWishlistPostgreRequest

		mock.ExpectQuery(query).
			WithArgs(1).
			WillReturnError(customerror.NewCustomError(errors.New(""), clientError.BadRequest, 1))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgreWishlistRepository(sqlxDb)

		err := rep.DeleteWishlist(1)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), clientError.BadRequest)
	})
}

func TestPostgreWishlistRepository_DeleteHotel(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	t.Run("DeleteHotel", func(t *testing.T) {
		query := DeleteHotelFromWishlistPostgreRequest

		mock.ExpectQuery(query).WithArgs(1, 3).WillReturnRows()

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgreWishlistRepository(sqlxDb)

		err := rep.DeleteHotel(3, 1)
		assert.NoError(t, err)
	})
	t.Run("DeleteHotelErr", func(t *testing.T) {
		query := DeleteHotelFromWishlistPostgreRequest

		mock.ExpectQuery(query).
			WithArgs(1, 3).
			WillReturnError(customerror.NewCustomError(errors.New(""), clientError.BadRequest, 1))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgreWishlistRepository(sqlxDb)

		err := rep.DeleteHotel(3, 1)
		assert.Error(t, err)
	})
}

func TestPostgreWishlistRepository_CreateWishlist(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	t.Run("CreateWishlist", func(t *testing.T) {

		rowsWishlist := sqlmock.NewRows([]string{"wishlist_id"}).AddRow(
			1)
		query := CreateWishlistPostgreRequest

		wishlistTest := wishlistmodel.Wishlist{1, "kekw", 4}
		mock.ExpectQuery(query).
			WithArgs("kekw", 4).
			WillReturnRows(rowsWishlist)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgreWishlistRepository(sqlxDb)

		wishlist, err := rep.CreateWishlist(wishlistTest)
		assert.NoError(t, err)
		assert.Equal(t, wishlist, wishlistTest)
	})
	t.Run("CreateWishlistErr", func(t *testing.T) {
		query := CreateWishlistPostgreRequest

		wishlistTest := wishlistmodel.Wishlist{1, "kekw", 4}
		mock.ExpectQuery(query).
			WithArgs(1, "kekw", 4).
			WillReturnError(customerror.NewCustomError(errors.New(""), clientError.BadRequest, 1))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgreWishlistRepository(sqlxDb)

		_, err := rep.CreateWishlist(wishlistTest)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), clientError.BadRequest)
	})
}

func TestPostgreWishlistRepository_CheckWishListOwner(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	t.Run("CheckWishListOwner", func(t *testing.T) {
		rowsUser := sqlmock.NewRows([]string{"user_id"}).AddRow(
			3)

		query := CheckWishListOwnerPostgreRequest

		mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rowsUser)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgreWishlistRepository(sqlxDb)

		res, err := rep.CheckWishListOwner(1, 3)
		assert.NoError(t, err)
		assert.Equal(t, res, true)
	})
	db, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	t.Run("CheckWishListOwner", func(t *testing.T) {
		rowsUser := sqlmock.NewRows([]string{"user_id"}).AddRow(
			3)

		query := CheckWishListOwnerPostgreRequest

		mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rowsUser)

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgreWishlistRepository(sqlxDb)

		res, err := rep.CheckWishListOwner(1, 2)
		assert.NoError(t, err)
		assert.Equal(t, res, false)
	})
	t.Run("CheckWishListOwnerErr", func(t *testing.T) {
		query := CheckWishListOwnerPostgreRequest

		mock.ExpectQuery(query).WithArgs(1).WillReturnError(customerror.NewCustomError(errors.New(""), clientError.BadRequest, 1))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgreWishlistRepository(sqlxDb)

		_, err := rep.CheckWishListOwner(1, 4)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), clientError.BadRequest)
	})
}

func TestPostgreWishlistRepository_AddHotel(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	t.Run("AddHotel", func(t *testing.T) {
		query := AddHotelToWishlistPostgreRequest

		mock.ExpectQuery(query).WithArgs(1, 4).WillReturnRows()

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgreWishlistRepository(sqlxDb)

		err := rep.AddHotel(4, 1)
		assert.NoError(t, err)
	})
	t.Run("AddHotelErr", func(t *testing.T) {
		query := GetWishlistMeta

		mock.ExpectQuery(query).WithArgs(1).WillReturnError(customerror.NewCustomError(errors.New(""), clientError.Conflict, 1))

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		defer sqlxDb.Close()

		rep := NewPostgreWishlistRepository(sqlxDb)

		err := rep.AddHotel(4, 1)
		assert.Error(t, err)
		assert.Equal(t, customerror.ParseCode(err), clientError.Conflict)
	})
}
