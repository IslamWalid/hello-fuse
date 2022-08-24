package fs

import (
	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

type EntryGetter interface {
	GetEntryType() fuse.DirentType
}

type FS struct{}

func NewFS() *FS {
	return &FS{}
}

// Root creates new directory to be the root directory of the filesystem.
func (f FS) Root() (fs.Node, error) {
	return NewDir(), nil
}
