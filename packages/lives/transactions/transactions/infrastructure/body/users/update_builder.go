package users

import (
	"errors"

	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
	concrete_cryptography "github.com/XMNBlockchain/core/packages/cryptography/infrastructure/rsa"
	"github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain/body/users"
	uuid "github.com/satori/go.uuid"
)

type updateBuilder struct {
	id *uuid.UUID
	pk cryptography.PublicKey
}

func createUpdateBuilder() users.UpdateBuilder {
	out := updateBuilder{
		id: nil,
		pk: nil,
	}

	return &out
}

// Create initializes the create builder
func (build *updateBuilder) Create() users.UpdateBuilder {
	build.id = nil
	build.pk = nil
	return build
}

// WithID adds an ID the create builder
func (build *updateBuilder) WithID(id *uuid.UUID) users.UpdateBuilder {
	build.id = id
	return build
}

// WithPublicKey adds a PublicKey the create builder
func (build *updateBuilder) WithNewPublicKey(pk cryptography.PublicKey) users.UpdateBuilder {
	build.pk = pk
	return build
}

// Now builds a new Create instance
func (build *updateBuilder) Now() (users.Update, error) {

	if build.id == nil {
		return nil, errors.New("the ID is mandatory in order to build a create instance")
	}

	if build.pk == nil {
		return nil, errors.New("the PublicKey is mandatory in order to build a create instance")
	}

	out := createUpdate(build.id, build.pk.(*concrete_cryptography.PublicKey))
	return out, nil
}
