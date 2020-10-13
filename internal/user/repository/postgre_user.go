package userRepository

import (
	"strconv"

	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/internal/error"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/user/models"
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
	err := p.conn.QueryRow("INSERT INTO users VALUES (default, $1, $2,$3,$4) RETURNING user_id", user.Username, user.Email, user.Password, user.Avatar).Scan(&id)
	if err != nil {
		return user, customerror.NewCustomError(err.Error())
	}
	user.ID = id
	return user, nil
}

func (p *PostgresUserRepository) GetByUserName(name string) (models.User, error) {
	rows, err := p.conn.Query("select user_id,username,email,password,avatar FROM users WHERE username=$1", name)
	defer rows.Close()
	user := models.User{}
	if err != nil {
		return user, customerror.NewCustomError(err.Error())
	}

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Avatar)
		if err != nil {
			return user, customerror.NewCustomError(err.Error())
		}
	}
	return user, nil
}

func (p *PostgresUserRepository) GetUserByID(ID int) (models.User, error) {
	row := p.conn.QueryRow("SELECT user_id,username,email,password,avatar FROM users WHERE user_id=$1", strconv.Itoa(ID))
	user := models.User{}

	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Avatar)
	if err != nil {
		return user, customerror.NewCustomError("such user doesn't exist")
	}
	return user, nil
}

func (p *PostgresUserRepository) UpdateUser(user models.User) error {
	_, err := p.conn.Query("UPDATE users SET username=$2,email=$3 WHERE user_id=$1",
		user.ID, user.Username, user.Email)
	if err != nil {
		return customerror.NewCustomError(err.Error())
	}
	return nil

}

func (p *PostgresUserRepository) UpdateAvatar(user models.User) error {
	_, err := p.conn.Query("UPDATE users SET avatar=$2 WHERE user_id=$1",
		user.ID, user.Avatar)
	if err != nil {
		return customerror.NewCustomError(err.Error())
	}
	return nil

}

func (p *PostgresUserRepository) UpdatePassword(user models.User) error {
	_, err := p.conn.Query("UPDATE users SET  password=$2 WHERE user_id=$1",
		user.ID, user.Password)
	if err != nil {
		return customerror.NewCustomError(err.Error())
	}
	return nil

}
