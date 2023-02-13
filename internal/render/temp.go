package render

import (
	"github.com/rs/zerolog/log"
	"github.com/rytsh/liz/utils/templatex"
)

func execTemplate(t *templatex.Template) func(name string, v any) (string, error) {
	return func(name string, v any) (string, error) {
		output, err := t.ExecuteBuffer(templatex.WithTemplate(name), templatex.WithData(v))
		return string(output), err
	}
}

func logOutput(v any) any {
	log.Info().Msgf("%v\n", v)

	return v
}

func nothing(v any) string {
	return ""
}
