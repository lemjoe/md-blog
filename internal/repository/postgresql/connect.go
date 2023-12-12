package postgresql

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lemjoe/md-blog/internal/repository/postgresql/article"
	"github.com/lemjoe/md-blog/internal/repository/postgresql/user"
	"github.com/lemjoe/md-blog/internal/repository/repotypes"
)

type DB struct {
	Pool *pgxpool.Pool
}

func ConnectDB(url, dbname string, user string, password string) (*DB, error) {
	url = "postgresql://" + user + ":" + password + "@" + url + "?sslmode=disable"

	fmt.Printf("url: %s, dbname: %s\n", url, dbname)

	dbpool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		log.Fatal("Unable to create connection pool: ", err)
	}
	err = dbpool.Ping(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()

	fmt.Println("Connected to PostgreSQL Database!")
	databaseInstance := &DB{
		Pool: dbpool,
	}

	return databaseInstance, nil
}

func (db *DB) Close() {
	db.Pool.Close()
}

func (db *DB) NewRepository() (*repotypes.Repository, error) {
	user, err := user.Init(db.Pool)
	if err != nil {
		return nil, err
	}
	art, err := article.Init(db.Pool)
	if err != nil {
		return nil, err
	}
	return &repotypes.Repository{
		User:    user,
		Article: art,
	}, nil
}
