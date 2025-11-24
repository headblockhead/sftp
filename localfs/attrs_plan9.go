package localfs

import (
	"io/fs"
	"syscall"

	sshfx "github.com/headblockhead/sftp/encoding/ssh/filexfer"
)

func fileStatFromInfoOs(fi fs.FileInfo, attrs *sshfx.Attributes) {
	if dir, ok := fi.Sys().(*syscall.Dir); ok {
		attrs.SetACModTime(dir.Atime, dir.Mtime)
	}
}
