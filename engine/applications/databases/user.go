package databases

import (
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	uuid "github.com/satori/go.uuid"
)

// User represents a user database
type User struct {
	dirPath    string
	repository users.UserRepository
}

// CreateUser creates a user database
func CreateUser(dirPath string, repository users.UserRepository) *User {
	out := User{
		dirPath:    dirPath,
		repository: repository,
	}

	return &out
}

// RetrieveByID retrieves a user by ID
func (db *User) RetrieveByID(id *uuid.UUID) (users.User, error) {
	return nil, nil
}
