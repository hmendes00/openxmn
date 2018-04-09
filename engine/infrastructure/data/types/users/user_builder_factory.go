package users

import (
	hashtrees "github.com/XMNBlockchain/openxmn/engine/domain/data/types/hashtrees"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	user "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// UserBuilderFactory represents a concrete UserBuilderFactory
type UserBuilderFactory struct {
	htBuilderFactory       hashtrees.HashTreeBuilderFactory
	metaDataBuilderFactory metadata.BuilderFactory
}

// CreateUserBuilderFactory creates a new UserBuilderFactory instance
func CreateUserBuilderFactory(htBuilderFactory hashtrees.HashTreeBuilderFactory, metaDataBuilderFactory metadata.BuilderFactory) user.UserBuilderFactory {
	out := UserBuilderFactory{
		htBuilderFactory:       htBuilderFactory,
		metaDataBuilderFactory: metaDataBuilderFactory,
	}
	return &out
}

// Create creates a new UserBuilder instance
func (fac *UserBuilderFactory) Create() user.UserBuilder {
	out := createUserBuilder(fac.htBuilderFactory, fac.metaDataBuilderFactory)
	return out
}
