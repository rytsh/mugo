package file

import (
	"path/filepath"
	"reflect"
	"testing"
)

func TestLoadFormats(t *testing.T) {
	tests := map[string][]byte{
		"data.json": []byte(`{"name":"mugo"}`),
		"data.yaml": []byte("name: mugo\n"),
		"data.toml": []byte("name = 'mugo'\n"),
	}

	api := New()
	for name, content := range tests {
		t.Run(name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), name)
			if err := api.SetRaw(path, content); err != nil {
				t.Fatal(err)
			}

			var got map[string]any
			if err := api.Load(path, &got); err != nil {
				t.Fatal(err)
			}
			if got["name"] != "mugo" {
				t.Fatalf("name = %v, want mugo", got["name"])
			}
		})
	}
}

func TestRawRoundTrip(t *testing.T) {
	api := New()
	path := filepath.Join(t.TempDir(), "nested", "data.txt")
	want := []byte("mugo")

	if err := api.SetRaw(path, want, WithFolderPerm("0700"), WithFilePerm("0600")); err != nil {
		t.Fatal(err)
	}
	got, err := api.LoadRaw(path)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("LoadRaw() = %q, want %q", got, want)
	}
}
