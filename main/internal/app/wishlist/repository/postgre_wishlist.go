package wishlistrepository

import (
	"fmt"

	"github.com/go-park-mail-ru/2020_2_JMickhs/main/configs"
	"github.com/spf13/viper"

	wishlistModel "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/wishlist/models"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/clientError"
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

func (s *PostgreWishlistRepository) GetWishlistMeta(wishlistID int) ([]wishlistModel.WishlistHotel, error) {
	bb := []wishlistModel.WishlistHotel{}
	err := s.conn.Select(&bb, GetWishlistMeta, wishlistID)
	if err != nil {
		return bb, customerror.NewCustomError(err, clientError.BadRequest, 1)
	}
	return bb, nil
}

func (s *PostgreWishlistRepository) CreateWishlist(wishlist wishlistModel.Wishlist) (wishlistModel.Wishlist, error) {
	err := s.conn.QueryRow(CreateWishlistPostgreRequest, wishlist.Name, wishlist.UserID).Scan(&wishlist.WishlistID)
	if err != nil {
		return wishlist, customerror.NewCustomError(err, clientError.BadRequest, 1)
	}
	return wishlist, nil
}

func (s *PostgreWishlistRepository) DeleteWishlist(wishlistID int) error {
	_, err := s.conn.Query(DeleteWishlistPostgreRequest, wishlistID)
	if err != nil {
		return customerror.NewCustomError(err, clientError.BadRequest, 1)
	}
	return nil
}

func (s *PostgreWishlistRepository) AddHotel(hotelID int, wishlistID int) error {
	_, err := s.conn.Query(AddHotelToWishlistPostgreRequest, wishlistID, hotelID)
	if err != nil {
		return customerror.NewCustomError(err, clientError.Conflict, 1)
	}
	return nil
}

func (s *PostgreWishlistRepository) DeleteHotel(hotelID int, wishlistID int) error {
	_, err := s.conn.Query(DeleteHotelFromWishlistPostgreRequest, wishlistID, hotelID)
	if err != nil {
		return customerror.NewCustomError(err, clientError.BadRequest, 1)
	}
	return nil
}

func (s *PostgreWishlistRepository) CheckWishListOwner(wishListID int, UserID int) (bool, error) {
	checkUserID := -1
	err := s.conn.QueryRow(CheckWishListOwnerPostgreRequest, wishListID).Scan(&checkUserID)
	if err != nil {
		return false, customerror.NewCustomError(err, clientError.BadRequest, 1)
	}
	if checkUserID == UserID {
		return true, nil
	}
	return false, nil
}

func (s *PostgreWishlistRepository) GetUserWishlists(UserID int) (wishlistModel.UserWishLists, error) {
	wishLists := wishlistModel.UserWishLists{}
	err := s.conn.Select(&wishLists.Wishlists, GetUserWithListsPostgreRequest, UserID)
	if err != nil {
		return wishLists, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return wishLists, nil
}

func (s *PostgreWishlistRepository) CheckHotelInWishlists(userID int, hotelID int) (string, error) {

	res, err := s.conn.Exec(CheckHotelInWishlistsPostgreRequest, userID, hotelID)
	if err != nil {
		return "", customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	count, _ := res.RowsAffected()
	fmt.Println(count, userID, hotelID)
	if count > 0 {
		return viper.GetString(configs.ConfigFields.WishListIn), nil
	}
	return viper.GetString(configs.ConfigFields.WishListOut), nil
}
