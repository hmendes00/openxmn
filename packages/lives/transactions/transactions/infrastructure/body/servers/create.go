package servers

import (
	servers "github.com/XMNBlockchain/core/packages/servers/domain"
	concrete_servers "github.com/XMNBlockchain/core/packages/servers/infrastructure"
	servers_trs "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain/body/servers"
	uuid "github.com/satori/go.uuid"
)

// Create represents the concrete create user transaction
type Create struct {
	ID     *uuid.UUID               `json:"id"`
	Serv   *concrete_servers.Server `json:"server"`
	IsTrs  bool                     `json:"is_transaction"`
	IsLead bool                     `json:"is_leader"`
	IsDB   bool                     `json:"is_database"`
}

func createCreate(id *uuid.UUID, serv *concrete_servers.Server, isTrs bool, isLead bool, isDB bool) servers_trs.Create {
	out := Create{
		ID:     id,
		Serv:   serv,
		IsTrs:  isTrs,
		IsLead: isLead,
		IsDB:   isDB,
	}

	return &out
}

// GetID returns the ID
func (cr *Create) GetID() *uuid.UUID {
	return cr.ID
}

// GetServer returns the Server instance
func (cr *Create) GetServer() servers.Server {
	return cr.Serv
}

// IsTransaction returns true if the server is a Transaction server, false otherwise
func (cr *Create) IsTransaction() bool {
	return cr.IsTrs
}

// IsLeader returns true if the server is a Leader server, false otherwise
func (cr *Create) IsLeader() bool {
	return cr.IsLead
}

// IsDatabase returns true if the server is a Database server, false otherwise
func (cr *Create) IsDatabase() bool {
	return cr.IsDB
}
