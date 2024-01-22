// cmd/neuro-news/main.go

package main

import (
	"context"

	"github.com/alekslesik/neuro-news/internal/app"
	"github.com/alekslesik/neuro-news/pkg/logger"

	"github.com/rs/zerolog/log"
)

func main() {
	const op = "main()"

	// set context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// set global logger
	logger.SetGlobalLog()

	// initialization application
	app, err := app.New(ctx, cancel)
	if err != nil {
		log.Fatal().Msgf("%s > create app error: %v", op, err)
	}

	// run application
	err = app.Run()
	if err != nil {
		log.Fatal().Msgf("%s > run app error: %v", op, err)
	}

	log.Warn().Msg("Application stopped")
}
