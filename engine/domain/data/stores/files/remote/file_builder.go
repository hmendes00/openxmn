package remote

import (
	fil "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
)

// FileBuilder represents a remote file builder
type FileBuilder interface {
	Create() FileBuilder
	WithServers(servs []fil.File) FileBuilder
	WithFile(file fil.File) FileBuilder
	Now() (File, error)
}
