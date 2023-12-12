package user

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lemjoe/md-blog/internal/models"
)

type userSchema struct {
	UserName     string            `json:"user_name"`
	FullName     string            `json:"full_name"`
	Password     string            `json:"passwd"`
	Email        string            `json:"email"`
	IsAdmin      bool              `json:"is_admin"`
	Id           string            `json:"id"`
	LastLogin    time.Time         `json:"last_login"`
	CreationDate time.Time         `json:"creation_date"`
	Settings     map[string]string `json:"settings"`
}

type User struct {
	tableName string
	pool      *pgxpool.Pool
}

func Init(pool *pgxpool.Pool) (*User, error) {
	collection := User{
		tableName: "users",
		pool:      pool,
	}
	sql := "CREATE TABLE IF NOT EXISTS users(user_name VARCHAR (255) UNIQUE NOT NULL, full_name VARCHAR (255) NOT NULL, passwd VARCHAR (255) NOT NULL, email VARCHAR (255) NOT NULL, is_admin boolean NOT NULL, id SERIAL PRIMARY KEY NOT NULL, last_login TIMESTAMP NOT NULL, creation_date TIMESTAMP NOT NULL, settings VARCHAR (255) NOT NULL)"

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(), sql)
	if err != nil {
		return nil, err
	}
	return &collection, nil
}

func (u *User) CreateUser(user models.User) (models.User, error) {
	return models.User{}, nil
}

func (u *User) GetUserByUsername(username string) (models.User, error) {
	return models.User{}, nil
}

func (u *User) GetUserById(id string) (models.User, error) {
	return models.User{}, nil
}

func (u *User) ChangeUserSettings(id string, settings map[string]string) error {
	return nil
}
