package render

import (
	"github.com/rs/zerolog/log"
	"github.com/rytsh/liz/utils/templatex"
)

func ExecTemplate(t *templatex.Template) func(name string, v any) (string, error) {
	return func(name string, v any) (string, error) {
		output, err := t.ExecuteBuffer(templatex.WithTemplate(name), templatex.WithData(v), templatex.WithParsed(true))
		return string(output), err
	}
}

func Log(level string, v any) any {
	switch level {
	case "debug":
		log.Debug().Msgf("%v\n", v)
	case "info":
		log.Info().Msgf("%v\n", v)
	case "warn":
		log.Warn().Msgf("%v\n", v)
	case "error":
		log.Error().Msgf("%v\n", v)
	case "fatal":
		log.Fatal().Msgf("%v\n", v)
	case "panic":
		log.Panic().Msgf("%v\n", v)
	default:
		log.Info().Msgf("%v\n", v)
	}

	return v
}

func Nothing(v any) string {
	return ""
}
