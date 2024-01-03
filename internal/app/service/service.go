package service

import (
	"github.com/alekslesik/neuro-news/internal/app/repository"
)

type Services struct {
    r *repository.MySQLRepository
}

func New(r *repository.MySQLRepository) *Services {
    return &Services{
        r: r,
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
