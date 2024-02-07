package model

import (
	"time"
)

type Article struct {
	ArticleID   int
	Title       string
	PreviewText string
	Image       string
	ArticleTime time.Time
	Tag         string
	DetailText  string
	Href        string
	Comments    int
	Category    string
	None        string
}

type ArticleModel interface {
	GetAllArticles() ([]Article, error)
	GetHomeCarouselArticles() ([]Article, error)
	GetHomeTrendingArticlesTop() ([]Article, error)
	GetHomeTrendingArticlesBottom() ([]Article, error)
	GetHomeNewsArticles() ([]Article, error)
	GetHomeSportArticles() ([]Article, error)
	GetHomeVideoArticles() ([]Article, error)
	GetHomePopularArticles() ([]Article, error)

	GetArticleByID(id int) (*Article, error)
}
