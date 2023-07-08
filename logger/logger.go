package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Setup() {
	logLevel := os.Getenv("GIT_LOG_LEVEL")

	var chosenLogLevel zerolog.Level

	switch logLevel {
	case "debug":
		chosenLogLevel = zerolog.DebugLevel
	case "warn":
		chosenLogLevel = zerolog.WarnLevel
	case "error":
		chosenLogLevel = zerolog.ErrorLevel
	case "fatal":
		chosenLogLevel = zerolog.FatalLevel
	case "panic":
		chosenLogLevel = zerolog.PanicLevel
	default:
		chosenLogLevel = zerolog.InfoLevel
	}

	log.Info().Msg("logger setup with level " + chosenLogLevel.String())
}
