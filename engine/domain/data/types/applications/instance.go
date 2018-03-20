package application

import (
	types "github.com/XMNBlockchain/openxmn/engine/domain/data/types"
	trees "github.com/XMNBlockchain/openxmn/engine/domain/data/types/applications/trees"
	servers "github.com/XMNBlockchain/openxmn/engine/domain/data/types/servers"
)

// Instance represents an application instance running on a server
type Instance interface {
	GetMetaData() types.MetaData
	GetServer() servers.Server
	GetApplication() trees.Leaf
}
