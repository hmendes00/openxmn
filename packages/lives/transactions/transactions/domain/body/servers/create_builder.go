package servers

import (
	servers "github.com/XMNBlockchain/core/packages/servers/domain"
	uuid "github.com/satori/go.uuid"
)

// CreateBuilder represents the builder of a create server transaction
type CreateBuilder interface {
	Create() CreateBuilder
	WithID(id *uuid.UUID) CreateBuilder
	WithServer(serv servers.Server) CreateBuilder
	IsTransaction() CreateBuilder
	IsLeader() CreateBuilder
	IsDatabase() CreateBuilder
	Now() (Create, error)
}
