package remote

import (
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
	servers "github.com/XMNBlockchain/openxmn/engine/domain/data/types/servers"
)

// FileBuilder represents a remote file builder
type FileBuilder interface {
	Create() FileBuilder
	WithServers(servs servers.Servers) FileBuilder
	WithFile(fil stored_files.File) FileBuilder
	Now() (File, error)
}
