package file

import (
	"io/fs"
	"strings"
)

func RequireExtensions(extensions []string) func(path fs.DirEntry) (bool, error) {
	return func(path fs.DirEntry) (bool, error) {
		for _, extension := range extensions {
			if strings.HasSuffix(path.Name(), extension) {
				return true, nil
			}
		}
		return false, nil
	}
}

func IgnoreExtensions(extensions []string) func(path fs.DirEntry) (bool, error) {
	return func(path fs.DirEntry) (bool, error) {
		for _, extension := range extensions {
			if !strings.HasSuffix(path.Name(), extension) {
				return true, nil
			}
		}
		return false, nil
	}
}
