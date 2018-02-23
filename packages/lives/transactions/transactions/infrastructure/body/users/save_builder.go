package users

import (
	"errors"

	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
	concrete_cryptography "github.com/XMNBlockchain/core/packages/cryptography/infrastructure/rsa"
	users "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain/body/users"
	uuid "github.com/satori/go.uuid"
)

type createBuilder struct {
	id *uuid.UUID
	pk cryptography.PublicKey
}

func createSaveBuilder() users.SaveBuilder {
	out := createBuilder{
		id: nil,
		pk: nil,
	}

	return &out
}

// Create initializes the create builder
func (build *createBuilder) Create() users.SaveBuilder {
	build.id = nil
	build.pk = nil
	return build
}

// WithID adds an ID the create builder
func (build *createBuilder) WithID(id *uuid.UUID) users.SaveBuilder {
	build.id = id
	return build
}

// WithPublicKey adds a PublicKey the create builder
func (build *createBuilder) WithPublicKey(pk cryptography.PublicKey) users.SaveBuilder {
	build.pk = pk
	return build
}

// Now builds a new Save instance
func (build *createBuilder) Now() (users.Save, error) {

	if build.id == nil {
		return nil, errors.New("the ID is mandatory in order to build a create instance")
	}

	if build.pk == nil {
		return nil, errors.New("the PublicKey is mandatory in order to build a create instance")
	}

	out := createSave(build.id, build.pk.(*concrete_cryptography.PublicKey))
	return out, nil
}
