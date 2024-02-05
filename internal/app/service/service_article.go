package service

import (
	"github.com/alekslesik/neuro-news/internal/app/model"
	"github.com/alekslesik/neuro-news/internal/pkg/template"
)

type ArticleService interface {
	GetAllArticles() ([]model.Article, error)
	GetHomeCarouselArticles() ([]model.Article, error)
	GetHomeTrendingArticles() ([]model.Article, error)
	GetHomeNewsArticles() ([]model.Article, error)
	GetHomeSportArticles() ([]model.Article, error)
	GetHomeVideoArticles() ([]model.Article, error)
	GetHomePopularArticles() ([]model.Article, error)

	GetArticleByID(id int) (*model.Article, error)
}

type articleService struct {
	ArticleRepository model.ArticleModel
	TemplateData template.TemplateData
}

// func NewArticleService(articleRepository model.ArticleModel) ArticleService {
// 	return &articleService{
// 		ArticleRepository: articleRepository,
// 	}
// }

func (as *articleService) GetAllArticles() ([]model.Article, error) {
	return as.ArticleRepository.GetAllArticles()
}

func (as *articleService) GetHomeCarouselArticles() ([]model.Article, error) {
	return as.ArticleRepository.GetHomeCarouselArticles()
}

func (as *articleService) GetHomeTrendingArticles() ([]model.Article, error) {
	return as.ArticleRepository.GetHomeTrendingArticles()
}

func (as *articleService) GetHomeNewsArticles() ([]model.Article, error) {
	return as.ArticleRepository.GetHomeNewsArticles()
}

func (as *articleService) GetHomeSportArticles() ([]model.Article, error) {
	return as.ArticleRepository.GetHomeSportArticles()
}

func (as *articleService) GetHomeVideoArticles() ([]model.Article, error) {
	return as.ArticleRepository.GetHomeVideoArticles()
}

func (as *articleService) GetHomePopularArticles() ([]model.Article, error) {
	return as.ArticleRepository.GetHomePopularArticles()
}

func (as *articleService) GetArticleByID(id int) (*model.Article, error) {
	return as.ArticleRepository.GetArticleByID(id)
}

func (as *articleService) GetHomeTemplateData() (*template.TemplateData, error) {
	var err error
	
	as.TemplateData.TemplateDataArticle.CarouselArticles, err = as.GetHomeCarouselArticles()
	if err != nil {
		return nil, err
	}

	return &as.TemplateData, nil
}
