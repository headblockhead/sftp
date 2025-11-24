package localfs

import (
	sftp "github.com/headblockhead/sftp"
)

var handler = &ServerHandler{}

// var _ sftp.HardlinkServerHandler = handler
var _ sftp.POSIXRenameServerHandler = handler
