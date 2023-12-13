package app

import (
	"database/sql"
	"flag"
	"log"

	"github.com/alekslesik/neuro-news/internal/app/model"
	"github.com/alekslesik/neuro-news/internal/pkg/middleware"
	"github.com/alekslesik/neuro-news/internal/pkg/router"
	"github.com/alekslesik/neuro-news/internal/pkg/template"
	"github.com/alekslesik/neuro-news/internal/pkg/db"

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
	model      *model.Model
	template   *template.Template
	dataBase   *sql.DB
	mailer     *mailer.Mailer
}

func New() (*Application, error) {
	const op = "app.New()"

	// config init
	//TODO add error returning
	config := config.New()

	flag.StringVar(&config.App.Env, "env", string(logger.DEVELOPMENT), "Environment (development|staging|production)")
	flag.IntVar(&config.App.Port, "port", 443, "API server port")
	config.MySQL.DSN = *flag.String("dsn", config.MySQL.DSN, "Name SQL data Source")
	flag.StringVar(&config.SMTP.Host, "smtp-host", "app.debugmail.io", "SMTP host")
	flag.IntVar(&config.SMTP.Port, "smtp-port", 25, "SMTP port")
	flag.StringVar(&config.SMTP.Username, "smtp-username", "d40e021c-f8d5-49af-a118-81f40f7b84b7", "SMTP username")
	flag.StringVar(&config.SMTP.Password, "smtp-password", "a8c960ed-d3ad-44e6-8461-37d40f15e569", "SMTP password")
	flag.StringVar(&config.SMTP.Sender, "smtp-sender", "alekslesik@gmail.com", "SMTP sender")
	flag.Parse()

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

	// Инициализация модели данных
	model := model.New(db)

	// Инициализация роутера
	appRouter := router.New()

	// Инициализация почтового сервиса
	// appMailer := mailer.New(appConfig.SMTPConfig)

	// Инициализация сессий
	// appSession := session.New()

	// Инициализация шаблонов
	// appTemplate := template.New()

	// Инициализация промежуточных обработчиков
	// appMiddleware := middleware.New()

	return &Application{
		config:     config,
		logger:     logger,
		router:     appRouter,
		// middleware: appMiddleware,
		// session:    appSession,
		model:      model,
		// template:   appTemplate,
		dataBase:   db,
		// mailer:     appMailer,
	}, nil
}

func (app *Application) Run() error {
	// Ваш код запуска приложения
	log.Println("Application is running...")
	return nil
}
