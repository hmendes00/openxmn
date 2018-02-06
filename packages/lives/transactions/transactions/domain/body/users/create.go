package users

import (
	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
	uuid "github.com/satori/go.uuid"
)

// Create represents a create user transaction
type Create interface {
	GetID() *uuid.UUID
	GetPublicKey() cryptography.PublicKey
}
