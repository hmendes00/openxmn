package read

import (
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	objects "github.com/XMNBlockchain/openxmn/highway/project/objects"
	uuid "github.com/satori/go.uuid"
)

// Holder represents a holder database
type Holder struct {
	userDB *User
	orgDB  *Organization
}

// CreateHolder creates a new Holder database
func CreateHolder(userDB *User, orgDB *Organization) *Holder {
	out := Holder{
		userDB: userDB,
		orgDB:  orgDB,
	}

	return &out
}

// RetrieveByUserOrOrganizationID retrieves an Holder by its organizationID, if not empty, otherwise the given user
func (db *Holder) RetrieveByUserOrOrganizationID(usr users.User, orgID *uuid.UUID) (*objects.Holder, error) {
	return nil, nil
}
