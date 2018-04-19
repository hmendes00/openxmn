package servers

import (
	"time"

	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	organizations "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations"
	servers "github.com/XMNBlockchain/openxmn/engine/domain/data/types/servers"
	uuid "github.com/satori/go.uuid"
)

// ServerBuilder represents a server builder
type ServerBuilder interface {
	Create() ServerBuilder
	WithID(id *uuid.UUID) ServerBuilder
	WithMetaData(met metadata.MetaData) ServerBuilder
	WithOwner(owner organizations.Organization) ServerBuilder
	WithServer(serv servers.Server) ServerBuilder
	WithPrice(pr Price) ServerBuilder
	CreatedOn(crOn time.Time) ServerBuilder
	LastUpdatedOn(lstUpOn time.Time) ServerBuilder
	Now() (Server, error)
}
