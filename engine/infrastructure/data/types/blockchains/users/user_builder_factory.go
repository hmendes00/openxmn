package users

import (
	hashtrees "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/hashtrees"
	metadata "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/blockchains/metadata"
	user "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/blockchains/users"
)

// UserBuilderFactory represents a concrete UserBuilderFactory
type UserBuilderFactory struct {
	htBuilderFactory       hashtrees.HashTreeBuilderFactory
	metaDataBuilderFactory metadata.MetaDataBuilderFactory
}

// CreateUserBuilderFactory creates a new UserBuilderFactory instance
func CreateUserBuilderFactory(htBuilderFactory hashtrees.HashTreeBuilderFactory, metaDataBuilderFactory metadata.MetaDataBuilderFactory) user.UserBuilderFactory {
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
