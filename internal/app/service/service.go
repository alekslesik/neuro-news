package service

import (
	"github.com/alekslesik/neuro-news/internal/app/repository"
)

type Services struct {}

func New(r *repository.MySQLRepository) *Services {
    return &Services{}
}

func (s *Services) GetArticleService() ArticleService {
    return &articleService{}
}

func (s *Services) GetUserService() UserService {
    return &userService{}
}
