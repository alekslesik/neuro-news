package app

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/alekslesik/config"
	"github.com/alekslesik/neuro-news/internal/app/handler"
	"github.com/alekslesik/neuro-news/internal/app/repository"
	"github.com/alekslesik/neuro-news/internal/app/service"
	"github.com/alekslesik/neuro-news/internal/pkg/db"
	"github.com/alekslesik/neuro-news/internal/pkg/flag"
	"github.com/alekslesik/neuro-news/internal/pkg/router"

	"github.com/alekslesik/neuro-news/pkg/logger"
	"github.com/rs/zerolog/log"
)

type Application struct {
	c *config.Config
	l *logger.Logger
	r *router.Router
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
	config, err := config.New()
	if err != nil {
		log.Fatal().Msgf("%s: config initialization error:  %s", op, err)
	}

	// flag init
	err = flag.Init(config)
	if err != nil {
		log.Fatal().Msgf("%s: flag initialization error:  %s", op, err)
	}

	// logger init
	logger, err := logger.New(logger.Level(config.Logger.LogLevel), config.Logger.LogFilePath)
	if err != nil {
		log.Fatal().Msgf("%s: logger initialization error:  %s", op, err)
	}

	// db init
	db, err := db.OpenDB(config.MySQL.DSN, config.MySQL.Driver)
	if err != nil {
		logger.Error().Msgf("%s: db initialization error: %v", op, err)
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
		c: config,
		l: logger,
		r: router,
		db:     db,
		// middleware: appMiddleware,
		// session:    appSession,
		// model:      model,
		// template:   appTemplate,
		// dataBase: db,
		// mailer:     appMailer,
	}, nil
}

func (a *Application) Run() error {
	const op = "app.Run()"

	defer a.closeDB()
	defer a.l.LogFile.Close()

	a.l.Info().Msg("Application is running ...")

	addr := a.c.App.Host + ":" + strconv.Itoa(a.c.App.Port)

	err := http.ListenAndServe(addr, a.r.Route())
	if err != nil {
		a.l.Err(err).Msgf("%s > failed to listen and serve", op)
		return err
	}

	return nil
}

func (a *Application) closeDB() {
	const op = "app.Close()"

	if err := a.db.Close(); err != nil {
		a.l.Err(err).Msgf("%s > failed to close data base", op)
	}
}
