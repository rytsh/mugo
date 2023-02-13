package render

import "github.com/rytsh/liz/utils/templatex"

func FuncMap(t *templatex.Template) map[string]interface{} {
	fMap := map[string]interface{}{
		"execTemplate": execTemplate(t),
		"rfc3339":      rfc3339,
		"saveFile":     saveFile,
		"logOutput":    logOutput,
		"nothing":      nothing,
	}

	return fMap
}
