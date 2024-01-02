package repository

import (
	"database/sql"

	"github.com/alekslesik/neuro-news/internal/app/model"
)

type MySQLArticleRepository struct {
	db *sql.DB
}

func (r *MySQLArticleRepository) GetAllArticles() ([]model.Article, error) {
	articles := []model.Article{{ID: 1}, {ID: 2}}
	return articles, nil
}

func (r *MySQLArticleRepository) GetArticleByID(id int) (*model.Article, error) {
	// Реализация получения статьи по ID из базы данных
	return nil, nil
}
