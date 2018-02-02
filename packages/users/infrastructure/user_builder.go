package infrastructure

import (
	"errors"

	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
	concrete_cryptography "github.com/XMNBlockchain/core/packages/cryptography/infrastructure/rsa"
	user "github.com/XMNBlockchain/core/packages/users/domain"
	uuid "github.com/satori/go.uuid"
)

type userBuilder struct {
	id  *uuid.UUID
	pub cryptography.PublicKey
}

func createUserBuilder() user.UserBuilder {
	out := userBuilder{
		id:  nil,
		pub: nil,
	}

	return &out
}

// Create initializes the UserBuilder instance
func (build *userBuilder) Create() user.UserBuilder {
	build.id = nil
	build.pub = nil
	return build
}

// WithID adds an ID to the UserBuilder
func (build *userBuilder) WithID(id uuid.UUID) user.UserBuilder {
	build.id = &id
	return build
}

// WithPublicKey adds a PublicKey to the UserBuilder
func (build *userBuilder) WithPublicKey(pub cryptography.PublicKey) user.UserBuilder {
	build.pub = pub
	return build
}

// Now builds a new User instance
func (build *userBuilder) Now() (user.User, error) {

	if build.id == nil {
		return nil, errors.New("the id is mandatory in order to build a User instance")
	}

	if build.pub == nil {
		return nil, errors.New("the PublicKey is mandatory in order to build a User instance")
	}

	out := createUser(build.id, build.pub.(*concrete_cryptography.PublicKey))
	return out, nil
}
