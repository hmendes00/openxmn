package projects

import (
	projects "github.com/XMNBlockchain/exmachina-network/core/domain/projects"
	servers "github.com/XMNBlockchain/exmachina-network/core/domain/servers"
	concrete_servers "github.com/XMNBlockchain/exmachina-network/core/infrastructure/servers"
	uuid "github.com/satori/go.uuid"
)

// Settings represents a concrete settings implementation
type Settings struct {
	AcceptedCurrencyID    *uuid.UUID                `json:"accepted_currency_id"`
	MainBlockchainServers *concrete_servers.Servers `json:"main_blockchain_servers"`
}

func createSettings(acceptedCurrencyID *uuid.UUID, mainBlockchainServers *concrete_servers.Servers) projects.Settings {
	out := Settings{
		AcceptedCurrencyID:    acceptedCurrencyID,
		MainBlockchainServers: mainBlockchainServers,
	}
	return &out
}

// GetAcceptedCurrencyID returns the accepted currency ID
func (set *Settings) GetAcceptedCurrencyID() *uuid.UUID {
	return set.AcceptedCurrencyID
}

// GetMainBlockchainServers returns the servers that can connect to the main blockchain
func (set *Settings) GetMainBlockchainServers() servers.Servers {
	return set.MainBlockchainServers
}
