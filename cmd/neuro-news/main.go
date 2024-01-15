// cmd/neuro-news/main.go

package main

import (
	"github.com/alekslesik/neuro-news/internal/app"
	"github.com/alekslesik/neuro-news/pkg/logger"

	"github.com/rs/zerolog/log"
)

func main() {
	const op = "main()"

	// set global logger
	logger.SetGlobalLog()

	// initialization application
	app, err := app.New()
	if err != nil {
		log.Fatal().Msgf("%s > create app error: %v", op, err)
	}

	// run application
	err = app.Run()
	if err != nil {
		log.Fatal().Msgf("%s > run app error: %v", op, err)
	}
}
