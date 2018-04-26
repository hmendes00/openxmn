package objects

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	servers "github.com/XMNBlockchain/openxmn/engine/domain/data/types/servers"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/metadata"
	concrete_servers "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/servers"
)

// Server represents a server that run projects
type Server struct {
	Met  *concrete_metadata.MetaData `json:"metadata"`
	Usr  *Holder                     `json:"holder"`
	Serv *concrete_servers.Server    `json:"server"`
	Proj *Project                    `json:"project"`
}

// CreateServer creates a new Server instance
func CreateServer(met metadata.MetaData, usr *Holder, serv servers.Server, proj *Project) *Server {
	out := Server{
		Met:  met.(*concrete_metadata.MetaData),
		Usr:  usr,
		Serv: serv.(*concrete_servers.Server),
		Proj: proj,
	}

	return &out
}
