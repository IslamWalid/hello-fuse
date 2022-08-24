package main

import (
	"fmt"
	"os"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
    myfs "hello-fuse/fs"
)

func main() {
	c, err := fuse.Mount("./mnt")
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
        os.Exit(1)
	}
	defer c.Close()

	err = fs.Serve(c, myfs.NewFS())
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
        os.Exit(1)
	}
}
