package log

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

var l *logger = NewLogger()

const API = "logger test"

type logger struct {
	logger zerolog.Logger
	// see https://github.com/rs/zerolog#leveled-logging
}

func NewLogger() *logger {
	loadEnvFile()
	println("ENV " + os.Getenv("APP_ENV"))
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
	return &logger{
		logger: zlog,
	}
}

// PROBLEM: I need to duplicate loadEnvFile() from conf.load_env.go
// because conf uses log ... but conversely, log need conf cause it needs the env var
func loadEnvFile() {
	curDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err, "banking jobs", "App", "gw - conf - LoadEnv - os.Getwd()")
	}
	loadErr := godotenv.Load(curDir + "/.env")
	if loadErr != nil {
		log.Fatal(err, "banking jobs", "conf - LoadEnv", "godotenv.Load("+curDir+"/.env\")")
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
