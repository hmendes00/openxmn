package servers

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	organizations "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations"
	servers "github.com/XMNBlockchain/openxmn/engine/domain/data/types/servers"
)

// Server represents a server owned by an organization
type Server interface {
	GetMetaData() metadata.MetaData
	GetOwner() organizations.Organization
	GetServer() servers.Server
	GetPrice() Price
}
