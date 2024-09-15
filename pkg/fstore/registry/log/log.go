package log

import (
	"github.com/rs/zerolog/log"
)

type Log struct{}

func (Log) Debug(v any) any {
	log.Debug().Msgf("%v\n", v)
	return v
}

func (Log) Info(v any) any {
	log.Info().Msgf("%v\n", v)
	return v
}

func (Log) Warn(v any) any {
	log.Warn().Msgf("%v\n", v)
	return v
}

func (Log) Error(v any) any {
	log.Error().Msgf("%v\n", v)
	return v
}

func (Log) Fatal(v any) any {
	log.Fatal().Msgf("%v\n", v)
	return v
}

func (Log) Panic(v any) any {
	log.Panic().Msgf("%v\n", v)
	return v
}
