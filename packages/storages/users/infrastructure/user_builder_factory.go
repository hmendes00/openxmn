package infrastructure

import (
	stored_users "github.com/XMNBlockchain/core/packages/storages/users/domain"
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
