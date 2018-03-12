package projects

import (
	servers "github.com/XMNBlockchain/exmachina-network/core/domain/servers"
	uuid "github.com/satori/go.uuid"
)

// SettingsBuilder represents a settings builder
type SettingsBuilder interface {
	Create() SettingsBuilder
	WithAcceptedCurrencyID(id *uuid.UUID) SettingsBuilder
	WithMainBlockchainServers(servs servers.Servers) SettingsBuilder
	Now() (Settings, error)
}
