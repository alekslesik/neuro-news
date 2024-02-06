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
}

type ArticleModel interface {
	GetAllArticles() ([]Article, error)
	GetHomeCarouselArticles() ([]Article, error)
	GetHomeTrendingArticles() ([]Article, error)
	GetHomeNewsArticles() ([]Article, error)
	GetHomeSportArticles() ([]Article, error)
	GetHomeVideoArticles() ([]Article, error)
	GetHomePopularArticles() ([]Article, error)

	GetArticleByID(id int) (*Article, error)
}
