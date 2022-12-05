package hugo

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/rytsh/mugo/internal/config"
	"github.com/spf13/afero"
	"github.com/spf13/cast"
)

type Namespace struct {
	workFs afero.Fs
}

func New(workDir string) *Namespace {
	return &Namespace{
		workFs: afero.NewOsFs(),
	}
}

// ReadDir lists the directory contents relative to the configured WorkingDir.
func (ns *Namespace) ReadDir(i any) ([]os.FileInfo, error) {
	path, err := cast.ToStringE(i)
	if err != nil {
		return nil, err
	}

	// check filename is relative
	if !filepath.IsAbs(path) {
		path = config.Checked.WorkDir + string(os.PathSeparator) + path
	}

	list, err := afero.ReadDir(ns.workFs, path)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory %q: %s", path, err)
	}

	return list, nil
}

// readFile reads the file named by filename in the given filesystem
// and returns the contents as a string.
func readFile(fs afero.Fs, filename string) (string, error) {
	filename = filepath.Clean(filename)
	if filename == "" || filename == "." || filename == string(os.PathSeparator) {
		return "", errors.New("invalid filename")
	}

	// check filename is relative
	if !filepath.IsAbs(filename) {
		filename = config.Checked.WorkDir + string(os.PathSeparator) + filename
	}

	b, err := afero.ReadFile(fs, filename)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// ReadFile reads the file named by filename relative to the configured WorkingDir.
// It returns the contents as a string.
// There is an upper size limit set at 1 megabytes.
func (ns *Namespace) ReadFile(i any) (string, error) {
	s, err := cast.ToStringE(i)
	if err != nil {
		return "", err
	}

	return readFile(ns.workFs, s)
}

// FileExists checks whether a file exists under the given path.
func (ns *Namespace) FileExists(i any) (bool, error) {
	path, err := cast.ToStringE(i)
	if err != nil {
		return false, err
	}

	if path == "" {
		return false, errors.New("fileExists needs a path to a file")
	}

	status, err := afero.Exists(ns.workFs, path)
	if err != nil {
		return false, err
	}

	return status, nil
}

// Stat returns the os.FileInfo structure describing file.
func (ns *Namespace) Stat(i any) (os.FileInfo, error) {
	path, err := cast.ToStringE(i)
	if err != nil {
		return nil, err
	}

	if path == "" {
		return nil, errors.New("fileStat needs a path to a file")
	}

	r, err := ns.workFs.Stat(path)
	if err != nil {
		return nil, err
	}

	return r, nil
}
