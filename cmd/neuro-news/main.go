// cmd/neuro-news/main.go

package main

import (
	"log"

	"github.com/alekslesik/neuro-news/internal/app"
)

func main() {
	const op = "main()"

	app := app.New()

	err := app.Run()
	if err != nil {
		log.Fatalf("%s > run app error: %v", op, err)
	}
}
