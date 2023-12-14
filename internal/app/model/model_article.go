package model

import (
	"time"
)

type Article struct {
	ID      int
	Title   string
	Body    string
	Image   string
	Created time.Time
}

type ArticleRepository interface {
    GetAllArticles() ([]Article, error)
    GetArticleByID(id int) (*Article, error)
}

