package app

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alekslesik/neuro-news/internal/app/handler"
	"github.com/alekslesik/neuro-news/internal/app/repository"
	"github.com/alekslesik/neuro-news/internal/app/service"
	"github.com/alekslesik/neuro-news/internal/pkg/router"
	"github.com/alekslesik/neuro-news/internal/pkg/server"

	"github.com/alekslesik/neuro-news/pkg/config"
	"github.com/alekslesik/neuro-news/pkg/logger"
)

type Application struct {
	ctx context.Context
	ccl context.CancelFunc
	cfg *config.Config
	log *logger.Logger
	db  *sql.DB
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

	// repository init
	repositories := repository.New(db)

	// services init
	services := service.New(repositories, logger)

	// handlers init
	handler := handler.New(services, logger)

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

	// db close
	defer a.closeDB()
	// logfile close
	defer a.log.LogFile.Close()

	errChan := make(chan error)

	// Set signals handler
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	var err error
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
		a.log.Info().Msgf("server started on %s/", a.srv.Addr)
	default:
		a.log.Error().Msgf("%s: address or port are not exists %s", op, a.srv.Addr)
	}

	if err != nil && err != http.ErrServerClosed {
		errChan <- err
	}

	select {
	case <-a.ctx.Done():
		a.log.Warn().Msg("Context signal received, initiating shutdown")
		a.srv.Shutdown(a.ctx)
		time.Sleep(2 * time.Second)
	case err := <-errChan:
		a.log.Err(err).Msgf("%s > server failure", op)
		return err
	case <-signals:
		a.log.Warn().Msg("Signal received, initiating shutdown")
		a.srv.Shutdown(a.ctx)
		time.Sleep(2 * time.Second)
	}

	return nil
}

func (a *Application) closeDB() {
	const op = "app.Close()"

	if err := a.db.Close(); err != nil {
		a.log.Err(err).Msgf("%s > failed to close data base", op)
	}
}
