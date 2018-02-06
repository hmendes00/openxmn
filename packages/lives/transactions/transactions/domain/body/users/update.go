package users

import (
	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
	uuid "github.com/satori/go.uuid"
)

// Update represents a transaction to update a user
type Update interface {
	GetID() *uuid.UUID
	GetNewPublicKey() cryptography.PublicKey
}
