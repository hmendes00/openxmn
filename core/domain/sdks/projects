package sdks

import (
	projects "github.com/XMNBlockchain/exmachina-network/core/domain/projects"
	servers "github.com/XMNBlockchain/exmachina-network/core/domain/servers"
	uuid "github.com/satori/go.uuid"
)

// Projects represents the projects SDK
type Projects interface {
	Retrieve(serv servers.Server, index int, amount int) (projects.Projects, error)
	RetrieveByID(serv servers.Server, id *uuid.UUID) (projects.Project, error)
}
