package app

import (
	"database/sql"

	"github.com/alekslesik/neuro-news/internal/app/handler"
	"github.com/alekslesik/neuro-news/internal/pkg/db"
	"github.com/alekslesik/neuro-news/internal/pkg/flag"
	"github.com/alekslesik/neuro-news/internal/pkg/grabber"
	"github.com/alekslesik/neuro-news/internal/pkg/router"
	"github.com/alekslesik/neuro-news/internal/pkg/server"
	"github.com/alekslesik/neuro-news/internal/pkg/template"

	"github.com/alekslesik/neuro-news/pkg/config"
	"github.com/alekslesik/neuro-news/pkg/logger"
	"github.com/rs/zerolog/log"
)

// config init
func configInit() *config.Config {
	const op = "configInit()"
	c, err := config.New()
	if err != nil {
		log.Fatal().Msgf("%s: config initialization error > %s", op, err)
	}

	return c
}

// flag init
func flagInit(c *config.Config) {
	const op = "flagInit()"
	err := flag.Init(c)
	if err != nil {
		log.Fatal().Msgf("%s: flag initialization error > %s", op, err)
	}
}

// logger init
func loggerInit(c *config.Config) *logger.Logger {
	const op = "loggerInit()"
	l, err := logger.New(logger.Level(c.Logger.LogLevel), c.Logger.LogFilePath)
	if err != nil {
		log.Fatal().Msgf("%s: logger initialization error > %s", op, err)
	}

	return l
}

// db init
func dbInit(c *config.Config, l *logger.Logger) *sql.DB {
	const op = "dbInit()"
	// db init
	db, err := db.OpenDB(c.MySQL.DSN, c.MySQL.Driver)
	if err != nil {
		l.Error().Msgf("%s: db initialization error > %v", op, err)
	}

	return db
}

// template init
func templateInit(l *logger.Logger) *template.Template {
	t := template.New(l)

	// appPath := os.Getenv("APP_PATH")

	// t.AddCache(appPath + "/website/content")
	t.AddCache("./website/content")

	return t
}

// grabber init
func grabberInit(l *logger.Logger, c *config.Config) *grabber.Grabber {
	home := "https://lenta.ru/"
	return grabber.New(l, c, home)
}

// router init
func routerInit(h *handler.AppHandler) *router.Router {
	return router.New(h)
}

// server init
func serverInit(c *config.Config, l *logger.Logger, r *router.Router) *server.Server {
	const op = "serverInit()"

	s, err := server.New(c, l, r)
	if err != nil {
		log.Fatal().Msgf("%s: server initialization error > %s", op, err)
	}

	return s
}
