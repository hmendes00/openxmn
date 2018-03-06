package infrastructure

import (
	hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/domain"
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	user "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
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
