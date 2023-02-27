package funcs

import (
	"github.com/rytsh/liz/loader/file"
	"github.com/rytsh/mugo/internal/render/generic"
)

func init() {
	generic.CallReg.AddFunction("file", new(File).init, "trust")
}

type File struct {
	trust bool
	api   *file.API
}

func (f *File) init(trust bool) *File {
	f.trust = trust
	f.api = file.New()

	return f
}

func (f *File) Save(fileName string, data []byte) (bool, error) {
	if !f.trust {
		return false, generic.ErrTrustRequired
	}

	if err := f.api.SetRaw(fileName, data); err != nil {
		return false, err
	}

	return true, nil
}

func (f *File) Read(fileName string) ([]byte, error) {
	return f.api.LoadRaw(fileName)
}
