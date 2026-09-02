package file

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

const (
	defaultFilePerm   fs.FileMode = 0o644
	defaultFolderPerm fs.FileMode = 0o755
)

type API struct {
	Codec map[string]Codec
}

func New() *API {
	jsonCodec := JSON{Indent: "  "}
	yamlCodec := YAML{}
	tomlCodec := TOML{}

	return &API{Codec: map[string]Codec{
		"JSON":  jsonCodec,
		".json": jsonCodec,
		"YAML":  yamlCodec,
		".yaml": yamlCodec,
		".yml":  yamlCodec,
		"TOML":  tomlCodec,
		".toml": tomlCodec,
	}}
}

func (a *API) OpenFile(path string, opts ...Option) (*os.File, error) {
	options, err := readOptions(opts)
	if err != nil {
		return nil, err
	}
	return openFileWrite(path, options)
}

func (a *API) LoadRaw(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", path, err)
	}
	return data, nil
}

func (a *API) Load(path string, dst any) error {
	codec, ok := a.Codec[filepath.Ext(path)]
	if !ok {
		return fmt.Errorf("failed to find codec for extension %s", filepath.Ext(path))
	}

	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", path, err)
	}
	defer f.Close()

	if err := codec.Decode(f, dst); err != nil {
		return fmt.Errorf("failed to unmarshal: %w", err)
	}
	return nil
}

func (a *API) LoadContent(data []byte, dst any, codec Codec) error {
	if codec == nil {
		return fmt.Errorf("failed codec is nil")
	}
	if err := codec.Decode(bytes.NewReader(data), dst); err != nil {
		return fmt.Errorf("failed to unmarshal: %w", err)
	}
	return nil
}

func (a *API) SetRaw(path string, data []byte, opts ...Option) error {
	return a.SetRawWithReader(path, bytes.NewReader(data), opts...)
}

func (a *API) SetRawWithReader(path string, data io.Reader, opts ...Option) error {
	options, err := readOptions(opts)
	if err != nil {
		return err
	}

	f, err := openFileWrite(path, options)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := io.Copy(f, data); err != nil {
		return fmt.Errorf("failed to write file %s: %w", path, err)
	}
	return nil
}

func openFileWrite(path string, options options) (*os.File, error) {
	folderPerm := defaultFolderPerm
	if options.folderPerm != nil {
		folderPerm = *options.folderPerm
	}
	if err := os.MkdirAll(filepath.Dir(path), folderPerm); err != nil {
		return nil, fmt.Errorf("failed to create folder %s: %w", filepath.Dir(path), err)
	}

	filePerm := defaultFilePerm
	if options.filePerm != nil {
		filePerm = *options.filePerm
	}
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, filePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", path, err)
	}
	return f, nil
}
