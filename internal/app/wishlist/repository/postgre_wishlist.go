package wishlistrepository

import (
	"fmt"
	"strconv"

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

func (s *PostgreWishlistRepository) GetWishlist(wishlistID int) ([]hotelmodel.Hotel, error) {
	rows, err := s.conn.Query(sqlrequests.GetWishlistPostgreRequest, strconv.Itoa(wishlistID))
	defer rows.Close()

	hotels := []hotelmodel.Hotel{}
	if err != nil {
		fmt.Errorf("Error while geting hotels from wishlists, %w", err)
		return hotels, customerror.NewCustomError(err, serverError.ServerInternalError, nil)
	}
	hotel := hotelmodel.Hotel{}
	for rows.Next() {
		err := rows.Scan(&hotel.HotelID, &hotel.Name, &hotel.Description, &hotel.Image, &hotel.Location, &hotel.Rating)
		if err != nil {
			fmt.Errorf("Error while unpacking hotel from bd, %w", err)
			return hotels, customerror.NewCustomError(err, serverError.ServerInternalError, nil)
		}
		hotels = append(hotels, hotel)
	}
	return hotels, nil
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
