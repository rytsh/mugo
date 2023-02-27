package funcs

import (
	"bytes"
	"fmt"

	"github.com/rytsh/mugo/internal/render/generic"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
	"github.com/tdewolff/minify/v2/json"
	"github.com/tdewolff/minify/v2/svg"
	"github.com/tdewolff/minify/v2/xml"
)

func init() {
	generic.CallReg.AddFunction("minify", generic.ReturnWithFn(Minify))
}

func Minify(mType string, data []byte) ([]byte, error) {
	buff := new(bytes.Buffer)

	switch mType {
	case "css":
		if err := css.Minify(minify.New(), buff, bytes.NewReader(data), nil); err != nil {
			return nil, err
		}
	case "js":
		if err := js.Minify(minify.New(), buff, bytes.NewReader(data), nil); err != nil {
			return nil, err
		}
	case "xml":
		if err := xml.Minify(minify.New(), buff, bytes.NewReader(data), nil); err != nil {
			return nil, err
		}
	case "html":
		if err := html.Minify(minify.New(), buff, bytes.NewReader(data), nil); err != nil {
			return nil, err
		}
	case "json":
		if err := json.Minify(minify.New(), buff, bytes.NewReader(data), nil); err != nil {
			return nil, err
		}
	case "svg":
		if err := svg.Minify(minify.New(), buff, bytes.NewReader(data), nil); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown minify type: %s", mType)
	}

	return buff.Bytes(), nil
}
