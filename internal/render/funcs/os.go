package funcs

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/rytsh/mugo/internal/render/generic"
	"github.com/spf13/afero"
	"github.com/spf13/cast"
)

func init() {
	generic.CallReg.AddFunction("os", new(Os).init, "workDir")
}

type Os struct {
	workFs  afero.Fs
	workDir string
}

func (o *Os) init(workDir string) any {
	o.workFs = afero.NewOsFs()
	o.workDir = workDir

	return o
}

// ReadDir lists the directory contents relative to the configured WorkingDir.
func (o *Os) ReadDir(i any) ([]os.FileInfo, error) {
	path, err := cast.ToStringE(i)
	if err != nil {
		return nil, err
	}

	// check filename is relative
	if !filepath.IsAbs(path) {
		path = filepath.Join(o.workDir, path)
	}

	list, err := afero.ReadDir(o.workFs, path)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory %q: %s", path, err)
	}

	return list, nil
}

// readFile reads the file named by filename in the given filesystem
// and returns the contents as a string.
func (o *Os) readFile(filename string) (string, error) {
	filename = filepath.Clean(filename)
	if filename == "" || filename == "." || filename == string(os.PathSeparator) {
		return "", errors.New("invalid filename")
	}

	// check filename is relative
	if !filepath.IsAbs(filename) {
		filename = filepath.Join(o.workDir, filename)
	}

	b, err := afero.ReadFile(o.workFs, filename)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// ReadFile reads the file named by filename relative to the configured WorkingDir.
// It returns the contents as a string.
// There is an upper size limit set at 1 megabytes.
func (o *Os) ReadFile(i any) (string, error) {
	s, err := cast.ToStringE(i)
	if err != nil {
		return "", err
	}

	return o.readFile(s)
}

// FileExists checks whether a file exists under the given path.
func (o *Os) FileExists(i any) (bool, error) {
	path, err := cast.ToStringE(i)
	if err != nil {
		return false, err
	}

	if path == "" {
		return false, errors.New("fileExists needs a path to a file")
	}

	// check filename is relative
	if !filepath.IsAbs(path) {
		path = filepath.Join(o.workDir, path)
	}

	status, err := afero.Exists(o.workFs, path)
	if err != nil {
		return false, err
	}

	return status, nil
}

// Stat returns the os.FileInfo structure describing file.
func (o *Os) Stat(i any) (os.FileInfo, error) {
	path, err := cast.ToStringE(i)
	if err != nil {
		return nil, err
	}

	if path == "" {
		return nil, errors.New("fileStat needs a path to a file")
	}

	r, err := o.workFs.Stat(path)
	if err != nil {
		return nil, err
	}

	return r, nil
}
