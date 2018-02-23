package users

import (
	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
	uuid "github.com/satori/go.uuid"
)

// SaveBuilder represents the builder of a save user transaction
type SaveBuilder interface {
	Create() SaveBuilder
	WithID(id *uuid.UUID) SaveBuilder
	WithPublicKey(pk cryptography.PublicKey) SaveBuilder
	Now() (Save, error)
}
