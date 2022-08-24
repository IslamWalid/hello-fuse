package fs

import (
	"context"
	"fmt"
	"time"

	"bazil.org/fuse"
)

type File struct {
	Type       fuse.DirentType
	Attributes fuse.Attr
    Content    []byte
}

func NewFile(content []byte) *File {
	fmt.Println("NewFile called")
	return &File{
		Type:    fuse.DT_File,
		Content: content,
		Attributes: fuse.Attr{
			Inode: 0,
			Atime: time.Now(),
			Mtime: time.Now(),
			Ctime: time.Now(),
			Mode:  0o666,
		},
	}
}

// GetEntryType returns the type of entry.
func (f *File) GetEntryType() fuse.DirentType {
	return f.Type
}

// Attr fills attr with the standard metadata for the node.
func (f *File) Attr(ctx context.Context, a *fuse.Attr) error {
	*a = f.Attributes
	return nil
}

func (f *File) ReadAll(ctx context.Context) ([]byte, error) {
	fmt.Println("ReadAll called")
	return f.Content, nil
}

func (f *File) Write(ctx context.Context, req *fuse.WriteRequest, resp *fuse.WriteResponse) error {
	fmt.Println("Write called: Size ", f.Attributes.Size)
	fmt.Println("Data to write: ", string(req.Data))
	f.Content = req.Data
	resp.Size = len(req.Data)
	f.Attributes.Size = uint64(resp.Size)
	return nil
}
