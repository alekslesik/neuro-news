package service

import (
	"math"
	"net/http"
	"strconv"

	"github.com/alekslesik/neuro-news/internal/app/model"
	"github.com/alekslesik/neuro-news/internal/pkg/grabber"
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
	GetHomePaginationArticles(page string) ([]model.Article, error)
	GetArticleByURL(string) (model.Article, error)
	InsertArticleImage(*model.Image, *model.Article) error

	GrabNewArticle() (*model.Article, error)

	GetHomeTemplateData() (*template.TemplateData, error)
	GetHomePaginateData(string) (*template.TemplateData, error)
	GetArticleTemplateData(string) (*template.TemplateData, error)
	RenderTemplate(w http.ResponseWriter, r *http.Request, name string, td *template.TemplateData) error
}

type articleService struct {
	ar model.ArticleModel
	t  *template.Template
	l  *logger.Logger
	g  *grabber.Grabber
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

// GetHomeTemplateData return template data for home page
func (as *articleService) GetHomePaginateData(page string) (*template.TemplateData, error) {
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

	as.t.TemplateData.TemplateDataArticle.PaginationArticles, err = as.GetHomePaginationArticles(page)
	if err != nil {
		as.l.Error().Msgf("%s: get all home articles error > %s", op, err)
		return nil, err
	}

	as.t.TemplateData.TemplateDataPage, err = as.GetPaginationTemplateData(page)
	if err != nil {
		as.l.Error().Msgf("%s: get all home articles error > %s", op, err)
		return nil, err
	}

	return &as.t.TemplateData, nil
}

func (as *articleService) GetAllArticles() ([]model.Article, error) {
	return as.ar.SelectAllArticles()
}

func (as *articleService) GetHomeCarouselArticles() ([]model.Article, error) {
	return as.ar.SelectHomeCarouselArticles()
}

func (as *articleService) GetHomeTrendingArticlesTop() ([]model.Article, error) {
	return as.ar.SelectHomeTrendingArticlesTop()
}

func (as *articleService) GetHomeTrendingArticlesBottom() ([]model.Article, error) {
	return as.ar.SelectHomeTrendingArticlesBottom()
}

func (as *articleService) GetHomeNewsArticles() ([]model.Article, error) {
	return as.ar.SelectHomeNewsArticles()
}

func (as *articleService) GetHomeSportArticles() ([]model.Article, error) {
	return as.ar.SelectHomeSportArticles()
}

func (as *articleService) GetHomeVideoArticles() ([]model.Article, error) {
	return as.ar.SelectHomeVideoArticles()
}

func (as *articleService) GetHomeAllArticles() ([]model.Article, error) {
	return as.ar.SelectHomeAllArticles()
}

// GetHomePaginationArticles return []model.Article for pagination part
func (as *articleService) GetHomePaginationArticles(page string) ([]model.Article, error) {
	op := "service.GetHomePaginationArticles"
	limit := 15
	var offset int

	if page == "" {
		offset = 0
	} else {
		p, err := strconv.Atoi(page)
		if err != nil {
			as.l.Error().Msgf("%s: atoi convert pagination page error > %s", op, err)
			return nil, err
		}
		offset = (p - 1) * limit
	}

	return as.ar.SelectPaginationArticles(limit, offset)
}

// GetPaginationTemplateData return pagination page data
func (as *articleService) GetPaginationTemplateData(page string) (*template.TemplateDataPage, error) {
	const op = "service.GetPaginationTemplateData()"

	articlesOnPage := 15
	articlesCount, err := as.ar.CountArticles()
	if err != nil {
		as.l.Error().Msgf("%s: render template error > %s", op, err)
		return nil, err
	}

	totalPages := math.Ceil(float64(articlesCount) / float64(articlesOnPage))
	currentPage, err := strconv.Atoi(page)
	if err != nil {
		as.l.Error().Msgf("%s: convert page string pagination number to int > %s", op, err)
		return nil, err
	}

	data := &template.TemplateDataPage{
		TotalPaginationPages:  int(totalPages),
		CurrentPaginationPage: currentPage,
	}

	return data, nil
}

// GetArticleTemplateData return template data for article page by article URL
func (as *articleService) GetArticleTemplateData(url string) (*template.TemplateData, error) {
	const op = "service.GetArticleByURL()"
	var err error

	as.t.TemplateData.TemplateDataArticle.Article, err = as.GetArticleByURL(url)
	if err != nil {
		as.l.Error().Msgf("%s: get article template data error > %s", op, err)
		return nil, err
	}

	return &as.t.TemplateData, nil
}

func (as *articleService) GetArticleByURL(url string) (model.Article, error) {
	return as.ar.SelectArticleByURL(url)
}

// GrabNewArticle grab new article from news site
func (as *articleService) GrabNewArticle() (*model.Article, error) {
	const op = "service.GetNewArticle()"

	// get article model without image
	a, err := as.g.GrabArticle()
	if err != nil {
		as.l.Error().Msgf("%s: grab article error > %s", op, err)
		return nil, err
	}

	return a, err
}

// InsertArticleImage insert article to DB
func (as *articleService) InsertArticleImage(image *model.Image, article *model.Article) error {
	const op = "service.InsertArticleImage()"

	err := as.ar.InsertArticleImage(image, article)
	if err != nil {
		as.l.Error().Msgf("%s: insert article error > %s", op, err)
		return err
	}
	return nil
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
