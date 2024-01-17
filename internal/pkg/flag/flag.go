package flag

import (
	"errors"
	"flag"
	"os"

	"github.com/alekslesik/config"
)

var (
	ErrWrongEnv  = errors.New("wrong value of 'env' flag, use development | staging | production")
	ErrWrongPort = errors.New("wrong value of 'port' flag, use range from 1 to 65535")
)

// Init initialize flags using config file
func Init(config *config.Config) error {
	// create flagset
	flagSet := flag.NewFlagSet("flag", flag.ContinueOnError)

	// define flags in flagset
	flagSet.StringVar(&config.App.Env, "env", config.App.Env, "Environment (development|staging|production)")
	flagSet.IntVar(&config.App.Port, "port", config.App.Port, "API server port")

	// take arguments transferred to application using os.Args slice
	args := os.Args[1:]
	err := flagSet.Parse(args)
	if err != nil {
		return err
	}

	// env validation
	validEnvs := map[string]bool{"development": true, "staging": true, "production": true}
	if _, valid := validEnvs[config.App.Env]; !valid {
		return ErrWrongEnv
	}

	// port validation
	if !(config.App.Port > 0 && config.App.Port < 65536) {
		return ErrWrongPort
	}

	return nil
}
