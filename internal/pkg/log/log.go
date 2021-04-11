package log

import (
	"os"
	"strings"

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
		Out:   os.Stderr,
		Level: levelOf(level),
	}
}

func levelOf(level string) log.Level {
	switch strings.ToUpper(level) {
	case TRACE:
		return log.TraceLevel
	case DEBUG:
		return log.DebugLevel
	case INFO:
		return log.InfoLevel
	case WARNING:
		return log.WarnLevel
	case ERROR:
		return log.ErrorLevel
	case FATAL:
		return log.FatalLevel
	default:
		return log.InfoLevel
	}
}
