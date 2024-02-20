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

type ArticleQueries struct {
	selectAllArticle   string
	selectArticleLimit string
	// WHERE category  = ?
	selectArticleWhereLimit string
	selectVideoLimit        string
}

var articleQueries = ArticleQueries{
	selectAllArticle: `SELECT article_id, title, preview_text, article_time, tag, detail_text, href, comments, category, image_path
	FROM
	article INNER JOIN image
	ON article.image_id = image.image_id
	WHERE kind = 'article'
	ORDER BY article_time DESC;`,

	selectArticleLimit: `SELECT article_id, title, preview_text, article_time, tag, detail_text, href, comments, category, image_path
	FROM
	article INNER JOIN image
	ON article.image_id = image.image_id
	WHERE kind = 'article'
	ORDER BY article_time DESC
	LIMIT ?;`,

	selectArticleWhereLimit: `SELECT article_id, title, preview_text, article_time, tag, detail_text, href, comments, category, image_path
	FROM
	article INNER JOIN image
	ON article.image_id = image.image_id
	WHERE kind = 'article' AND category  = ?
	ORDER BY article_time DESC
	LIMIT ?;`,

	selectVideoLimit: `SELECT article_id, title, preview_text, article_time, tag, detail_text, href, comments, category, video_path
	FROM
	article INNER JOIN video
	ON article.video_id = video.video_id
	WHERE kind = 'video'
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

	rows, err := r.db.Query(articleQueries.selectArticleLimit, 4)
	if err != nil {
		r.l.Warn().Msgf("%s: query select articles for carousel > %s", op, err)
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

	rows, err := r.db.Query(articleQueries.selectArticleLimit, 4)
	if err != nil {
		r.l.Error().Msgf("%s: query select trending articles top > %s", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var a model.Article
		err = rows.Scan(&a.ArticleID, &a.Title, &a.PreviewText,
			&a.ArticleTime, &a.Tag, &a.DetailText, &a.Href, &a.Comments, &a.Category, &a.Image)
		if err != nil {
			r.l.Error().Msgf("%s: query scan trending articles top > %s", op, err)
			return nil, err
		}

		as = append(as, a)
	}

	return as, nil
}

// GetHomeTrendingArticlesBottom return last six articles with // TODO large number of comments
func (r *MySQLArticleRepository) GetHomeTrendingArticlesBottom() ([]model.Article, error) {
	const op = "repository.GetHomeTrendingArticlesBottom()"

	var as []model.Article

	rows, err := r.db.Query(articleQueries.selectArticleLimit, 11)
	if err != nil {
		r.l.Error().Msgf("%s: query select trending articles bottom > %s", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var a model.Article
		err = rows.Scan(&a.ArticleID, &a.Title, &a.PreviewText,
			&a.ArticleTime, &a.Tag, &a.DetailText, &a.Href, &a.Comments, &a.Category, &a.Image)
		if err != nil {
			r.l.Error().Msgf("%s: query scan trending articles bottom > %s", op, err)
			return nil, err
		}

		as = append(as, a)
	}

	as = as[4:]

	return as, nil
}

// GetHomeNewsArticles return 3 news for news/sport block
func (r *MySQLArticleRepository) GetHomeNewsArticles() ([]model.Article, error) {
	const op = "repository.GetHomeNewsArticles()"

	var as []model.Article

	rows, err := r.db.Query(articleQueries.selectArticleLimit, 3)
	if err != nil {
		r.l.Error().Msgf("%s: query select home news articles > %s", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var a model.Article
		err = rows.Scan(&a.ArticleID, &a.Title, &a.PreviewText,
			&a.ArticleTime, &a.Tag, &a.DetailText, &a.Href, &a.Comments, &a.Category, &a.Image)
		if err != nil {
			r.l.Error().Msgf("%s: query scan home news articles > %s", op, err)
			return nil, err
		}

		as = append(as, a)
	}

	return as, nil
}

// GetHomeSportArticles return 3 sport news for news/sport block
func (r *MySQLArticleRepository) GetHomeSportArticles() ([]model.Article, error) {
	const op = "repository.GetHomeSportArticles()"

	var as []model.Article

	rows, err := r.db.Query(articleQueries.selectArticleWhereLimit, "sport", 3)
	if err != nil {
		r.l.Error().Msgf("%s: query select home sport articles > %s", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var a model.Article
		err = rows.Scan(&a.ArticleID, &a.Title, &a.PreviewText,
			&a.ArticleTime, &a.Tag, &a.DetailText, &a.Href, &a.Comments, &a.Category, &a.Image)
		if err != nil {
			r.l.Error().Msgf("%s: query scan home sport articles > %s", op, err)
			return nil, err
		}

		as = append(as, a)
	}

	return as, nil
}

// GetHomeVideoArticles return 3 video for video block
func (r *MySQLArticleRepository) GetHomeVideoArticles() ([]model.Article, error) {
	const op = "repository.GetHomeVideoArticles()"

	var as []model.Article

	rows, err := r.db.Query(articleQueries.selectVideoLimit, 3)
	if err != nil {
		r.l.Warn().Msgf("%s: query select videos for video block > %s", op, err)
	}

	defer rows.Close()

	for rows.Next() {
		var a model.Article
		err = rows.Scan(&a.ArticleID, &a.Title, &a.PreviewText,
			&a.ArticleTime, &a.Tag, &a.DetailText, &a.Href, &a.Comments, &a.Category, &a.Video)
		if err != nil {
			r.l.Error().Msgf("%s: query scan videos for video block > %s", op, err)
			return nil, err
		}

		as = append(as, a)
	}

	return as, nil
}

// GetHomeAllArticles return all articles except video
func (r *MySQLArticleRepository) GetHomeAllArticles() ([]model.Article, error) {
	const op = "repository.GetHomeAllArticles()"

	var as []model.Article

	rows, err := r.db.Query(articleQueries.selectAllArticle)
	if err != nil {
		r.l.Error().Msgf("%s: query select all articles > %s", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var a model.Article
		err = rows.Scan(&a.ArticleID, &a.Title, &a.PreviewText,
			&a.ArticleTime, &a.Tag, &a.DetailText, &a.Href, &a.Comments, &a.Category, &a.Image)
		if err != nil {
			r.l.Error().Msgf("%s: query scan all articles > %s", op, err)
			return nil, err
		}

		as = append(as, a)
	}

	return as, nil
}

func (r *MySQLArticleRepository) GetArticleByID(id int) (*model.Article, error) {
	return nil, nil
}
