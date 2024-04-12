package main

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	zerolog.SetGlobalLevel(zerolog.TraceLevel)

	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	logger := zerolog.New(output).With().Timestamp().Caller().Logger()

	log.Trace().Msg("this is a debug message")
	logger.Debug().Msg("this is a debug message")
	logger.Info().Msg("this is an info message")
	logger.Warn().Msg("this is a warn message")
	logger.Error().Msg("this is an error message")
	// log.Fatal().Msg("this is a fatal message")
	// log.Panic().Msg("This is a panic message")
}
