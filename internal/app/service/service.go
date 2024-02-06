package service

import (
	"github.com/alekslesik/neuro-news/internal/pkg/template"

	"github.com/alekslesik/neuro-news/internal/app/repository"
	"github.com/alekslesik/neuro-news/pkg/logger"
)

type Services struct {
	r *repository.MySQLRepository
	l *logger.Logger
	t *template.Template
}

func New(r *repository.MySQLRepository, l *logger.Logger, t *template.Template) *Services {
	return &Services{
		r: r,
		t: t,
		l: l,
	}
}

func (s *Services) GetArticleService() ArticleService {
	return &articleService{
		ar: s.r.GetArticleRepository(),
		t:  s.t,
		l:  s.l,
	}
}

func (s *Services) GetUserService() UserService {
	return &userService{
		UserRepository: s.r.GetUserRepository(),
	}
}
