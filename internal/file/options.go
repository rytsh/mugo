package file

import (
	"io/fs"
	"strconv"
)

type options struct {
	filePerm   *fs.FileMode
	folderPerm *fs.FileMode
}

type Option func(*options) error

func WithFilePerm(value string) Option {
	return func(options *options) error {
		if value == "" {
			return nil
		}
		perm, err := strconv.ParseUint(value, 8, 32)
		if err != nil {
			return err
		}
		mode := fs.FileMode(perm)
		options.filePerm = &mode
		return nil
	}
}

func WithFolderPerm(value string) Option {
	return func(options *options) error {
		if value == "" {
			return nil
		}
		perm, err := strconv.ParseUint(value, 8, 32)
		if err != nil {
			return err
		}
		mode := fs.FileMode(perm)
		options.folderPerm = &mode
		return nil
	}
}

func readOptions(opts []Option) (options, error) {
	var result options
	for _, opt := range opts {
		if err := opt(&result); err != nil {
			return result, err
		}
	}
	return result, nil
}
