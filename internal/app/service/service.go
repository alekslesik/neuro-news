package service

import (
	"github.com/alekslesik/neuro-news/internal/app/repository"
	"github.com/alekslesik/neuro-news/pkg/logger"
)

type Services struct {
	r *repository.MySQLRepository
	l *logger.Logger
}

func New(r *repository.MySQLRepository, l *logger.Logger) *Services {
	return &Services{
		r: r,
		l: l,
	}
}

func (s *Services) GetArticleService() ArticleService {
	return &articleService{
		ArticleRepository: s.r.GetArticleRepository(),
	}
}

func (s *Services) GetUserService() UserService {
	return &userService{
		UserRepository: s.r.GetUserRepository(),
	}
}
