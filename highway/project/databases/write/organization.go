package write

import (
	"errors"
	"fmt"

	objects "github.com/XMNBlockchain/openxmn/highway/project/objects"
)

// Organization represents an organization write database
type Organization struct {
	organizations map[string]*objects.Organization
}

// CreateOrganization creates a new Organization instance
func CreateOrganization(organizations map[string]*objects.Organization) *Organization {
	out := Organization{
		organizations: organizations,
	}

	return &out
}

// Insert inserts a new organization
func (db *Organization) Insert(org *objects.Organization) {
	db.organizations[org.Met.GetID().String()] = org
}

// Update updates an existing organization
func (db *Organization) Update(original *objects.Organization, new *objects.Organization) error {
	delErr := db.Delete(original)
	if delErr != nil {
		return delErr
	}

	db.Insert(new)
	return nil
}

// Delete deletes an existing organization
func (db *Organization) Delete(org *objects.Organization) error {
	idAsString := org.Met.GetID().String()
	if _, ok := db.organizations[idAsString]; ok {
		delete(db.organizations, idAsString)
		return nil
	}

	str := fmt.Sprintf("the organization (ID: %s) could not be found", idAsString)
	return errors.New(str)
}
