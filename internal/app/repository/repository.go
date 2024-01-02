package repository

import (
	"database/sql"

	"github.com/alekslesik/neuro-news/internal/app/model"
)

type MySQLRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{
		db: db,
	}
}

func (r *MySQLRepository) GetArticleRepository() model.ArticleModel {
	return &MySQLArticleRepository{db: r.db}
}

func (r *MySQLRepository) GetUserRepository() model.UserModel {
	return &MySQLUserRepository{db: r.db}
}
