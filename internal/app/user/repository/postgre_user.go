package userRepository

import (
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/sqlrequests"

	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/error"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/user/models"
	"github.com/jmoiron/sqlx"
)

type PostgresUserRepository struct {
	conn *sqlx.DB
}

func NewPostgresUserRepository(conn *sqlx.DB) PostgresUserRepository {
	userStorage := PostgresUserRepository{conn}
	return userStorage
}

func (p *PostgresUserRepository) Add(user models.User) (models.User, error) {
	var id int
	err := p.conn.QueryRow(sqlrequests.AddUserPostgreRequest, user.Username, user.Email, user.Password, user.Avatar).Scan(&id)
	if err != nil {
		return user, customerror.NewCustomError(err, http.StatusConflict, nil)
	}
	user.ID = id
	return user, nil
}

func (p *PostgresUserRepository) GetByUserName(name string) (models.User, error) {
	rows, err := p.conn.Query(sqlrequests.GetUserByNamePostgreRequest, name)
	defer rows.Close()
	user := models.User{}
	if err != nil {
		return user, customerror.NewCustomError(err, http.StatusUnauthorized, nil)
	}

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Avatar)
		if err != nil {
			return user, customerror.NewCustomError(err, http.StatusInternalServerError, nil)
		}
	}
	return user, nil
}

func (p *PostgresUserRepository) GetUserByID(ID int) (models.User, error) {
	row := p.conn.QueryRow(sqlrequests.GetUserByIDPostgreRequest, strconv.Itoa(ID))
	user := models.User{}

	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Avatar)
	if err != nil {
		return user, customerror.NewCustomError(err, http.StatusGone, nil)
	}
	return user, nil
}

func (p *PostgresUserRepository) UpdateUser(user models.User) error {
	_, err := p.conn.Query(sqlrequests.UpdateUserCredPostgreRequest,
		user.ID, user.Username, user.Email)
	if err != nil {
		return customerror.NewCustomError(err, http.StatusConflict, nil)
	}
	return nil

}

func (p *PostgresUserRepository) UpdateAvatar(user models.User) error {
	_, err := p.conn.Query(sqlrequests.UpdateUserAvatarPostgreRequest,
		user.ID, user.Avatar)
	if err != nil {
		return customerror.NewCustomError(err, http.StatusInternalServerError, nil)
	}
	return nil

}

func (p *PostgresUserRepository) UpdatePassword(user models.User) error {
	_, err := p.conn.Query(sqlrequests.UpdateUserPasswordPostgreRequest,
		user.ID, user.Password)
	if err != nil {
		return customerror.NewCustomError(err, http.StatusInternalServerError, nil)
	}
	return nil

}
