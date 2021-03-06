package read

import (
	"errors"
	"fmt"

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
	idAsString := id.String()
	if oneOrg, ok := db.org[idAsString]; ok {
		return oneOrg, nil
	}

	str := fmt.Sprintf("the organization (ID: %s) could not be found", idAsString)
	return nil, errors.New(str)
}

// CanUpdate verifies if a given user can update the given organization
func (db *Organization) CanUpdate(org *objects.Organization, user users.User) bool {
	return true
}

// CanDelete verifies if a given user can delete the given organization
func (db *Organization) CanDelete(org *objects.Organization, user users.User) bool {
	return true
}
