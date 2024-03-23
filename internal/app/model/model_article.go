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
	GetAllArticles() ([]Article, error)
	GetHomeCarouselArticles() ([]Article, error)
	GetHomeTrendingArticlesTop() ([]Article, error)
	GetHomeTrendingArticlesBottom() ([]Article, error)
	GetHomeNewsArticles() ([]Article, error)
	GetHomeSportArticles() ([]Article, error)
	GetHomeVideoArticles() ([]Article, error)
	GetHomeAllArticles() ([]Article, error)
	InsertArticleImage(*Image, *Article) error

	GetArticleByURL(string) (Article, error)
}
