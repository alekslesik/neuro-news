package logger

import (
	"errors"
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Level string

const (
	DEVELOPMENT Level = "development"
	PRODUCTION  Level = "production"
)

var ErrLevelMissing error = errors.New("logging level missing")
var ErrCreateFile error = errors.New("create log file error")

type Logger struct {
	zerolog.Logger
}

// Create new logger
func New(l Level, file string) (*Logger, error) {
	setGlobalLogger()

	switch l {
	case DEVELOPMENT:
		logFile, err := createLogFile(file)
		if err != nil {
			return nil, err
		}

		return getDevLogger(logFile), nil

	case PRODUCTION:
		logFile, err := createLogFile(file)
		if err != nil {
			return nil, err
		}

		return getProdLogger(logFile), nil
	}

	return nil, ErrLevelMissing
}

// Create log file in specified filePath
func createLogFile(logFilePath string) (*os.File, error) {
	// Get dir where log file must be
	logDir := filepath.Dir(logFilePath)

	// Check existing dir, and create if not exists
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err := os.MkdirAll(logDir, 0755); err != nil {
			return nil, err
		}
	}

	// Create or open log file for writing
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	return logFile, nil
}

// Logging within log file only
func getProdLogger(file *os.File) *Logger {
	zerolog.TimeFieldFormat = time.RFC1123
	z := zerolog.New(file).
		Level(zerolog.InfoLevel).
		With().
		Timestamp().
		Logger()

	return &Logger{z}
}

// Logging within log file and console
func getDevLogger(file *os.File) *Logger {
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC1123}
	multi := zerolog.MultiLevelWriter(consoleWriter, file)

	z := zerolog.New(multi).
		Level(zerolog.TraceLevel).
		With().
		Stack().
		Timestamp().
		Caller().
		Logger()

	return &Logger{z}
}

// Set global logger
func setGlobalLogger() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
}
