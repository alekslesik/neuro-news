package app

import (
	"database/sql"

	"github.com/alekslesik/neuro-news/internal/pkg/middleware"
	"github.com/alekslesik/neuro-news/internal/pkg/router"
	"github.com/alekslesik/neuro-news/internal/pkg/sqlmodel"
	"github.com/alekslesik/neuro-news/internal/pkg/template"

	"github.com/alekslesik/neuro-news/pkg/config"
	"github.com/alekslesik/neuro-news/pkg/logger"
	"github.com/alekslesik/neuro-news/pkg/mailer"
	"github.com/alekslesik/neuro-news/pkg/session"
)

type Application struct {
	config     *config.Config
	logger     *logger.Logger
	router     *router.Router
	middleware *middleware.Middleware
	session    *session.Session
	model      *sqlmodel.Model
	template   *template.Template
	dataBase   *sql.DB
	mailer     *mailer.Mailer
}

// Create new instance of application
func New() *Application {
	return &Application{}
}

func (app *Application) New() error {
	return nil
}

func (app *Application) Run() error {
	return nil
}
