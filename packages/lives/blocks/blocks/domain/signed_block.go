package domain

import (
	"time"

	users "github.com/XMNBlockchain/core/packages/users/domain"
	uuid "github.com/satori/go.uuid"
)

// SignedBlock represents a SignedBlock instance
type SignedBlock interface {
	GetID() *uuid.UUID
	GetBlock() Block
	GetSignature() users.Signature
	CreatedOn() time.Time
}
