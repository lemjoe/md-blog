package article

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lemjoe/md-blog/internal/models"
)

type articleSchema struct {
	Title            string    `json:"article_title"`
	Author           string    `json:"article_author"`
	AuthorId         string    `json:"author_id"`
	CreationDate     time.Time `json:"creation_date"`
	ModificationDate time.Time `json:"modification_date"`
	IsLocked         bool      `json:"is_locked"`
	Id               string    `json:"id"`
}

type Article struct {
	tableName string
	pool      *pgxpool.Pool
}

func Init(pool *pgxpool.Pool) (*Article, error) {
	collection := Article{
		tableName: "articles",
		pool:      pool,
	}
	sql := "CREATE TABLE IF NOT EXISTS articles(article_title VARCHAR (255) UNIQUE NOT NULL, article_author VARCHAR (255) NOT NULL, author_id VARCHAR (255) NOT NULL, creation_date TIMESTAMP NOT NULL, modification_date TIMESTAMP NOT NULL, is_locked boolean NOT NULL, id SERIAL PRIMARY KEY NOT NULL)"

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

func (a *Article) CreateArticle(article models.Article) (models.Article, error) {

	return models.Article{}, nil
}

func (a *Article) GetAllArticles() ([]models.Article, error) {
	var findedArticles []models.Article
	return findedArticles, nil
}

func (a *Article) GetArticleById(id string) (models.Article, error) {
	return models.Article{}, nil
}

func (a *Article) DeleteArticleById(id string) error {
	return nil
}

func (a *Article) UpdateArticleById(id string) error {
	return nil
}

func (a *Article) LockArticleById(id string) error {
	return nil
}
