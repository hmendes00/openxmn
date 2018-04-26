package write

import (
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
func (db *Organization) Insert(org *objects.Organization) error {
	return nil
}

// Update updates an existing organization
func (db *Organization) Update(original *objects.Organization, new *objects.Organization) error {
	return nil
}

// Delete deletes an existing organization
func (db *Organization) Delete(org *objects.Organization) error {
	return nil
}
