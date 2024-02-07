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

type Queries struct {
	selectArticle string
}

var queries = Queries{
	selectArticle: `SELECT article_id, title, preview_text, article_time, tag, detail_text, href, comments, category, image_path
		FROM
		article INNER JOIN image
		ON article.image_id = image.image_id
		ORDER BY article_time DESC
		LIMIT ?;`,
}

func (r *MySQLArticleRepository) GetAllArticles() ([]model.Article, error) {
	articles := []model.Article{{ArticleID: 1}, {ArticleID: 2}}
	return articles, nil
}

// GetHomeCarouselArticles get articles for carousel on home page
func (r *MySQLArticleRepository) GetHomeCarouselArticles() ([]model.Article, error) {
	const op = "repository.GetHomeCarouselArticles()"

	var as []model.Article

	rows, err := r.db.Query(queries.selectArticle, 4)
	if err != nil {
		r.l.Error().Msgf("%s: query select articles for carousel > %s", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var a model.Article
		err = rows.Scan(&a.ArticleID, &a.Title, &a.PreviewText,
			&a.ArticleTime, &a.Tag, &a.DetailText, &a.Href, &a.Comments, &a.Category, &a.Image)
		if err != nil {
			r.l.Error().Msgf("%s: query scan articles for carousel > %s", op, err)
			return nil, err
		}

		as = append(as, a)
	}

	return as, nil
}

// GetHomeTrendingArticlesTop return last four articles with // TODO large number of comments
func (r *MySQLArticleRepository) GetHomeTrendingArticlesTop() ([]model.Article, error) {
	const op = "repository.GetHomeTrendingArticlesTop()"

	var as []model.Article

	rows, err := r.db.Query(queries.selectArticle, 4)
	if err != nil {
		r.l.Error().Msgf("%s: query select articles for carousel > %s", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var a model.Article
		err = rows.Scan(&a.ArticleID, &a.Title, &a.PreviewText,
			&a.ArticleTime, &a.Tag, &a.DetailText, &a.Href, &a.Comments, &a.Category, &a.Image)
		if err != nil {
			r.l.Error().Msgf("%s: query scan articles for carousel > %s", op, err)
			return nil, err
		}

		as = append(as, a)
	}

	return as, nil
}

// GetHomeTrendingArticlesBottom return last four articles with // TODO large number of comments
func (r *MySQLArticleRepository) GetHomeTrendingArticlesBottom() ([]model.Article, error) {
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
