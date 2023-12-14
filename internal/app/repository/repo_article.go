package repository

import (
	"database/sql"

	"github.com/alekslesik/neuro-news/internal/app/model"
)

type MySQLArticleRepository struct {
    db *sql.DB
}

func (r *MySQLArticleRepository) GetAllArticles() ([]model.Article, error) {
    // Реализация получения всех статей из базы данных
	return nil, nil
}

func (r *MySQLArticleRepository) GetArticleByID(id int) (*model.Article, error) {
    // Реализация получения статьи по ID из базы данных
	return nil, nil
}
