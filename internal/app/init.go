package app

import (
	"database/sql"

	"github.com/alekslesik/neuro-news/internal/pkg/db"
	"github.com/alekslesik/neuro-news/internal/pkg/flag"
	"github.com/alekslesik/neuro-news/pkg/config"
	"github.com/alekslesik/neuro-news/pkg/logger"
	"github.com/rs/zerolog/log"
)

// config init
func configInit() *config.Config {
	const op = "configInit()"
	c, err := config.New()
	if err != nil {
		log.Fatal().Msgf("%s: config initialization error:  %s", op, err)
	}

	return c
}

// flag init
func flagInit(c *config.Config)  {
	const op = "flagInit()"
	err := flag.Init(c)
	if err != nil {
		log.Fatal().Msgf("%s: flag initialization error:  %s", op, err)
	}
}

// logger init
func loggerInit(c *config.Config) *logger.Logger {
	const op = "loggerInit()"
	l, err := logger.New(logger.Level(c.Logger.LogLevel), c.Logger.LogFilePath)
	if err != nil {
		log.Fatal().Msgf("%s: logger initialization error:  %s", op, err)
	}

	return l
}

// db init
func dbInit(c *config.Config, l *logger.Logger) *sql.DB {
	const op = "dbInit()"
	// db init
	db, err := db.OpenDB(c.MySQL.DSN, c.MySQL.Driver)
	if err != nil {
		l.Error().Msgf("%s: db initialization error: %v", op, err)
	}

	return db
}