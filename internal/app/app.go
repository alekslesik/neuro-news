package app

// 		log.Info().Msgf(")
// 		log.Warn().Msg("")
// 		log.Error().Msgf("%s:  > %s", op, err)
// 		log.Fatal().Msgf("%s:  > %s", op, err)

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	// "time"

	"github.com/alekslesik/neuro-news/internal/app/handler"
	"github.com/alekslesik/neuro-news/internal/app/repository"
	"github.com/alekslesik/neuro-news/internal/app/service"
	"github.com/alekslesik/neuro-news/internal/pkg/grabber"
	"github.com/alekslesik/neuro-news/internal/pkg/router"
	"github.com/alekslesik/neuro-news/internal/pkg/server"
	"github.com/alekslesik/neuro-news/internal/pkg/template"

	"github.com/alekslesik/neuro-news/pkg/config"
	"github.com/alekslesik/neuro-news/pkg/logger"
)

type Application struct {
	ctx context.Context
	ccl context.CancelFunc
	cfg *config.Config
	log *logger.Logger
	db  *sql.DB
	tp  *template.Template
	svs *service.Services
	grb *grabber.Grabber
	rtr *router.Router
	srv *server.Server
	// middleware *middleware.Middleware
	// session    *session.Session
	// template *template.Template
	// mailer   *mailer.Mailer
}

func New(context context.Context, cancel context.CancelFunc) (*Application, error) {
	const op = "app.New()"

	config := configInit()

	// flag init
	flagInit(config)

	// logger init
	logger := loggerInit(config)

	// db init
	db := dbInit(config, logger)

	// template init
	templates := templateInit(logger)

	// repository init
	repositories := repository.New(db, logger)

	// grabber init
	grabber := grabberInit(logger, config)

	// services init
	services := service.New(repositories, logger, templates, grabber)

	// handlers init
	handler := handler.New(services, logger, templates)

	// TODO Инициализация промежуточных обработчиков
	// appMiddleware := middleware.New()

	// router init
	router := routerInit(handler)

	server := serverInit(config, logger, router)

	// Инициализация почтового сервиса
	// appMailer := mailer.New(appConfig.SMTPConfig)

	// Инициализация сессий
	// appSession := session.New()

	// Инициализация шаблонов
	// appTemplate := template.New()

	return &Application{
		ctx: context,
		ccl: cancel,
		cfg: config,
		log: logger,
		rtr: router,
		db:  db,
		tp:  templates,
		svs: services,
		grb: grabber,
		srv: server,
		// middleware: appMiddleware,
		// session:    appSession,
		// model:      model,
		// template:   appTemplate,
		// mailer:     appMailer,
	}, nil
}

func (a *Application) Run() error {
	const op = "app.Run()"

	var err error
	errChan := make(chan error)

	go func() {
		for {
			err = a.svs.GetArticleService().GetNewArticle()
			if err != nil {
				a.log.Error().Msgf("%s: get new article error > %s", op, err)
			}

			time.Sleep(time.Minute * 10)
		}
	}()

	// db close
	defer a.closeDB()
	// logfile close
	defer a.log.LogFile.Close()

	// Set signals handler
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	switch a.srv.Addr {
	case "localhost:80", "localhost:8080", ":80", ":8080":
		go func() {
			err = a.srv.ListenAndServe()
		}()
		a.log.Info().Msgf("server started on %s/", a.srv.Addr)

	case "localhost:443", "localhost:8443", ":443", ":8443":
		go func() {
			err = a.srv.ListenAndServeTLS(a.cfg.TLS.CertPath, a.cfg.TLS.KeyPath)
		}()
		a.log.Info().Msgf("Server started on %s/", a.srv.Addr)

	default:
		a.log.Error().Msgf("%s: address or port are not exists > %s", op, a.srv.Addr)
	}

	if err != nil && err != http.ErrServerClosed {
		errChan <- err
	}

	select {
	case <-a.ctx.Done():
		a.log.Warn().Msg("Context signal received, initiating shutdown")
		a.srv.Shutdown(a.ctx)
		// TODO temprorary
		// time.Sleep(2 * time.Second)

	case err := <-errChan:
		a.log.Error().Msgf("%s: server failure > %s", op, err)
		return err

	case <-signals:
		a.log.Warn().Msg("Signal received, initiating shutdown")
		a.srv.Shutdown(a.ctx)
		// time.Sleep(2 * time.Second)
	}

	return nil
}

func (a *Application) closeDB() {
	const op = "app.Close()"

	if err := a.db.Close(); err != nil {
		a.log.Error().Msgf("%s: failed to close data base > %s", op, err)
	}
}
