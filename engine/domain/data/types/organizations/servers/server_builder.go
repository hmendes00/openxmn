package servers

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	organizations "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations"
	servers "github.com/XMNBlockchain/openxmn/engine/domain/data/types/servers"
)

// ServerBuilder represents a server builder
type ServerBuilder interface {
	Create() ServerBuilder
	WithMetaData(met metadata.MetaData) ServerBuilder
	WithOwner(owner organizations.Organization) ServerBuilder
	WithServer(serv servers.Server) ServerBuilder
	WithPrice(pr Price) ServerBuilder
	Now() (Server, error)
}
