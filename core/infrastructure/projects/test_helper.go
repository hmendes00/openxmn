package projects

import (
	concrete_blockchains "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains"
	concrete_servers "github.com/XMNBlockchain/exmachina-network/core/infrastructure/servers"
	uuid "github.com/satori/go.uuid"
)

// CreateSettingsForTests creates a Settings for tests
func CreateSettingsForTests() *Settings {
	//variables:
	acceptedCurrencyID := uuid.NewV4()
	mainBlockchainServers := concrete_servers.CreateServersForTests()

	set := createSettings(&acceptedCurrencyID, mainBlockchainServers)
	return set.(*Settings)
}

// CreateProjectForTests creates a Settings for tests
func CreateProjectForTests() *Project {
	//variables:
	id := uuid.NewV4()
	name := "some name"
	description := "some description"
	blockchain := concrete_blockchains.CreateBlockchainForTests()
	settings := CreateSettingsForTests()

	proj := createProject(&id, name, description, blockchain, settings)
	return proj.(*Project)
}

// CreateProjectsForTests creates a Settings for tests
func CreateProjectsForTests() *Projects {
	//variables:
	lst := []*Project{
		CreateProjectForTests(),
		CreateProjectForTests(),
		CreateProjectForTests(),
		CreateProjectForTests(),
		CreateProjectForTests(),
	}

	projs := createProjects(lst)
	return projs.(*Projects)
}
