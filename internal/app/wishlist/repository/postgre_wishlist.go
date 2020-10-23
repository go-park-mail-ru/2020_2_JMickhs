package wishlistrepository

import (
	wishlistModel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/wishlist/models"
	"github.com/jmoiron/sqlx"
)

type PostgreWishlistRepository struct {
	conn *sqlx.DB
}

func NewPostgreWishlistRepository(conn *sqlx.DB) PostgreWishlistRepository {
	return PostgreWishlistRepository{conn}
}

func (s *PostgreWishlistRepository) GetWhishlist(wishlistID int) (wishlistModel.Wishlist, error) {
	panic("not implemented") // TODO: Implement
}

func (s *PostgreWishlistRepository) CreateWishlist(wishlist wishlistModel.Wishlist) error {
	panic("not implemented") // TODO: Implement
}

func (s *PostgreWishlistRepository) DeleteWishlist(wishlistID int) error {
	panic("not implemented") // TODO: Implement
}

func (s *PostgreWishlistRepository) UpdateWishlist(wishlist wishlistModel.Wishlist) error {
	panic("not implemented") // TODO: Implement
}

func (s *PostgreWishlistRepository) AddHotel(hotelID int, wishlistID int) error {
	panic("not implemented") // TODO: Implement
}

func (s *PostgreWishlistRepository) DeleteHotel(hotelID string, wishlistID int) error {
	panic("not implemented") // TODO: Implement
}
