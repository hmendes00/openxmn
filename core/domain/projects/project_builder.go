package projects

import (
	"github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains"
	uuid "github.com/satori/go.uuid"
)

// ProjectBuilder represents a project builder
type ProjectBuilder interface {
	Create() ProjectBuilder
	WithID(id *uuid.UUID) ProjectBuilder
	WithName(name string) ProjectBuilder
	WithDescription(description string) ProjectBuilder
	WithBlockchain(blkChain blockchains.Blockchain) ProjectBuilder
	WithDependencies(dependencies Projects) ProjectBuilder
	WithSettings(set Settings) ProjectBuilder
	Now() (Project, error)
}
