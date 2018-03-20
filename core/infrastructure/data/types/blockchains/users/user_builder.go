package users

import (
	"errors"
	"strconv"
	"time"

	cryptography "github.com/XMNBlockchain/exmachina-network/core/domain/cryptography"
	hashtrees "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/hashtrees"
	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/metadata"
	user "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/users"
	concrete_cryptography "github.com/XMNBlockchain/exmachina-network/core/infrastructure/cryptography/rsa"
	concrete_metadata "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/blockchains/metadata"
	uuid "github.com/satori/go.uuid"
)

type userBuilder struct {
	htBuilderFactory       hashtrees.HashTreeBuilderFactory
	metaDataBuilderFactory metadata.MetaDataBuilderFactory
	id                     *uuid.UUID
	met                    metadata.MetaData
	pub                    cryptography.PublicKey
	crOn                   *time.Time
}

func createUserBuilder(htBuilderFactory hashtrees.HashTreeBuilderFactory, metaDataBuilderFactory metadata.MetaDataBuilderFactory) user.UserBuilder {
	out := userBuilder{
		htBuilderFactory:       htBuilderFactory,
		metaDataBuilderFactory: metaDataBuilderFactory,
		id:   nil,
		met:  nil,
		pub:  nil,
		crOn: nil,
	}

	return &out
}

// Create initializes the UserBuilder instance
func (build *userBuilder) Create() user.UserBuilder {
	build.id = nil
	build.met = nil
	build.pub = nil
	build.crOn = nil
	return build
}

// WithID adds an ID to the UserBuilder
func (build *userBuilder) WithID(id uuid.UUID) user.UserBuilder {
	build.id = &id
	return build
}

// WithMetaData adds MetaData to the UserBuilder
func (build *userBuilder) WithMetaData(met metadata.MetaData) user.UserBuilder {
	build.met = met
	return build
}

// WithPublicKey adds a PublicKey to the UserBuilder
func (build *userBuilder) WithPublicKey(pub cryptography.PublicKey) user.UserBuilder {
	build.pub = pub
	return build
}

// CreatedOn adds a creation time to the UserBuilder
func (build *userBuilder) CreatedOn(crOn time.Time) user.UserBuilder {
	build.crOn = &crOn
	return build
}

// Now builds a new User instance
func (build *userBuilder) Now() (user.User, error) {

	if build.pub == nil {
		return nil, errors.New("the PublicKey is mandatory in order to build a User instance")
	}

	if build.met == nil {
		if build.id == nil {
			return nil, errors.New("the ID is mandatory in order to build a User instance")
		}

		if build.crOn == nil {
			return nil, errors.New("the creation time is mandatory in order to build a User instance")
		}

		pubKeyAsString, pubKeyAsStringErr := build.pub.String()
		if pubKeyAsStringErr != nil {
			return nil, pubKeyAsStringErr
		}

		blocks := [][]byte{
			build.id.Bytes(),
			[]byte(strconv.Itoa(int(build.crOn.UnixNano()))),
			[]byte(pubKeyAsString),
		}

		ht, htErr := build.htBuilderFactory.Create().Create().WithBlocks(blocks).Now()
		if htErr != nil {
			return nil, htErr
		}

		met, metErr := build.metaDataBuilderFactory.Create().Create().WithID(build.id).WithHashTree(ht).CreatedOn(*build.crOn).Now()
		if metErr != nil {
			return nil, metErr
		}

		build.met = met
	}

	if build.met == nil {
		return nil, errors.New("the MetaData is mandatory in order to build a User instance")
	}

	out := createUser(build.met.(*concrete_metadata.MetaData), build.pub.(*concrete_cryptography.PublicKey))
	return out, nil
}
