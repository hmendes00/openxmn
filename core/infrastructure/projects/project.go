package projects

import (
	projects "github.com/XMNBlockchain/exmachina-network/core/domain/projects"
	blockchains "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains"
	concrete_blockchains "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains"
	uuid "github.com/satori/go.uuid"
)

// Project represents a concrete project implementation
type Project struct {
	ID          *uuid.UUID                       `json:"id"`
	Name        string                           `json:"name"`
	Description string                           `json:"description"`
	Blockchain  *concrete_blockchains.Blockchain `json:"blockchain"`
	Sett        *Settings                        `json:"settings"`
	Dep         *Projects                        `json:"projects"`
}

func createProject(id *uuid.UUID, name string, description string, blockchain *concrete_blockchains.Blockchain, settings *Settings) projects.Project {
	out := Project{
		ID:          id,
		Name:        name,
		Description: description,
		Blockchain:  blockchain,
		Sett:        settings,
		Dep:         nil,
	}

	return &out
}

func createProjectWithDependencies(id *uuid.UUID, name string, description string, blockchain *concrete_blockchains.Blockchain, settings *Settings, dep *Projects) projects.Project {
	out := Project{
		ID:          id,
		Name:        name,
		Description: description,
		Blockchain:  blockchain,
		Sett:        settings,
		Dep:         dep,
	}

	return &out
}

// GetID returns the ID
func (proj *Project) GetID() *uuid.UUID {
	return proj.ID
}

// GetName returns the name
func (proj *Project) GetName() string {
	return proj.Name
}

// GetDescription returns the description
func (proj *Project) GetDescription() string {
	return proj.Description
}

// GetBlockchain returns the blockchain
func (proj *Project) GetBlockchain() blockchains.Blockchain {
	return proj.Blockchain
}

// GetSettings returns the settings
func (proj *Project) GetSettings() projects.Settings {
	return proj.Sett
}

// HasDependencies returns true if there is dependencies, false otherwise
func (proj *Project) HasDependencies() bool {
	return proj.Dep != nil
}

// GetDependencies returns the dependencies, if any
func (proj *Project) GetDependencies() projects.Projects {
	return proj.Dep
}
