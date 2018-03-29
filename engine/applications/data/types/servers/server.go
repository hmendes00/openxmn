package servers

import (
	"github.com/XMNBlockchain/openxmn/engine/applications/data/types/metadata"
	"github.com/XMNBlockchain/openxmn/engine/applications/data/types/wealth"
	servers "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/servers"
)

// Server represents a server attached to an organization
type Server struct {
	Met   *metadata.MetaData   `json:"metadata"`
	Org   *wealth.Organization `json:"organization"`
	Serv  *servers.Server      `json:"server"`
	Price *Price               `json:"price"`
}

// CreateServer creates a server instance
func CreateServer(met *metadata.MetaData, org *wealth.Organization, serv *servers.Server, price *Price) *Server {
	out := Server{
		Met:   met,
		Org:   org,
		Serv:  serv,
		Price: price,
	}

	return &out
}

// GetMetaData returns the metadata
func (serv *Server) GetMetaData() *metadata.MetaData {
	return serv.Met
}

// GetOrganization returns the organization
func (serv *Server) GetOrganization() *wealth.Organization {
	return serv.Org
}

// GetServer returns the server
func (serv *Server) GetServer() *servers.Server {
	return serv.Serv
}

// GetPrice returns the price
func (serv *Server) GetPrice() *Price {
	return serv.Price
}
