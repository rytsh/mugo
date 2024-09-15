package file

import (
	"github.com/rytsh/liz/file"

	"github.com/rytsh/mugo/pkg/fstore/errors"
)

type File struct {
	trust bool
	api   *file.API
}

func New(trust bool) *File {
	return &File{
		trust: trust,
		api:   file.New(),
	}
}

// Deprecated: Use Write instead.
func (f *File) Save(fileName string, data []byte) (bool, error) {
	return f.Write(fileName, data)
}

func (f *File) Write(fileName string, data []byte) (bool, error) {
	if !f.trust {
		return false, errors.ErrTrustRequired
	}

	if err := f.api.SetRaw(fileName, data); err != nil {
		return false, err
	}

	return true, nil
}

func (f *File) Read(fileName string) ([]byte, error) {
	return f.api.LoadRaw(fileName)
}
