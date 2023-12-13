package repository

import (
	"database/sql"

	"github.com/alekslesik/neuro-news/internal/app/model"
)

type MySQLRepository struct {
    db *sql.DB
}

func NewMySQLRepository(db *sql.DB) model.ArticleRepository {
    return &MySQLRepository{
        db: db,
    }
}

func (r *MySQLRepository) GetAllArticles() ([]model.Article, error) {
    return nil, nil
}

func (r *MySQLRepository) GetArticleByID(id int) (*model.Article, error) {
	return nil, nil
}
