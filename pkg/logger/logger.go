package logger

import (
	"errors"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Level string

const (
	DEVELOPMENT Level = "development"
	PRODUCTION  Level = "production"
)

var (
	once            sync.Once
	ErrCreateFile   error = errors.New("create log file error")
	ErrLevelMissing error = errors.New("logging level missing")
)

type Logger struct {
	zerolog.Logger
}

// New create new logger instance with level. File string must be like "./path/logname.log"
func New(l Level, file string) (*Logger, error) {
	SetGlobalLog()

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

// createLogFile create log file in specified filePath
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

// getProdLogger return logger with logging in file only
func getProdLogger(file *os.File) *Logger {
	zerolog.TimeFieldFormat = time.RFC1123
	z := zerolog.New(file).
		Level(zerolog.InfoLevel).
		With().
		Timestamp().
		Logger()

	return &Logger{z}
}

// getDevLogger return logger with logging in file and console
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

// SetGlobalLog set global logger
func SetGlobalLog() {
	// set up once
	once.Do(setOnceGlobalLog)
}

func setOnceGlobalLog() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		NoColor:    false,
		TimeFormat: "15:04:05",
		PartsOrder: []string{
			zerolog.TimestampFieldName,
			zerolog.LevelFieldName,
			zerolog.CallerFieldName,
			zerolog.MessageFieldName,
		},
	})
}
