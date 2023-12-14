// cmd/neuro-news/main.go

package main

import (
	"log"

	"github.com/alekslesik/neuro-news/internal/app"
)

func main() {
	const op = "main()"

	app, err := app.New()
	if err != nil {
		log.Fatalf("%s > create app error: %v", op, err)

	}

	err = app.Run()
	if err != nil {
		log.Fatalf("%s > run app error: %v", op, err)
	}
}
