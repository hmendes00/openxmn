package remote

import (
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
	servers "github.com/XMNBlockchain/openxmn/engine/domain/data/types/servers"
)

// File represents a remote file
type File interface {
	GetServers() servers.Servers
	GetFile() stored_files.File
}
