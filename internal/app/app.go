package app

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/alekslesik/neuro-news/internal/app/handler"
	// "github.com/alekslesik/neuro-news/internal/app/model"
	"github.com/alekslesik/neuro-news/internal/app/repository"
	"github.com/alekslesik/neuro-news/internal/app/service"
	"github.com/alekslesik/neuro-news/internal/pkg/db"
	"github.com/alekslesik/neuro-news/internal/pkg/flag"
	"github.com/alekslesik/neuro-news/internal/pkg/router"

	"github.com/alekslesik/neuro-news/pkg/config"
	"github.com/alekslesik/neuro-news/pkg/logger"
)

type Application struct {
	config *config.Config
	logger *logger.Logger
	router *router.Router
	db     *sql.DB
	// middleware *middleware.Middleware
	// session    *session.Session
	// template *template.Template
	// dataBase *sql.DB
	// mailer   *mailer.Mailer
}

func New() (*Application, error) {
	const op = "app.New()"

	// config init
	//TODO add error returning
	config, err := config.New()
	if err != nil {
		log.Fatalf("%s: config initialization error:  %s", op, err)
	}

	// flag init
	err = flag.Init(config)
	if err != nil {
		log.Fatalf("%s: flag initialization error:  %s", op, err)
	}

	// logger init
	logger, err := logger.New(logger.Level(config.Logger.LogLevel), config.Logger.LogFilePath)
	if err != nil {
		log.Fatal(err)
	}

	// data base init
	db, err := db.OpenDB(config.MySQL.DSN, config.MySQL.Driver)
	if err != nil {
		logger.Error().Msgf("%s: open db error: %v", op, err)
	}

	// repository init
	repositories := repository.New(db)

	// services init
	services := service.New(repositories)

	// handlers init
	handler := handler.New(services)

	// TODO Инициализация промежуточных обработчиков
	// appMiddleware := middleware.New()

	// router init
	router := router.New(handler)

	// Инициализация почтового сервиса
	// appMailer := mailer.New(appConfig.SMTPConfig)

	// Инициализация сессий
	// appSession := session.New()

	// Инициализация шаблонов
	// appTemplate := template.New()

	return &Application{
		config: config,
		logger: logger,
		router: router,
		db:     db,
		// middleware: appMiddleware,
		// session:    appSession,
		// model:      model,
		// template:   appTemplate,
		// dataBase: db,
		// mailer:     appMailer,
	}, nil
}

func (app *Application) Run() error {
	// const op = "app.Run()"

	defer app.closeDB()

	log.Println("Application is running...")

	err := http.ListenAndServe(":8080", app.router.Route())
	if err != nil {
		return err
	}

	return nil
}

func (app *Application) closeDB() {
	const op = "app.Close()"

	if err := app.db.Close(); err != nil {
		app.logger.Err(err).Msgf("%s > failed to close data base", op)
	}
}
