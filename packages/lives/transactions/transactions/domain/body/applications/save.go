package applications

import uuid "github.com/satori/go.uuid"

// Save represents a save application transaction
type Save interface {
	GetID() *uuid.UUID
	BeginAtBlockIndex() int
	FromVersionID() int
	GetOrganizationID() *uuid.UUID
	GetBlockchainID() *uuid.UUID
	GetCode() string
}
