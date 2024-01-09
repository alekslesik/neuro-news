package router

import (
	"net/http"

	"github.com/alekslesik/neuro-news/internal/app/handler"
)

type Router struct {
	h *handler.AppHandler
}

func New(handler *handler.AppHandler) *Router {
	return &Router{
		h: handler,
	}
}

func (r *Router) Route() http.Handler {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	mux := http.NewServeMux()

	mux.HandleFunc("/", r.h.ArticleHandler.GetAllArticles)
	// router.POST("/somePost", posting)
	// router.PUT("/somePut", putting)
	// router.DELETE("/someDelete", deleting)
	// router.PATCH("/somePatch", patching)
	// router.HEAD("/someHead", head)
	// router.OPTIONS("/someOptions", options)

	// Пример использования функции Dir
	// В данном случае, мы указываем путь к директории "static", которая будет доступна на сервере
	// Если listDirectory установлено в true, то файлы директории будут отображаться, иначе - нет
	// engine.StaticFS("/static", http.Dir("./"))

	// for end-to-end testing
	// mux.Get("/ping", http.HandlerFunc(ping))

	// return standardMiddleware.Then(mux)

	return mux
}
