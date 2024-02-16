package service

import (
	"net/http"

	"github.com/alekslesik/neuro-news/internal/app/model"
	"github.com/alekslesik/neuro-news/internal/pkg/template"
	"github.com/alekslesik/neuro-news/pkg/logger"
)

type ArticleService interface {
	GetAllArticles() ([]model.Article, error)
	GetHomeCarouselArticles() ([]model.Article, error)
	GetHomeTrendingArticlesTop() ([]model.Article, error)
	GetHomeTrendingArticlesBottom() ([]model.Article, error)
	GetHomeNewsArticles() ([]model.Article, error)
	GetHomeSportArticles() ([]model.Article, error)
	GetHomeVideoArticles() ([]model.Article, error)
	GetHomeAllArticles() ([]model.Article, error)
	GetArticleByID(id int) (*model.Article, error)

	GetHomeTemplateData() (*template.TemplateData, error)
	RenderTemplate(w http.ResponseWriter, r *http.Request, name string, td *template.TemplateData) error
}

type articleService struct {
	ar model.ArticleModel
	t  *template.Template
	l  *logger.Logger
}

// func NewArticleService(articleRepository model.ArticleModel) ArticleService {
// 	return &articleService{
// 		ArticleRepository: articleRepository,
// 	}
// }

func (as *articleService) GetAllArticles() ([]model.Article, error) {
	return as.ar.GetAllArticles()
}

func (as *articleService) GetHomeCarouselArticles() ([]model.Article, error) {
	return as.ar.GetHomeCarouselArticles()
}

func (as *articleService) GetHomeTrendingArticlesTop() ([]model.Article, error) {
	return as.ar.GetHomeTrendingArticlesTop()
}

func (as *articleService) GetHomeTrendingArticlesBottom() ([]model.Article, error) {
	return as.ar.GetHomeTrendingArticlesBottom()
}

func (as *articleService) GetHomeNewsArticles() ([]model.Article, error) {
	return as.ar.GetHomeNewsArticles()
}

func (as *articleService) GetHomeSportArticles() ([]model.Article, error) {
	return as.ar.GetHomeSportArticles()
}

func (as *articleService) GetHomeVideoArticles() ([]model.Article, error) {
	return as.ar.GetHomeVideoArticles()
}

func (as *articleService) GetHomeAllArticles() ([]model.Article, error) {
	return as.ar.GetHomeAllArticles()
}

func (as *articleService) GetArticleByID(id int) (*model.Article, error) {
	return as.ar.GetArticleByID(id)
}

// GetHomeTemplateData return template data for home page
func (as *articleService) GetHomeTemplateData() (*template.TemplateData, error) {
	const op = "service.GetHomeTemplateData()"
	var err error

	as.t.TemplateData.TemplateDataArticle.CarouselArticles, err = as.GetHomeCarouselArticles()
	if err != nil {
		as.l.Error().Msgf("%s: get home carouser articles error > %s", op, err)
		return nil, err
	}

	as.t.TemplateData.TemplateDataArticle.TrendingArticlesTop, err = as.GetHomeTrendingArticlesTop()
	if err != nil {
		as.l.Error().Msgf("%s: get home trending articles top error > %s", op, err)
		return nil, err
	}

	as.t.TemplateData.TemplateDataArticle.TrendingArticlesBottom, err = as.GetHomeTrendingArticlesBottom()
	if err != nil {
		as.l.Error().Msgf("%s: get home trending articles bottom error > %s", op, err)
		return nil, err
	}

	as.t.TemplateData.TemplateDataArticle.NewsArticles, err = as.GetHomeNewsArticles()
	if err != nil {
		as.l.Error().Msgf("%s: get home news articles error > %s", op, err)
		return nil, err
	}

	as.t.TemplateData.TemplateDataArticle.SportArticles, err = as.GetHomeSportArticles()
	if err != nil {
		as.l.Error().Msgf("%s: get home sport news articles error > %s", op, err)
		return nil, err
	}

	as.t.TemplateData.TemplateDataArticle.VideoArticles, err = as.GetHomeVideoArticles()
	if err != nil {
		as.l.Error().Msgf("%s: get home video articles error > %s", op, err)
		return nil, err
	}

	as.t.TemplateData.TemplateDataArticle.AllArticles, err = as.GetHomeAllArticles()
	if err != nil {
		as.l.Error().Msgf("%s: get all home articles error > %s", op, err)
		return nil, err
	}

	return &as.t.TemplateData, nil
}

// RenderTemplate render page with received data
func (as *articleService) RenderTemplate(w http.ResponseWriter, r *http.Request, n string, td *template.TemplateData) error {
	const op = "service.RenderTemplate()"

	err := as.t.Render(w, r, n, td)
	if err != nil {
		as.l.Error().Msgf("%s: render template error > %s", op, err)
		return err
	}

	return nil
}
