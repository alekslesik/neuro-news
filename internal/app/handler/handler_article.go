package handler

import (
	"net/http"
	"strings"

	"github.com/alekslesik/neuro-news/internal/pkg/template"
	"github.com/alekslesik/neuro-news/pkg/logger"
)

// ArticleHandler handle requests related with articles
type ArticleHandler struct {
	AppHandler *AppHandler
	l          *logger.Logger
}

// NewArticleHandler create new instance of ArticleHandler.
func NewArticleHandler(appHandler *AppHandler, l *logger.Logger) *ArticleHandler {
	return &ArticleHandler{
		AppHandler: appHandler,
		l:          l,
	}
}

// GetHomeArticles GET handler for home page /?PAGEN_1
func (a *ArticleHandler) GetHomeArticles(w http.ResponseWriter, r *http.Request) {
	const (
		op       = "GetHomeArticles()"
		tmplFile = "home.page.html"
	)

	var (
		td  *template.TemplateData
		err error
	)

	page := r.URL.Query().Get("PAGEN_1")

	td, err = a.AppHandler.articleService.GetHomePaginateData(page)
	if err != nil {
		a.l.Error().Msgf("%s: GetHomePaginateData error > %s", op, err)
	}

	err = a.AppHandler.articleService.RenderTemplate(w, r, tmplFile, td)
	if err != nil {
		a.l.Error().Msgf("%s: RenderTemplate home page error > %s", op, err)
	}
}

func (a *ArticleHandler) GetArticle(w http.ResponseWriter, r *http.Request) {
	const (
		op   = "GetArticle()"
		tmplFile = "article.page.html"
	)

	// get article URL
	urlParts := strings.Split(r.URL.Path, "/")
	if len(urlParts) < 2 {
		http.NotFound(w, r)
		return
	}

	url := urlParts[len(urlParts)-1]

	td, err := a.AppHandler.articleService.GetArticleTemplateData(url)
	if err != nil {
		a.l.Error().Msgf("%s: GetArticle error > %s", op, err)
	}

	err = a.AppHandler.articleService.RenderTemplate(w, r, tmplFile, td)
	if err != nil {
		a.l.Error().Msgf("%s: RenderTemplate article page error > %s", op, err)
	}
}
