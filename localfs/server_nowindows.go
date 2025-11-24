//go:build !windows
// +build !windows

package localfs

import (
	"fmt"
	"io/fs"
)

func (h *ServerHandler) openfile(path string, flag int, mod fs.FileMode) (*File, error) {
	f, err := h.Root.OpenFile(path, flag, mod)
	if err != nil {
		return nil, err
	}

	return &File{
		filename: path,
		handle:   fmt.Sprint(h.handles.Add(1)),
		File:     f,
	}, nil
}
