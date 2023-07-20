package file

import "io/fs"

type Filter func(path fs.DirEntry) (bool, error)
