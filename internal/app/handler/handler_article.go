package handler

import (
	"net/http"

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

func (a *ArticleHandler) GetHomeArticles(w http.ResponseWriter, r *http.Request) {
	const op = "GetHomeArticles()"


	// взять статьи для карусели (последние 4)
	a.AppHandler.articleService.GetAllArticles()

	// fmt.Fprint(w, op)

	a.AppHandler.templates.Render(w, r, "home.page.html", &template.TemplateData{
		CurrentYear: 2024,
	})

	// _, err := a.AppHandler.articleService.GetAllArticles()
	// if err != nil {
	// 	a.l.Err(err).Msgf("%s > get all articles error", op)
	// 	http.Error(w, "get all articles error", http.StatusInternalServerError)
	// 	return
	// }
}
