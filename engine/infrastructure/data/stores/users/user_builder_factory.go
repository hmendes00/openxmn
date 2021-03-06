package users

import (
	stored_users "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/users"
)

// UserBuilderFactory represents a concrete UserBuilderFactory implementation
type UserBuilderFactory struct {
}

// CreateUserBuilderFactory creates a new UserBuilderFactory instance
func CreateUserBuilderFactory() stored_users.UserBuilderFactory {
	out := UserBuilderFactory{}
	return &out
}

// Create creates a new UserBuilder instance
func (fac *UserBuilderFactory) Create() stored_users.UserBuilder {
	out := createUserBuilder()
	return out
}
