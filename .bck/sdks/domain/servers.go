package domain

import (
	servers "github.com/XMNBlockchain/core/packages/servers/domain"
	uuid "github.com/satori/go.uuid"
)

// Servers represents the servers SDK
type Servers interface {
	RetrieveFirstByBlockchainID(blockchainID *uuid.UUID, typ string) (servers.Server, error)
	RetrieveAllByBlockchainID(blockchainID *uuid.UUID, typ string, index int, amount int) ([]servers.Server, error)
}
