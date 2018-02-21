package domain

import (
	"time"

	users "github.com/XMNBlockchain/core/packages/lives/users/domain"
	uuid "github.com/satori/go.uuid"
)

// MetaData represents an object metadata
type MetaData interface {
	GetID() *uuid.UUID
	CreatedOn() time.Time
	HasSignature() bool
	GetSignature() users.Signature
}
