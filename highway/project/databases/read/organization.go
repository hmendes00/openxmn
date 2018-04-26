package read

import (
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	objects "github.com/XMNBlockchain/openxmn/highway/project/objects"
	uuid "github.com/satori/go.uuid"
)

// Organization represents an organization read database
type Organization struct {
	org map[string]*objects.Organization
}

// CreateOrganization creates a new Organization instance
func CreateOrganization(org map[string]*objects.Organization) *Organization {
	out := Organization{
		org: org,
	}

	return &out
}

// RetrieveByID retrieves an organization by ID
func (db *Organization) RetrieveByID(id *uuid.UUID) (*objects.Organization, error) {
	return nil, nil
}

// CanUpdate verifies if a given user can update the given organization
func (db *Organization) CanUpdate(org *objects.Organization, user users.User) bool {
	return true
}

// CanDelete verifies if a given user can delete the given organization
func (db *Organization) CanDelete(org *objects.Organization, user users.User) bool {
	return true
}
