//go:build !aix && !darwin && !dragonfly && !freebsd && !openbsd && !linux && !plan9
// +build !aix,!darwin,!dragonfly,!freebsd,!openbsd,!linux,!plan9

package statvfs

import (
	"runtime"

	sshfx "github.com/headblockhead/sftp/encoding/ssh/filexfer"
	"github.com/headblockhead/sftp/encoding/ssh/filexfer/openssh"
)

// StatVFS stubs the OpenSSH StatVFS with an sshfx.StatusOpUnsupported Status.
func StatVFS(name string) (*openssh.StatVFSExtendedReplyPacket, error) {
	return nil, &sshfx.StatusPacket{
		StatusCode:   sshfx.StatusOpUnsupported,
		ErrorMessage: "not supported by " + runtime.GOOS,
	}
}
