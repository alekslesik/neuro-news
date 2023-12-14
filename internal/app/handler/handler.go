package handler

import "github.com/alekslesik/neuro-news/internal/app/service"



type AppHandler struct {
    ArticleService service.ArticleService
    UserService    service.UserService
}

func NewAppHandler(articleService service.ArticleService, userService service.UserService) *AppHandler {
    return &AppHandler{
        ArticleService: articleService,
        UserService:    userService,
    }
}

// Теперь вы можете использовать AppHandler в ваших обработчиках, обращаясь к ArticleService и UserService
