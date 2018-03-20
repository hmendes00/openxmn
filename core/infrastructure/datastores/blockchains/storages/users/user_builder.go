package users

import (
	"errors"

	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/files"
	stored_users "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/users"
	concrete_stored_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/datastores/blockchains/storages/files"
)

type userBuilder struct {
	met    stored_files.File
	pubKey stored_files.File
}

func createUserBuilder() stored_users.UserBuilder {
	out := userBuilder{
		met:    nil,
		pubKey: nil,
	}

	return &out
}

// Create initializes the UserBuilder
func (build *userBuilder) Create() stored_users.UserBuilder {
	build.met = nil
	build.pubKey = nil
	return build
}

// WithMetaData adds MetaData to the UserBuilder
func (build *userBuilder) WithMetaData(met stored_files.File) stored_users.UserBuilder {
	build.met = met
	return build
}

// WithPublicKey adds a PublicKey to the UserBuilder
func (build *userBuilder) WithPublicKey(pk stored_files.File) stored_users.UserBuilder {
	build.pubKey = pk
	return build
}

// Now builds a new User instance
func (build *userBuilder) Now() (stored_users.User, error) {
	if build.met == nil {
		return nil, errors.New("the MetaData is mandatory in order to build a User instance")
	}

	if build.pubKey == nil {
		return nil, errors.New("the PublicKey is mandatory in order to build a User instance")
	}

	out := createUser(build.met.(*concrete_stored_files.File), build.pubKey.(*concrete_stored_files.File))
	return out, nil
}
