package users

import (
	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
	uuid "github.com/satori/go.uuid"
)

// UpdateBuilder represents the builder of an update user transaction
type UpdateBuilder interface {
	Create() UpdateBuilder
	WithID(id *uuid.UUID) UpdateBuilder
	WithNewPublicKey(pk cryptography.PublicKey) UpdateBuilder
	Now() (Update, error)
}
