package projects

import (
	servers "github.com/XMNBlockchain/exmachina-network/core/domain/servers"
	uuid "github.com/satori/go.uuid"
)

// Settings represents a project settings
type Settings interface {
	GetAcceptedCurrencyID() *uuid.UUID
	GetMainBlockchainServers() servers.Servers
}
