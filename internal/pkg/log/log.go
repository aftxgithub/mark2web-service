package web

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// Log levels
const (
	NO_LEVEL = "NO_LEVEL"

	TRACE = "TRACE"

	DEBUG = "DEBUG"

	INFO = "INFO"

	WARNING = "WARNING"

	ERROR = "ERROR"

	FATAL = "FATAL"
)

func New(level string) *log.Logger {
	return &log.Logger{
		Out:       os.Stderr,
		Formatter: new(log.TextFormatter),
		Level:     levelOf(level),
	}
}

func levelOf(level string) log.Level {
	switch level {
	default:
		return log.InfoLevel
	}
}
