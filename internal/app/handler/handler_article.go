package handler

import (
	"net/http"
	"strings"

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

func (a *ArticleHandler) GetHomeArticles(w http.ResponseWriter, r *http.Request) {
	const (
		op   = "GetHomeArticles()"
		page = "home.page.html"
	)

	td, err := a.AppHandler.articleService.GetHomeTemplateData()
	if err != nil {
		a.l.Error().Msgf("%s: GetHomeTemplateData error > %s", op, err)
	}

	err = a.AppHandler.articleService.RenderTemplate(w, r, page, td)
	if err != nil {
		a.l.Error().Msgf("%s: RenderTemplate error > %s", op, err)
	}
}

func (a *ArticleHandler) GetArticle(w http.ResponseWriter, r *http.Request) {
	const (
		op   = "GetArticle()"
		page = "article.page.html"
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

	err = a.AppHandler.articleService.RenderTemplate(w, r, page, td)
	if err != nil {
		a.l.Error().Msgf("%s: RenderTemplate error > %s", op, err)
	}
}
