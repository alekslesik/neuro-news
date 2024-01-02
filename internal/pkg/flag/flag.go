package flag

import (
	"flag"

	"github.com/alekslesik/neuro-news/pkg/config"
	"github.com/alekslesik/neuro-news/pkg/logger"
)

// flags init
func Init(config *config.Config)  {
	flag.StringVar(&config.App.Env, "env", string(logger.DEVELOPMENT), "Environment (development|staging|production)")
	flag.IntVar(&config.App.Port, "port", 443, "API server port")
	flag.StringVar(&config.SMTP.Host, "smtp-host", "app.debugmail.io", "SMTP host")
	flag.IntVar(&config.SMTP.Port, "smtp-port", 25, "SMTP port")
	flag.StringVar(&config.SMTP.Username, "smtp-username", "d40e021c-f8d5-49af-a118-81f40f7b84b7", "SMTP username")
	flag.StringVar(&config.SMTP.Password, "smtp-password", "a8c960ed-d3ad-44e6-8461-37d40f15e569", "SMTP password")
	flag.StringVar(&config.SMTP.Sender, "smtp-sender", "alekslesik@gmail.com", "SMTP sender")
	flag.Parse()
}