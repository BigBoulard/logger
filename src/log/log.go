package log

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

var l *logger = NewLogger()

type logger struct {
	logger zerolog.Logger
	// see https://github.com/rs/zerolog#leveled-logging
}

func NewLogger() *logger {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	var zlog zerolog.Logger
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
		Timestamp().
		Logger()

	return &logger{
		logger: zlog,
	}
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
