package blockchains

import (
	uuid "github.com/satori/go.uuid"
)

// Save represents a save blockchain transaction
type Save interface {
	GetID() *uuid.UUID
	GetOrganizationID() *uuid.UUID
	GetName() string
	GetDescription() string
	GetVoteTreshold() float64
}
