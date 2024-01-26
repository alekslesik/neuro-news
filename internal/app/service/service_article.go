package service

import "github.com/alekslesik/neuro-news/internal/app/model"

type ArticleService interface {
	GetAllArticles() ([]model.Article, error)
	GetArticleByID(id int) (*model.Article, error)
}

type articleService struct {
	ArticleRepository model.ArticleModel
}

func NewArticleService(articleRepository model.ArticleModel) ArticleService {
	return &articleService{
		ArticleRepository: articleRepository,
	}
}

func (as *articleService) GetAllArticles() ([]model.Article, error) {
	return as.ArticleRepository.GetAllArticles()
}

func (as *articleService) GetArticleByID(id int) (*model.Article, error) {
	return as.ArticleRepository.GetArticleByID(id)
}

