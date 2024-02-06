package repository

import (
	"database/sql"

	"github.com/alekslesik/neuro-news/internal/app/model"
	"github.com/alekslesik/neuro-news/pkg/logger"
)

type MySQLArticleRepository struct {
	db *sql.DB
	l  *logger.Logger
}

func (r *MySQLArticleRepository) GetAllArticles() ([]model.Article, error) {
	articles := []model.Article{{ArticleID: 1}, {ArticleID: 2}}
	return articles, nil
}

// GetHomeCarouselArticles get articles for carousel on home page
func (r *MySQLArticleRepository) GetHomeCarouselArticles() ([]model.Article, error) {
	const op = "repository.GetHomeCarouselArticles()"

	var as []model.Article

	q := `SELECT article_id, title, preview_text, image_id, date, tag, detail_text
		FROM neuronews.article
		ORDER BY date DESC
		LIMIT ?;`

	rows, err := r.db.Query(q, 4)
	if err != nil {
		r.l.Error().Msgf("%s: query select articles for carousel > %s", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var a model.Article
		err = rows.Scan(&a.ArticleID, &a.Title, &a.PreviewText, &a.Image, &a.Date, &a.Tag, &a.DetailText)
		if err != nil {
			r.l.Error().Msgf("%s: query scan articles for carousel > %s", op, err)
			return nil, err
		}

		as = append(as, a)
	}

	return as, nil
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
