package funcs

import (
	"bytes"
	"fmt"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
	"github.com/tdewolff/minify/v2/json"
	"github.com/tdewolff/minify/v2/svg"
	"github.com/tdewolff/minify/v2/xml"

	"github.com/rytsh/mugo/pkg/fstore/registry"
)

func init() {
	registry.CallReg.AddFunction("minify", new(Minify).init)
}

type Minify struct {
	css  css.Minifier
	js   js.Minifier
	xml  xml.Minifier
	html html.Minifier
	json json.Minifier
	svg  svg.Minifier
}

func (m *Minify) init() any {
	m.css = css.Minifier{
		KeepCSS2:  true,
		Precision: 0,
	}
	m.js = js.Minifier{}
	m.xml = xml.Minifier{
		KeepWhitespace: false,
	}
	m.html = html.Minifier{
		KeepDocumentTags:        true,
		KeepConditionalComments: true,
		KeepEndTags:             true,
		KeepDefaultAttrVals:     true,
		KeepWhitespace:          false,
	}
	m.json = json.Minifier{}
	m.svg = svg.Minifier{
		KeepComments: false,
		Precision:    0,
	}

	return m.Minify
}

func (m *Minify) Minify(mType string, data []byte) ([]byte, error) {
	buff := new(bytes.Buffer)

	switch mType {
	case "css":
		if err := m.css.Minify(minify.New(), buff, bytes.NewReader(data), nil); err != nil {
			return nil, err
		}
	case "js":

		if err := m.js.Minify(minify.New(), buff, bytes.NewReader(data), nil); err != nil {
			return nil, err
		}
	case "xml":
		if err := m.xml.Minify(minify.New(), buff, bytes.NewReader(data), nil); err != nil {
			return nil, err
		}
	case "html":
		if err := m.html.Minify(minify.New(), buff, bytes.NewReader(data), nil); err != nil {
			return nil, err
		}
	case "json":
		if err := m.json.Minify(minify.New(), buff, bytes.NewReader(data), nil); err != nil {
			return nil, err
		}
	case "svg":
		if err := m.svg.Minify(minify.New(), buff, bytes.NewReader(data), nil); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown minify type: %s", mType)
	}

	return buff.Bytes(), nil
}
