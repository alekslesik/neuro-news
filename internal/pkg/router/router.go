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

	// engine := gin.New()

	// engine.GET("/", r.h.ArticleHandler.GetAllArticles)

	mux.HandleFunc("/", r.h.ArticleHandler.GetHomeArticles)
	mux.HandleFunc("/post", r.h.ArticleHandler.GetArticle)

	// static pages
	mux.HandleFunc("/about", r.h.CommonHandler.GetAboutPage)
	mux.HandleFunc("/contact", r.h.CommonHandler.GetContactPage)
	mux.HandleFunc("/advertisement", r.h.CommonHandler.GetAdvertisementPage)
	mux.HandleFunc("/privacy", r.h.CommonHandler.GetPrivacyPage)

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

	// file server for static files
	fileServer := http.FileServer(http.Dir("./website/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// for end-to-end testing
	// mux.Get("/ping", http.HandlerFunc(ping))

	// return standardMiddleware.Then(mux)

	// return engine

	return mux
}
