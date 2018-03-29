package wealth

import (
	user "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// Entity represents a user or an organization entity
type Entity struct {
	Usr *user.User    `json:"user"`
	Org *Organization `json:"organization"`
}

// CreateEntityWithUser creates an entity with a user
func CreateEntityWithUser(usr *user.User) *Entity {
	out := Entity{
		Usr: usr,
		Org: nil,
	}

	return &out
}

// HasUser returns true if there is a user, false otherwise
func (en *Entity) HasUser() bool {
	return en.Usr != nil
}

// GetUser returns the user, if any
func (en *Entity) GetUser() *user.User {
	return en.Usr
}

// HasOrganization returns true if there is an organization, false otherwise
func (en *Entity) HasOrganization() bool {
	return en.Org != nil
}

// GetOrganization returns the organization, if any
func (en *Entity) GetOrganization() *Organization {
	return en.Org
}
