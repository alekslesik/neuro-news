package repository

import (
	"database/sql"

	"github.com/alekslesik/neuro-news/internal/app/model"
)

type MySQLArticleRepository struct {
	db *sql.DB
}

func (r *MySQLArticleRepository) GetAllArticles() ([]model.Article, error) {
	articles := []model.Article{{Article_id: 1}, {Article_id: 2}}
	return articles, nil
}

func (r *MySQLArticleRepository) GetHomeCarouselArticles() ([]model.Article, error) {
	return nil, nil
}

func (r *MySQLArticleRepository) GetHomeTrendingArticles() ([]model.Article, error) {
	return nil, nil
}

func (r *MySQLArticleRepository) GetHomeNewsArticles() ([]model.Article, error) {
	return nil, nil
}

func (r *MySQLArticleRepository) GetHomeSportArticles() ([]model.Article, error) {
	return nil, nil
}

func (r *MySQLArticleRepository) GetHomeVideoArticles() ([]model.Article, error) {
	return nil, nil
}

func (r *MySQLArticleRepository) GetHomePopularArticles() ([]model.Article, error) {
	return nil, nil
}

func (r *MySQLArticleRepository) GetArticleByID(id int) (*model.Article, error) {
	return nil, nil
}
