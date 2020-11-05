package wishlistrepository

import (
	"fmt"

	hotelmodel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/hotels/models"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/sqlrequests"
	wishlistModel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/wishlist/models"
	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/error"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/serverError"
	"github.com/jmoiron/sqlx"
)

type PostgreWishlistRepository struct {
	conn *sqlx.DB
}

func NewPostgreWishlistRepository(conn *sqlx.DB) PostgreWishlistRepository {
	return PostgreWishlistRepository{conn}
}

func (s *PostgreWishlistRepository) GetWishlist(wishlistID int) ([]hotelmodel.MiniHotel, error) {
	bb := []hotelmodel.MiniHotel{}

	err := s.conn.Select(&bb, sqlrequests.GetWishlistPostgreRequest, wishlistID)

	if err != nil {
		return bb, customerror.NewCustomError(err, serverError.ServerInternalError, nil)
	}
	return bb, nil
}

func (s *PostgreWishlistRepository) CreateWishlist(wishlist wishlistModel.Wishlist) error {
	err := s.conn.MustExec(sqlrequests.CreateWishlistPostgreRequest, wishlist.WishistID, wishlist.Name, wishlist.UserID)
	if err != nil {
		fmt.Errorf("Error while creating wishlist, %w", err)
	}
	return nil
}

func (s *PostgreWishlistRepository) DeleteWishlist(wishlistID int) error {
	err := s.conn.MustExec(sqlrequests.DeleteWishlistPostgreRequest, wishlistID)
	if err != nil {
		fmt.Errorf("Error while deleting wishlist, %w", err)
	}
	return nil
}

func (s *PostgreWishlistRepository) AddHotel(hotelID int, wishlistID int) error {
	err := s.conn.MustExec(sqlrequests.AddHotelToWishlistPostgreRequest, wishlistID, hotelID)
	if err != nil {
		fmt.Errorf("Error while adding hotel to wishlist, %w", err)
	}
	return nil
}

func (s *PostgreWishlistRepository) DeleteHotel(hotelID int, wishlistID int) error {
	err := s.conn.MustExec(sqlrequests.DeleteHotelFromWishlistPostgreRequest, wishlistID, hotelID)
	if err != nil {
		fmt.Errorf("Error while deleting hotel from wishlist, %w", err)
	}
	return nil
}

func (s *PostgreWishlistRepository) GetTable() ([]wishlistModel.WishlisstHotel, error) {
	bb := []wishlistModel.WishlisstHotel{}
	err := s.conn.Select(&bb, "SELECT * FROM wishlistshotels")

	if err != nil {
		return bb, customerror.NewCustomError(err, serverError.ServerInternalError, nil)
	}

	return bb, nil
}
