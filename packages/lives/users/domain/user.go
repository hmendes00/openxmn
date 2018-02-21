package domain

import (
	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
	uuid "github.com/satori/go.uuid"
)

// User represents a container of coins
type User interface {
	GetID() *uuid.UUID
	GetPublicKey() cryptography.PublicKey
}
