package app

import (
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alekslesik/neuro-news/internal/app/handler"
	"github.com/alekslesik/neuro-news/internal/app/repository"
	"github.com/alekslesik/neuro-news/internal/app/service"
	"github.com/alekslesik/neuro-news/internal/pkg/router"

	"github.com/alekslesik/neuro-news/pkg/config"
	"github.com/alekslesik/neuro-news/pkg/logger"
)

type Application struct {
	c  *config.Config
	l  *logger.Logger
	r  *router.Router
	db *sql.DB
	// middleware *middleware.Middleware
	// session    *session.Session
	// template *template.Template
	// dataBase *sql.DB
	// mailer   *mailer.Mailer
}

func New() (*Application, error) {
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
		c:  config,
		l:  logger,
		r:  router,
		db: db,
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

	// addr := a.c.App.Host + ":" + strconv.Itoa(a.c.App.Port)

	// a.l.Info().Msgf("Application is running on %v", addr)

	var serverErr error

	// Get root certificate from system storage
	rootCAs, err := x509.SystemCertPool()
	if err != nil {
		a.l.Err(err).Msgf("%s > get root certificate", op)
		return err
	}

	// Set up server
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", a.c.App.Host, a.c.App.Port),
		Handler: a.r.Route(),
		TLSConfig: &tls.Config{
			MinVersion:         tls.VersionTLS12,
			RootCAs:            rootCAs,
			InsecureSkipVerify: false,
		},
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	switch srv.Addr {
	case "localhost:80", "localhost:8080", ":80", ":8080":
		go func() {
			if err := srv.ListenAndServe(); err != nil {
				serverErr = err
				a.l.Err(err).Msgf("%s > failed to start server", op)
			}
		}()
		a.l.Info().Msgf("server started on %s/", srv.Addr)
	case "localhost:443", "localhost:8443", ":443", ":8443":
		go func() {
			if err := srv.ListenAndServeTLS(a.c.TLS.CertPath, a.c.TLS.KeyPath); err != nil {
				serverErr = err
				a.l.Err(err).Msgf("%s > failed to start server", op)
			}
		}()
		a.l.Info().Msgf("server started on %s/", srv.Addr)
	default:
		a.l.Error().Msgf("%s: port not exists %s", op, srv.Addr)
	}

	<-done
	a.l.Info().Msg("server stopped")

	return serverErr
}

func (a *Application) closeDB() {
	const op = "app.Close()"

	if err := a.db.Close(); err != nil {
		a.l.Err(err).Msgf("%s > failed to close data base", op)
	}
}
