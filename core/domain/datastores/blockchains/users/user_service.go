package domain

import (
	stored_users "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/users"
)

// UserService represents a user service
type UserService interface {
	Save(dirPath string, usr User) (stored_users.User, error)
}
