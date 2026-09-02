package fstore_test

import (
	"bytes"
	"fmt"
	"path/filepath"
	"testing"
	"text/template"

	_ "github.com/rytsh/mugo/fstore/registry"

	"github.com/rytsh/mugo/fstore"
	registryrandom "github.com/rytsh/mugo/fstore/registry/random"
)

func TestDocumentationExamples(t *testing.T) {
	registryrandom.SetDefaultRandomSeed(1)

	tests := []struct {
		name     string
		template string
		want     string
		trust    bool
	}{
		{name: "ungrouped", template: `{{ nothing 1 }}mugo`, want: "mugo"},
		{name: "sprig", template: `{{ upper "mugo" }}`, want: "MUGO"},
		{name: "cast", template: `{{ cast.ToInt "42" }}`, want: "42"},
		{name: "codec", template: "{{ $v := codec.JsonDecode (codec.StringToByte `{\"name\":\"mugo\"}`) }}{{ $v.name }}", want: "mugo"},
		{name: "crypto", template: `{{ crypto.SHA256 "mugo" }}`, want: "db3e53a360a0f4ebebff1d40ef5b3305ce1928e6e31b91ac4184971c4e2388bc"},
		{name: "exec", template: `{{ (exec.Exec "printf mugo").stdout | codec.ByteToString }}`, want: "mugo", trust: true},
		{name: "faker", template: `{{ len (faker.UUID.V4) }}`, want: "36"},
		{name: "html2", template: `{{ html2.EscapeString "<mugo>" }}`, want: "&lt;mugo&gt;"},
		{name: "humanize", template: `{{ humanize.Bytes 82854982 }}`, want: "83 MB"},
		{name: "log", template: `{{ index (log.Info "rendered" "name" "mugo") 1 }}`, want: "mugo"},
		{name: "map", template: `{{ map.Set "app/name" "mugo" | nothing }}{{ map.Get "app/name" nil }}`, want: "mugo"},
		{name: "math", template: `{{ math.Add "1.2" "2.3" }}`, want: "3.5"},
		{name: "minify", template: "{{ codec.ByteToString (minify \"json\" (codec.StringToByte `{ \"name\": \"mugo\" }`)) }}", want: `{"name":"mugo"}`},
		{name: "random", template: `{{ len (random.AlphaNum 8) }}`, want: "8"},
		{name: "time", template: `{{ time.Duration "2h" }}`, want: "2h0m0s"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			options := []fstore.OptionFunc{fstore.WithLog(fstore.Noop{})}
			if tt.trust {
				options = append(options, fstore.WithTrust(true))
			}
			if got := executeTemplate(t, tt.template, options...); got != tt.want {
				t.Fatalf("output = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestDocumentationFileExamples(t *testing.T) {
	path := filepath.Join(t.TempDir(), "mugo.txt")
	quotedPath := fmt.Sprintf("%q", path)
	content := fmt.Sprintf(`{{ file.Write %s (codec.StringToByte "mugo") | nothing }}{{ file.Read %s | codec.ByteToString }}`, quotedPath, quotedPath)

	if got := executeTemplate(t, content, fstore.WithTrust(true)); got != "mugo" {
		t.Fatalf("file output = %q, want mugo", got)
	}

	content = fmt.Sprintf(`{{ os.FileExists %s }}`, quotedPath)
	if got := executeTemplate(t, content); got != "true" {
		t.Fatalf("os output = %q, want true", got)
	}
}

func executeTemplate(t *testing.T, content string, opts ...fstore.OptionFunc) string {
	t.Helper()

	tpl, err := template.New("documentation").Funcs(fstore.FuncMap(opts...)).Parse(content)
	if err != nil {
		t.Fatal(err)
	}

	var output bytes.Buffer
	if err := tpl.Execute(&output, nil); err != nil {
		t.Fatal(err)
	}
	return output.String()
}
