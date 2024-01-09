package handler

import "github.com/alekslesik/neuro-news/internal/app/service"

type AppHandler struct {
	articleService service.ArticleService
	userService    service.UserService
	ArticleHandler ArticleHandler
	UserHandler    UserHandler
}

func New(services *service.Services) *AppHandler {

	appHandler := &AppHandler{
		articleService: services.GetArticleService(),
		userService:    services.GetUserService(),
	}

	articleHandler := NewArticleHandler(appHandler)
	userHandler := NewUserHandler(appHandler)

	appHandler.ArticleHandler = *articleHandler
	appHandler.UserHandler = *userHandler

	return appHandler
}
