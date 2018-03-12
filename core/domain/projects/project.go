package projects

import (
	blockchains "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains"
	uuid "github.com/satori/go.uuid"
)

// Project represents a project
type Project interface {
	GetID() *uuid.UUID
	GetName() string
	GetDescription() string
	GetBlockchain() blockchains.Blockchain
	GetSettings() Settings
	HasDependencies() bool
	GetDependencies() Projects
}
