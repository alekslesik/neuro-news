package model

import (
	"database/sql"

	service "github.com/alekslesik/neuro-news/internal/app/service/mysql"
)

func New(db *sql.DB) *Model {
	return &Model{
		Article: &service.ArticleModel{DB: db},
		Users: &service.UserModel{DB: db},
	}
}

type Model struct {
	Articles interface {
		Insert(name, email, password string) error
		Authenticate(email, password string) (int, string, error)
		Get(id int) (*models.Article, error)
	}
	Users interface {
		Insert(name, email, password string) error
		Authenticate(email, password string) (int, string, error)
		Get(id int) (*models.User, error)
	}
}