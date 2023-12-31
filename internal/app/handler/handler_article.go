package handler

import (
    "encoding/json"
    "net/http"
)

// ArticleHandler handle requests related with articles
type ArticleHandler struct {
    AppHandler *AppHandler
}

// Create new instance of ArticleHandler.
func NewArticleHandler(appHandler *AppHandler) *ArticleHandler {
    return &ArticleHandler{
        AppHandler: appHandler,
    }
}

// GetAllArticles handle request to get all articles
func (ah *ArticleHandler) GetAllArticles(w http.ResponseWriter, r *http.Request) {
    articles, err := ah.AppHandler.articleService.GetAllArticles()
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    // Send articles as json response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(articles)
}
