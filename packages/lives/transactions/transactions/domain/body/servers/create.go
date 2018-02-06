package servers

import (
	servers "github.com/XMNBlockchain/core/packages/servers/domain"
	uuid "github.com/satori/go.uuid"
)

// Create represents a create server transaction
type Create interface {
	GetID() *uuid.UUID
	GetServer() servers.Server
	IsTransaction() bool
	IsLeader() bool
	IsDatabase() bool
}
