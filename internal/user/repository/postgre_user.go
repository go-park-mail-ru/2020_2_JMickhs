package userRepository

import (
	"strconv"

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
	err := p.conn.QueryRow("INSERT INTO users VALUES (default, $1, $2,$3,$4) RETURNING id", user.Username, user.Email, user.Password, user.Avatar).Scan(&id)
	user.ID = id
	return user, err
}

func (p *PostgresUserRepository) GetByUserName(name string) (models.User, error) {
	rows, err := p.conn.Query("select id,username,email,password,avatar FROM users WHERE username=$1", name)
	defer rows.Close()
	user := models.User{}
	if err != nil {
		return user, err
	}

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Avatar)
		if err != nil {
			return user, err
		}
	}
	return user, nil
}

func (p *PostgresUserRepository) GetUserByID(ID int) (models.User, error) {
	rows, err := p.conn.Query("SELECT id,username,email,password,avatar FROM users WHERE id=$1", strconv.Itoa(ID))
	defer rows.Close()
	user := models.User{}
	if err != nil {
		return user, err
	}

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Avatar)
		if err != nil {
			return user, err
		}
	}
	return user, nil
}

func (p *PostgresUserRepository) UpdateUser(user models.User) error {
	_, err := p.conn.Query("UPDATE users SET username=$2,email=$3 WHERE id=$1",
		user.ID, user.Username, user.Email)
	if err != nil {
		return err
	}
	return nil

}

func (p *PostgresUserRepository) UpdateAvatar(user models.User) error {
	_, err := p.conn.Query("UPDATE users SET avatar=$2 WHERE id=$1",
		user.ID, user.Avatar)
	if err != nil {
		return err
	}
	return nil

}

func (p *PostgresUserRepository) UpdatePassword(user models.User) error {
	_, err := p.conn.Query("UPDATE users SET  password=$2 WHERE id=$1",
		user.ID, user.Password)
	if err != nil {
		return err
	}
	return nil

}
