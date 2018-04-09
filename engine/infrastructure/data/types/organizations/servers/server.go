package servers

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	organizations "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations"
	organizations_server "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations/servers"
	servers "github.com/XMNBlockchain/openxmn/engine/domain/data/types/servers"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/metadata"
	concrete_organizations "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/organizations"
	concrete_servers "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/servers"
)

// Server represents a concrete server implementation
type Server struct {
	Met   *concrete_metadata.MetaData          `json:"metadata"`
	Owner *concrete_organizations.Organization `json:"owner"`
	Serv  *concrete_servers.Server             `json:"server"`
	Pr    *Price                               `json:"price"`
}

func createServer(met *concrete_metadata.MetaData, owner *concrete_organizations.Organization, serv *concrete_servers.Server, pr *Price) organizations_server.Server {
	out := Server{
		Met:   met,
		Owner: owner,
		Serv:  serv,
		Pr:    pr,
	}

	return &out
}

// GetMetaData returns the metadata
func (serv *Server) GetMetaData() metadata.MetaData {
	return serv.Met
}

// GetOwner returns the owner
func (serv *Server) GetOwner() organizations.Organization {
	return serv.Owner
}

// GetServer returns the server
func (serv *Server) GetServer() servers.Server {
	return serv.Serv
}

// GetPrice returns the price
func (serv *Server) GetPrice() organizations_server.Price {
	return serv.Pr
}
