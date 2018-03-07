package domain

import (
	stored_users "github.com/XMNBlockchain/core/packages/storages/users/domain"
)

// UserService represents a user service
type UserService interface {
	Save(dirPath string, usr User) (stored_users.User, error)
}
