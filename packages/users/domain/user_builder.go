package domain

import (
	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
	uuid "github.com/satori/go.uuid"
)

// UserBuilder represents a User builder
type UserBuilder interface {
	Create() UserBuilder
	WithID(id uuid.UUID) UserBuilder
	WithPublicKey(pub cryptography.PublicKey) UserBuilder
	Now() (User, error)
}
