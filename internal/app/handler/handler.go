package handler

import (
    "net/http"
    "github.com/alekslesik/neuro-news/internal/app/service"
)

type Handler struct {
    // Зависимости handler, например, сервисы
    UserService    *service.UserService
    ArticleService *service.ArticleService
    // ...
}

func NewHandler(userService *service.UserService, articleService *service.ArticleService) *Handler {
    return &Handler{
        UserService:    userService,
        ArticleService: articleService,
        // ...
    }
}

// Обработчики HTTP-запросов, например:
func (h *Handler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
    // Обработка запроса для получения пользователя
    // ...
}
