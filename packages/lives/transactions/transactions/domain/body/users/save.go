package users

import (
	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
	uuid "github.com/satori/go.uuid"
)

// Save represents a save user transaction
type Save interface {
	GetID() *uuid.UUID
	GetPublicKey() cryptography.PublicKey
}
