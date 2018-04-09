package databases

import (
	files "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands/files"
	organizations "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations"
	uuid "github.com/satori/go.uuid"
)

// Organization represents the organization database
type Organization struct {
	dirPath    string
	repository organizations.OrganizationRepository
}

// CreateOrganization creates an organization database
func CreateOrganization(dirPath string, repository organizations.OrganizationRepository) *Organization {
	out := Organization{
		dirPath:    dirPath,
		repository: repository,
	}

	return &out
}

// RetrieveByID retrieves an organization by ID
func (db *Organization) RetrieveByID(id *uuid.UUID) (organizations.Organization, error) {
	return nil, nil
}

// Insert inserts a new organizartion to the database
func (db *Organization) Insert(org organizations.Organization) (files.File, error) {
	return nil, nil
}

// Update uopdates an organization to the database
func (db *Organization) Update(old organizations.Organization, new organizations.Organization) (files.File, files.File, error) {
	return nil, nil, nil
}
