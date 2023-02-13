package render

import (
	"github.com/rytsh/liz/loader/file"
)

var fileAPI = file.New()

func saveFile(fileName string, data string) (bool, error) {
	if err := fileAPI.SetRaw(fileName, []byte(data)); err != nil {
		return false, err
	}

	return true, nil
}
