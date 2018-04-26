package read

import (
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	uuid "github.com/satori/go.uuid"
)

// User represents a user read database
type User struct {
	usr map[string]users.User
}

// CreateUser creates a new User instance
func CreateUser(usr map[string]users.User) *User {
	out := User{
		usr: usr,
	}

	return &out
}

// RetrieveByID retrieves a user by ID
func (db *User) RetrieveByID(id *uuid.UUID) (users.User, error) {
	return nil, nil
}
