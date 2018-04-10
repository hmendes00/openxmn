package databases

import (
	files "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands/files"
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

// Insert inserts a new user to the database
func (db *User) Insert(usr users.User) (files.File, error) {
	return nil, nil
}

// Update updates a user to the database
func (db *User) Update(old users.User, new users.User) (files.File, files.File, error) {
	return nil, nil, nil
}
