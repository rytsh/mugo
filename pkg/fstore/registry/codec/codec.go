package codec

import (
	"bufio"
	"bytes"
	"encoding/json"

	"github.com/BurntSushi/toml"
	"github.com/gomarkdown/markdown"
	"gopkg.in/yaml.v3"
)

type Codec struct{}

func (Codec) JsonDecode(v []byte) (any, error) {
	var data any
	err := json.Unmarshal(v, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (Codec) JsonEncode(v any, pretty bool) ([]byte, error) {
	if pretty {
		return json.MarshalIndent(v, "", "  ")
	}

	return json.Marshal(v)
}

func (Codec) YamlDecode(v []byte) (any, error) {
	var data any
	err := yaml.Unmarshal(v, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (Codec) YamlEncode(v any) ([]byte, error) {
	return yaml.Marshal(v)
}

func (Codec) TomlDecode(v []byte) (any, error) {
	var data any
	if _, err := toml.NewDecoder(bytes.NewReader(v)).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}

func (Codec) TomlEncode(v any) ([]byte, error) {
	var buf bytes.Buffer
	if err := toml.NewEncoder(&buf).Encode(v); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (Codec) Markdown(data []byte) []byte {
	return markdown.ToHTML(data, nil, nil)
}

func (Codec) ByteToString(b []byte) string {
	return string(b)
}

func (Codec) StringToByte(s string) []byte {
	return []byte(s)
}

func (Codec) IndentByte(i int, data []byte) []byte {
	dataR := bufio.NewReader(bytes.NewReader(data))
	var buf bytes.Buffer

	for {
		line, err := dataR.ReadBytes('\n')
		if err != nil {
			buf.Write(line)

			break
		}

		buf.Write(bytes.Repeat([]byte(" "), i))
		buf.Write(line)
	}

	return buf.Bytes()
}
