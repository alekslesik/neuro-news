package handler

import (
	"net/http"

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

// GetAllArticles handle request to get all articles
// func (a *ArticleHandler) GetAllArticles(c *gin.Context) {
// 	const op = "GetAllArticles()"
// 	articles, err := a.AppHandler.articleService.GetAllArticles()
// 	if err != nil {
// 		a.l.Err(err).Msgf("%s > get all articles error", op)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, articles)
// 	a.l.Info().Msgf("success GET %s response", c.Request.URL)
// }

func (a *ArticleHandler) GetAllArticles(w http.ResponseWriter, r *http.Request) {
	const op = "GetAllArticles()"

	_, err := a.AppHandler.articleService.GetAllArticles()
	if err != nil {
		a.l.Err(err).Msgf("%s > get all articles error", op)
		http.Error(w, "get all articles error", http.StatusInternalServerError)
		return
	}
}
