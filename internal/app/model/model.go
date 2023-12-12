package model

import (
	"database/sql"
	"errors"
	"time"

	"github.com/alekslesik/neuro-news/internal/app/service"
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

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

func New(db *sql.DB) *Model {
	return &Model{
		Article: &service.ArticleModel{DB: db},
		User:    &service.UserModel{DB: db},
	}
}

type Model struct {
	Article interface {
		Insert(title, body, image string) error
		Get(id int) (*Article, error)
	}
	User interface {
		Insert(name, email, password string) error
		Authenticate(email, password string) (int, string, error)
		Get(id int) (*User, error)
	}
}
