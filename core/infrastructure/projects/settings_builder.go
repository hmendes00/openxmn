package projects

import (
	"errors"

	projects "github.com/XMNBlockchain/exmachina-network/core/domain/projects"
	servers "github.com/XMNBlockchain/exmachina-network/core/domain/servers"
	concrete_servers "github.com/XMNBlockchain/exmachina-network/core/infrastructure/servers"
	uuid "github.com/satori/go.uuid"
)

type settingsBuilder struct {
	acceptedCurrencyID    *uuid.UUID
	mainBlockchainServers servers.Servers
}

func createSettingsBuilder() projects.SettingsBuilder {
	out := settingsBuilder{
		acceptedCurrencyID:    nil,
		mainBlockchainServers: nil,
	}
	return &out
}

// Create initializes the SettingsBuilder instance
func (build *settingsBuilder) Create() projects.SettingsBuilder {
	build.acceptedCurrencyID = nil
	build.mainBlockchainServers = nil
	return build
}

// WithAcceptedCurrencyID adds an accepted currency ID to the SettingsBuilder instance
func (build *settingsBuilder) WithAcceptedCurrencyID(id *uuid.UUID) projects.SettingsBuilder {
	build.acceptedCurrencyID = id
	return build
}

// WithMainBlockchainServers adds a main blockchain Servers to the SettingsBuilder instance
func (build *settingsBuilder) WithMainBlockchainServers(servs servers.Servers) projects.SettingsBuilder {
	build.mainBlockchainServers = servs
	return build
}

// Now builds a new Settings instance
func (build *settingsBuilder) Now() (projects.Settings, error) {
	if build.acceptedCurrencyID == nil {
		return nil, errors.New("the accepted currency ID is mandatory in order to build a Settings instance")
	}

	if build.mainBlockchainServers == nil {
		return nil, errors.New("the main blockchain servers are mandatory in order to build a Settings instance")
	}

	out := createSettings(build.acceptedCurrencyID, build.mainBlockchainServers.(*concrete_servers.Servers))
	return out, nil
}
