package file

import (
	"io/fs"
	"path/filepath"
	"strings"
)

func IgnoreDirname(ignore []string) func(path fs.DirEntry) (bool, error) {
	return func(path fs.DirEntry) (bool, error) {
		for _, dir := range ignore {
			if path.IsDir() && strings.Contains(path.Name(), dir) {
				return false, filepath.SkipDir
			}
		}
		return true, nil
	}
}
