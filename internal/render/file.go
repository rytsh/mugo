package render

import (
	"github.com/rytsh/liz/loader/file"
)

var fileAPI = file.New()

func SaveFile(fileName string, data []byte) (bool, error) {
	if err := fileAPI.SetRaw(fileName, data); err != nil {
		return false, err
	}

	return true, nil
}

func ReadFile(fileName string) ([]byte, error) {
	return fileAPI.LoadRaw(fileName)
}
