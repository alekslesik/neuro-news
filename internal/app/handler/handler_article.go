package handler

import (
    "encoding/json"
    "net/http"
)

// ArticleHandler обрабатывает запросы, связанные со статьями.
type ArticleHandler struct {
    AppHandler *AppHandler
}

// NewArticleHandler создает новый экземпляр ArticleHandler.
func NewArticleHandler(appHandler *AppHandler) *ArticleHandler {
    return &ArticleHandler{
        AppHandler: appHandler,
    }
}

// GetAllArticles обрабатывает запрос на получение всех статей.
func (ah *ArticleHandler) GetAllArticles(w http.ResponseWriter, r *http.Request) {
    articles, err := ah.AppHandler.ArticleService.GetAllArticles()
    if err != nil {
        // Обработка ошибки
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    // Отправляем статьи в виде JSON-ответа
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(articles)
}
