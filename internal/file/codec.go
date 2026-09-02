package file

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

type Codec interface {
	Encode(io.Writer, any) error
	Decode(io.Reader, any) error
}

type JSON struct {
	Indent string
}

func (j JSON) Encode(w io.Writer, v any) error {
	encoder := json.NewEncoder(w)
	if j.Indent != "" {
		encoder.SetIndent("", j.Indent)
	}
	if err := encoder.Encode(v); err != nil {
		return fmt.Errorf("json encode: %w", err)
	}
	return nil
}

func (JSON) Decode(r io.Reader, v any) error {
	if err := json.NewDecoder(r).Decode(v); err != nil {
		return fmt.Errorf("json decode: %w", err)
	}
	return nil
}

type YAML struct{}

func (YAML) Encode(w io.Writer, v any) error {
	if err := yaml.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("failed to encode YAML: %w", err)
	}
	return nil
}

func (YAML) Decode(r io.Reader, v any) error {
	if err := yaml.NewDecoder(r).Decode(v); err != nil {
		return fmt.Errorf("failed to decode YAML: %w", err)
	}
	return nil
}

type TOML struct{}

func (TOML) Encode(w io.Writer, v any) error {
	if err := toml.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("toml encode: %w", err)
	}
	return nil
}

func (TOML) Decode(r io.Reader, v any) error {
	if _, err := toml.NewDecoder(r).Decode(v); err != nil {
		return fmt.Errorf("toml decode: %w", err)
	}
	return nil
}
