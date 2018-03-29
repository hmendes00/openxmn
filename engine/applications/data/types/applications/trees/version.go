package trees

import (
	"github.com/XMNBlockchain/openxmn/engine/applications/data/types/metadata"
	wealth "github.com/XMNBlockchain/openxmn/engine/applications/data/types/wealth"
	uuid "github.com/satori/go.uuid"
)

// Version represents a version of an application, in a branch
type Version struct {
	Met          *metadata.MetaData   `json:"metadata"`
	Prev         *Version             `json:"previous_version"`
	Org          *wealth.Organization `json:"organization"`
	Name         string               `json:"name"`
	BlockchainID *uuid.UUID           `json:"blockchain_id"`
	DockerFile   string               `json:"docker_file"`
}

// CreateVersion creates a version instance
func CreateVersion(met *metadata.MetaData, org *wealth.Organization, name string, blockchainID *uuid.UUID, dockerFile string) *Version {
	out := Version{
		Met:          met,
		Prev:         nil,
		Org:          org,
		Name:         name,
		BlockchainID: blockchainID,
		DockerFile:   dockerFile,
	}

	return &out
}

// CreateVersionFromPreviousVersion creates a version instance from a previous version
func CreateVersionFromPreviousVersion(met *metadata.MetaData, prev *Version, org *wealth.Organization, name string, blockchainID *uuid.UUID, dockerFile string) *Version {
	out := Version{
		Met:          met,
		Prev:         prev,
		Org:          org,
		Name:         name,
		BlockchainID: blockchainID,
		DockerFile:   dockerFile,
	}

	return &out
}

// GetMetaData returns the metadata
func (ver *Version) GetMetaData() *metadata.MetaData {
	return ver.Met
}

// GetPrevious returns the previous version
func (ver *Version) GetPrevious() *Version {
	return ver.Prev
}

// GetOrganization returns the organization
func (ver *Version) GetOrganization() *wealth.Organization {
	return ver.Org
}

// GetName returns the name
func (ver *Version) GetName() string {
	return ver.Name
}

// GetBlockchainID returns the blockchain ID
func (ver *Version) GetBlockchainID() *uuid.UUID {
	return ver.BlockchainID
}

// GetDockerFile returns the docker file
func (ver *Version) GetDockerFile() string {
	return ver.DockerFile
}
