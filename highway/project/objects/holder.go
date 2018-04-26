package objects

import (
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	concrete_users "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/users"
)

// Holder represents a user or an organization that hold currencies or assets
type Holder struct {
	Usr *concrete_users.User `json:"user"`
	Org *Organization        `json:"organization"`
}

// CreateHolderWithUser creates a new Holder instance from a user instance
func CreateHolderWithUser(usr users.User) *Holder {
	out := Holder{
		Usr: usr.(*concrete_users.User),
		Org: nil,
	}

	return &out
}

// CreateHolderWithOrganization creates a new Holder instance from an organization instance
func CreateHolderWithOrganization(org *Organization) *Holder {
	out := Holder{
		Usr: nil,
		Org: org,
	}

	return &out
}
