package render

import (
	"bufio"
	"bytes"

	"github.com/gomarkdown/markdown"
)

func ByteToString(b []byte) string {
	return string(b)
}

func StringToByte(s string) []byte {
	return []byte(s)
}

func Md(data []byte) []byte {
	return markdown.ToHTML(data, nil, nil)
}

func IndentByte(i int, data []byte) []byte {
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
