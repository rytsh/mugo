package file

import (
	"github.com/rytsh/mugo/fstore"
	internalfile "github.com/rytsh/mugo/internal/file"
)

func init() {
	fstore.AddStructWithOptions(func(o fstore.Option) (string, *File) {
		return "file", New(o.Trust)
	})
}

type File struct {
	trust bool
	api   *internalfile.API
}

func New(trust bool) *File {
	return &File{
		trust: trust,
		api:   internalfile.New(),
	}
}

// Deprecated: Use Write instead.
func (f *File) Save(fileName string, data []byte) (bool, error) {
	return f.Write(fileName, data)
}

func (f *File) Write(fileName string, data []byte) (bool, error) {
	if !f.trust {
		return false, fstore.ErrTrustRequired
	}

	if err := f.api.SetRaw(fileName, data); err != nil {
		return false, err
	}

	return true, nil
}

func (f *File) Read(fileName string) ([]byte, error) {
	return f.api.LoadRaw(fileName)
}
