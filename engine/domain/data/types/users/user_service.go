package domain

import (
	stored_users "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/users"
)

// UserService represents a user service
type UserService interface {
	Save(dirPath string, usr User) (stored_users.User, error)
}