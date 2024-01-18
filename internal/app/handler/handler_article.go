package handler

import (

	"net/http"

	"github.com/gin-gonic/gin"
)

// ArticleHandler handle requests related with articles
type ArticleHandler struct {
	AppHandler *AppHandler
}

// NewArticleHandler create new instance of ArticleHandler.
func NewArticleHandler(appHandler *AppHandler) *ArticleHandler {
	return &ArticleHandler{
		AppHandler: appHandler,
	}
}

// GetAllArticles handle request to get all articles
func (ah *ArticleHandler) GetAllArticles(c *gin.Context) {
	articles, err := ah.AppHandler.articleService.GetAllArticles()
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, articles)
}
