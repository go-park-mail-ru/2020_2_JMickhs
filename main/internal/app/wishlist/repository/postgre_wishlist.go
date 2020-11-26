package wishlistrepository

import (
	wishlistModel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/wishlist/models"
	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/package/error"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/serverError"
	"github.com/jmoiron/sqlx"
)

type PostgreWishlistRepository struct {
	conn *sqlx.DB
}

func NewPostgreWishlistRepository(conn *sqlx.DB) PostgreWishlistRepository {
	return PostgreWishlistRepository{conn}
}

func (s *PostgreWishlistRepository) GetWishlistMeta(wishlistID int) ([]wishlistModel.WishlisstHotel, error) {
	bb := []wishlistModel.WishlisstHotel{}
	err := s.conn.Select(&bb, GetWishlistMeta, wishlistID)
	if err != nil {
		return bb, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return bb, nil
}

func (s *PostgreWishlistRepository) CreateWishlist(wishlist wishlistModel.Wishlist) error {
	_, err := s.conn.Query(CreateWishlistPostgreRequest, wishlist.WishistID, wishlist.Name, wishlist.UserID)
	if err != nil {
		return customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return nil
}

func (s *PostgreWishlistRepository) DeleteWishlist(wishlistID int) error {
	_, err := s.conn.Query(DeleteWishlistPostgreRequest, wishlistID)
	if err != nil {
		return customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return nil
}

func (s *PostgreWishlistRepository) AddHotel(hotelID int, wishlistID int) error {
	_, err := s.conn.Query(AddHotelToWishlistPostgreRequest, wishlistID, hotelID)
	if err != nil {
		return customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return nil
}

func (s *PostgreWishlistRepository) DeleteHotel(hotelID int, wishlistID int) error {
	_, err := s.conn.Query(DeleteHotelFromWishlistPostgreRequest, wishlistID, hotelID)
	if err != nil {
		return customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return nil
}
