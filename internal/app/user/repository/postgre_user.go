package userRepository

import (
	"strconv"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/clientError"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/serverError"

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
	err := p.conn.QueryRow(AddUserPostgreRequest, user.Username, user.Email, user.Password, user.Avatar).Scan(&id)
	if err != nil {
		return user, customerror.NewCustomError(err, clientError.Conflict, 1)
	}
	user.ID = id
	return user, nil
}

func (p *PostgresUserRepository) GetByUserName(name string) (models.User, error) {
	user := models.User{}
	err := p.conn.QueryRow(GetUserByNamePostgreRequest, name).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Avatar)
	if err != nil {
		return user, customerror.NewCustomError(err, clientError.Unauthorizied, 1)
	}
	return user, nil
}

func (p *PostgresUserRepository) GetUserByID(ID int) (models.User, error) {
	row := p.conn.QueryRow(GetUserByIDPostgreRequest, strconv.Itoa(ID))
	user := models.User{}

	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Avatar)
	if err != nil {
		return user, customerror.NewCustomError(err, clientError.Gone, 1)
	}
	return user, nil
}

func (p *PostgresUserRepository) UpdateUser(user models.User) error {
	_, err := p.conn.Query(UpdateUserPostgreRequest,
		user.ID, user.Username, user.Email)
	if err != nil {
		return customerror.NewCustomError(err, clientError.Conflict, 1)
	}
	return nil

}

func (p *PostgresUserRepository) UpdateAvatar(user models.User) error {
	_, err := p.conn.Query(UpdateUserAvatarPostgreRequest,
		user.ID, user.Avatar)
	if err != nil {
		return customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return nil

}

func (p *PostgresUserRepository) UpdatePassword(user models.User) error {
	_, err := p.conn.Query(UpdateUserPasswordPostgreRequest,
		user.ID, user.Password)
	if err != nil {
		return customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return nil

}
