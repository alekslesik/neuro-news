package model

import (
	"errors"
	"time"
)

var (
	ErrNoRecord = errors.New("models: no suitable entry was found")
	//If a user tries to login with an incorrect email address or password.
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	//If a user tries to signup with an email address that's already in use.
	ErrDuplicateEmail = errors.New("models: duplicate email")
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

