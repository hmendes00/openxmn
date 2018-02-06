package users

import (
	users "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain/body/users"
)

// UserBuilderFactory represents a concrete UserBuilderFactory
type UserBuilderFactory struct {
}

// CreateUserBuilderFactory creates a new UserBuilderFactory instance
func CreateUserBuilderFactory() users.BuilderFactory {
	out := UserBuilderFactory{}
	return &out
}

// Create creates a new UserBuilder instance
func (fac *UserBuilderFactory) Create() users.Builder {
	out := createUserBuilder()
	return out
}
