package fs

import (
	"context"
	"fmt"
	"os"
	"syscall"
	"time"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

type Dir struct {
	Type       fuse.DirentType
	Attributes fuse.Attr
	Entries    map[string]any
}

func NewDir() *Dir {
	fmt.Println("NewDir called")
	return &Dir{
		Type: fuse.DT_Dir,
		Attributes: fuse.Attr{
			Inode: 0,
			Atime: time.Now(),
			Mtime: time.Now(),
			Ctime: time.Now(),
			Mode:  os.ModeDir | 0o777,
		},
		Entries: map[string]any{},
	}
}

// GetEntryType returns the type of entry.
func (d *Dir) GetEntryType() fuse.DirentType {
	return d.Type
}

// Mkdir handles mkdir requests.
func (d *Dir) Mkdir(ctx context.Context, req *fuse.MkdirRequest) (fs.Node, error) {
	fmt.Println("Mkdir called with name: ", req.Name)
	dir := NewDir()
	d.Entries[req.Name] = dir
	return dir, nil
}

// Attr fills attr with the standard metadata for the node.
func (d *Dir) Attr(ctx context.Context, a *fuse.Attr) error {
	*a = d.Attributes
	return nil
}

// LookUp serves look up requests comming from the kernel.
func (d *Dir) LookUp(ctx context.Context, name string) (fs.Node, error) {
	node, ok := d.Entries[name]
	if ok {
		return node.(fs.Node), nil
	}
	return nil, syscall.ENOENT
}

// ReadDirAll reads all entries
func (d *Dir) ReadDirAll(ctx context.Context) ([]fuse.Dirent, error) {
	fmt.Println("ReadDirAll called")
	var entries []fuse.Dirent

	for k, v := range d.Entries {
		var a fuse.Attr
		v.(fs.Node).Attr(ctx, &a)
		entries = append(entries, fuse.Dirent{
			Inode: a.Inode,
			Type:  v.(EntryGetter).GetEntryType(),
			Name:  k,
		})
	}
	return entries, nil
}

// Create creates a new directory entry in the receiver, which must be a directory.
func (d *Dir) Create(ctx context.Context, req *fuse.CreateRequest, resp *fuse.CreateResponse) (fs.Node, fs.Handle, error) {
	fmt.Println("Create called with filename: ", req.Name)
	f := NewFile(nil)
	fmt.Println("Create: Modified at", f.Attributes.Mtime.String())
	d.Entries[req.Name] = f
	return f, f, nil
}

func (d *Dir) Remove(ctx context.Context, req *fuse.RemoveRequest) error {
	delete(d.Entries, req.Name)
	return nil
}
