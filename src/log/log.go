package log

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

var l *logger

const API = "test api"

type logger struct {
	logger zerolog.Logger
	// see https://github.com/rs/zerolog#leveled-logging
}

func NewLogger() {
	// println("ENV " + os.Getenv("APP_ENV"))
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	var zlog zerolog.Logger
	if os.Getenv("APP_ENV") == "dev" {
		zlog = zerolog.New(
			zerolog.ConsoleWriter{
				Out:        os.Stderr,
				TimeFormat: time.RFC3339,
				FormatLevel: func(i interface{}) string {
					return strings.ToUpper(fmt.Sprintf("[%s]", i))
				},
				FormatMessage: func(i interface{}) string {
					return fmt.Sprintf("| %s |", i)
				},
			}).
			Level(zerolog.TraceLevel).
			With().
			Str("api", API).
			Timestamp().
			Logger()

	} else { // prod
		zlog = zerolog.New(os.Stdout).
			Level(zerolog.InfoLevel).
			With().
			Timestamp().
			Logger()
	}
	l = &logger{
		logger: zlog,
	}
}

func Info(path string, msg string) {
	l.logger.
		Info().
		Str("path", path).
		Msg(msg)
}

func Error(path string, err error) {
	l.logger.
		Error().
		Stack().
		Str("path", path).
		Msg(err.Error())
}

func Fatal(path string, err error) {
	l.logger.
		Fatal().
		Stack().
		Str("path", path).
		Msg(err.Error())
}
