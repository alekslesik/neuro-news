package model

import (
	"time"
)

type Article struct {
	ArticleID   int
	Title       string
	PreviewText string
	ImageID     int64
	ImagePath   string
	ArticleTime time.Time
	Tag         string
	DetailText  string
	Href        string
	Comments    int
	Category    string
	Kind        string
	VideoID     int64
	VideoPath   string
}

type ArticleModel interface {
	SelectAllArticles() ([]Article, error)
	SelectHomeCarouselArticles() ([]Article, error)
	SelectHomeTrendingArticlesTop() ([]Article, error)
	SelectHomeTrendingArticlesBottom() ([]Article, error)
	SelectHomeNewsArticles() ([]Article, error)
	SelectHomeSportArticles() ([]Article, error)
	SelectHomeVideoArticles() ([]Article, error)
	SelectPaginationArticles(limit, offset int) ([]Article, error)
	CountArticles() (int, error)

	InsertArticleImage(*Image, *Article) error
	SelectArticleByURL(string) (Article, error)
}
