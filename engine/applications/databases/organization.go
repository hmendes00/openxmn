package databases

import (
	"errors"
	"fmt"

	organizations "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations"
	uuid "github.com/satori/go.uuid"
)

// Organization represents the organization database
type Organization struct {
	orgs map[string]organizations.Organization
}

// CreateOrganization creates an organization database
func CreateOrganization() *Organization {
	out := Organization{
		orgs: map[string]organizations.Organization{},
	}

	return &out
}

// RetrieveByID retrieves an organization by ID
func (db *Organization) RetrieveByID(id *uuid.UUID) (organizations.Organization, error) {
	idAsString := id.String()
	if oneOrg, ok := db.orgs[idAsString]; ok {
		return oneOrg, nil
	}

	str := fmt.Sprintf("the organization (ID: %s) could not be found", idAsString)
	return nil, errors.New(str)
}

// Insert inserts a new organizartion to the database
func (db *Organization) Insert(org organizations.Organization) error {
	id := org.GetMetaData().GetID()
	idAsString := id.String()
	_, retOrgErr := db.RetrieveByID(id)
	if retOrgErr == nil {
		str := fmt.Sprintf("there is already an organization with ID: %s", idAsString)
		return errors.New(str)
	}

	db.orgs[idAsString] = org
	return nil
}

// Update uopdates an organization to the database
func (db *Organization) Update(old organizations.Organization, new organizations.Organization) error {
	newOrgID := new.GetMetaData().GetID()
	newOrgIDAsString := newOrgID.String()
	_, retNewOrgErr := db.RetrieveByID(newOrgID)
	if retNewOrgErr == nil {
		str := fmt.Sprintf("the new organization (ID: %s) already exists", newOrgIDAsString)
		return errors.New(str)
	}

	delErr := db.Delete(old)
	if delErr != nil {
		return delErr
	}

	insErr := db.Insert(new)
	if insErr != nil {
		return insErr
	}

	return nil
}

// Delete deletes an organizartion from the database
func (db *Organization) Delete(org organizations.Organization) error {
	id := org.GetMetaData().GetID()
	_, retOrgErr := db.RetrieveByID(id)
	if retOrgErr != nil {
		return retOrgErr
	}

	idAsString := id.String()
	delete(db.orgs, idAsString)
	return nil
}
