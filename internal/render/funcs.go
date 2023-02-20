package render

import (
	"github.com/rytsh/liz/utils/templatex"
)

func FuncMap(trust bool) func(t *templatex.Template) map[string]interface{} {
	return func(t *templatex.Template) map[string]interface{} {
		// safe functions
		fMap := map[string]interface{}{
			"execTemplate": ExecTemplate(t),
			"rfc3339":      Rfc3339,
			"readFile":     ReadFile,
			"log":          Log,
			"nothing":      Nothing,
			"byteToString": ByteToString,
			"stringToByte": StringToByte,
			"md":           Md,
			"hold":         Hold,
			"getHold":      GetHold,
			"getData":      GetData,
			"indentByte":   IndentByte,
			"md5":          MD5,
			"sha1":         SHA1,
			"sha256":       SHA256,
			"fnv32a":       FNV32a,
			"hmac":         HMAC,
			"minify":       Minify,
		}

		if trust {
			// unsafe functions
			fMap["exec"] = Exec
			fMap["saveFile"] = SaveFile
		}

		return fMap
	}
}
