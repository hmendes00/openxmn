package remote

import (
	fil "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
)

// File represents a stored remote file
type File interface {
	GetServers() []fil.File
	GetFile() fil.File
}
