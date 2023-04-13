package main

import (
	"flag"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	prod := flag.Bool("prod", false, "write log to file")
	debug := flag.Bool("debug", false, "enable debug mode")

	flag.Parse()

	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	logger := zerolog.New(os.Stdout)

	if *prod {
		if !*debug {
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}

		file, err := os.OpenFile("zerolog.log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			log.Panic().Err(err)
		}
		defer file.Close()

		logger = zerolog.New(file).With().Timestamp().Logger()
	}

	logger.Debug().Msg("Isso é só para debug")
	logger.Info().Msg("Olá Mundo!")
	logger.Warn().Msg("CUIDADO!!!")
}
