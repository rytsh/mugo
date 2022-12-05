// This package functions, scraped from Hugo's functions.
package hugo

import (
	"github.com/rytsh/mugo/internal/config"
)

func FuncMap() map[string]interface{} {
	ns := New(config.Checked.WorkDir)

	fMap := map[string]interface{}{
		"readFile":   ns.ReadFile,
		"readDir":    ns.ReadDir,
		"stat":       ns.Stat,
		"fileExists": ns.FileExists,
	}

	return fMap
}
