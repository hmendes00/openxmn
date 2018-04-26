package write

import (
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// User represents a user write database
type User struct {
	users map[string]users.User
}

// CreateUser creates a new User instance
func CreateUser(users map[string]users.User) *User {
	out := User{
		users: users,
	}

	return &out
}

// Insert inserts a new user
func (db *User) Insert(user users.User) error {
	return nil
}

// Update updates an existing user
func (db *User) Update(original users.User, new users.User) error {
	return nil
}

// Delete deletes an existing user
func (db *User) Delete(user users.User) error {
	return nil
}
