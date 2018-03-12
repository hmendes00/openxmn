package projects

import (
	"errors"

	projects "github.com/XMNBlockchain/exmachina-network/core/domain/projects"
	blockchains "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains"
	concrete_blockchains "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains"
	uuid "github.com/satori/go.uuid"
)

type projectBuilder struct {
	id          *uuid.UUID
	name        string
	description string
	blockchain  blockchains.Blockchain
	settings    projects.Settings
	dep         projects.Projects
}

func createProjectBuilder() projects.ProjectBuilder {
	out := projectBuilder{
		id:          nil,
		name:        "",
		description: "",
		blockchain:  nil,
		settings:    nil,
		dep:         nil,
	}

	return &out
}

// Create initializes the project builder
func (build *projectBuilder) Create() projects.ProjectBuilder {
	build.id = nil
	build.name = ""
	build.description = ""
	build.blockchain = nil
	build.settings = nil
	build.dep = nil
	return build
}

// WithID adds an ID to the project builder
func (build *projectBuilder) WithID(id *uuid.UUID) projects.ProjectBuilder {
	build.id = id
	return build
}

// WithName adds a name to the project builder
func (build *projectBuilder) WithName(name string) projects.ProjectBuilder {
	build.name = name
	return build
}

// WithDescription adds a description to the project builder
func (build *projectBuilder) WithDescription(description string) projects.ProjectBuilder {
	build.description = description
	return build
}

// WithBlockchain adds a blockchain to the project builder
func (build *projectBuilder) WithBlockchain(blkChain blockchains.Blockchain) projects.ProjectBuilder {
	build.blockchain = blkChain
	return build
}

// WithDependencies adds Projects dependencies to the project builder
func (build *projectBuilder) WithDependencies(dependencies projects.Projects) projects.ProjectBuilder {
	build.dep = dependencies
	return build
}

// WithSettings adds Settings to the project builder
func (build *projectBuilder) WithSettings(set projects.Settings) projects.ProjectBuilder {
	build.settings = set
	return build
}

// Now builds a new Project instance
func (build *projectBuilder) Now() (projects.Project, error) {
	if build.id == nil {
		return nil, errors.New("the ID is mandatory in order to build a Project instance")
	}

	if build.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Project instance")
	}

	if build.description == "" {
		return nil, errors.New("the description is mandatory in order to build a Project instance")
	}

	if build.blockchain == nil {
		return nil, errors.New("the blockchain is mandatory in order to build a Project instance")
	}

	if build.settings == nil {
		return nil, errors.New("the settings is mandatory in order to build a Project instance")
	}

	if build.dep == nil {
		out := createProject(build.id, build.name, build.description, build.blockchain.(*concrete_blockchains.Blockchain), build.settings.(*Settings))
		return out, nil
	}

	out := createProjectWithDependencies(build.id, build.name, build.description, build.blockchain.(*concrete_blockchains.Blockchain), build.settings.(*Settings), build.dep.(*Projects))
	return out, nil
}
