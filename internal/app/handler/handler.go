package handler

import (
	"github.com/alekslesik/neuro-news/internal/app/service"
	"github.com/alekslesik/neuro-news/internal/pkg/template"
	"github.com/alekslesik/neuro-news/pkg/logger"
)

type AppHandler struct {
	articleService service.ArticleService
	userService    service.UserService
	ArticleHandler ArticleHandler
	UserHandler    UserHandler
	templates      *template.Template
}

func New(services *service.Services, l *logger.Logger, templates *template.Template) *AppHandler {

	appHandler := &AppHandler{
		articleService: services.GetArticleService(),
		userService:    services.GetUserService(),
		templates:      templates,
	}

	articleHandler := NewArticleHandler(appHandler, l)
	userHandler := NewUserHandler(appHandler, l)

	appHandler.ArticleHandler = *articleHandler
	appHandler.UserHandler = *userHandler

	return appHandler
}
