package infrastructure

import (
	user "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
)

// UserBuilderFactory represents a concrete UserBuilderFactory
type UserBuilderFactory struct {
}

// CreateUserBuilderFactory creates a new UserBuilderFactory instance
func CreateUserBuilderFactory() user.UserBuilderFactory {
	out := UserBuilderFactory{}
	return &out
}

// Create creates a new UserBuilder instance
func (fac *UserBuilderFactory) Create() user.UserBuilder {
	out := createUserBuilder()
	return out
}
