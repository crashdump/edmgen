// Package files implements file and directory filters.
//
// They are implemented as functions with the following signature:
// type Filter func(path fs.DirEntry) (bool, error)

package file
