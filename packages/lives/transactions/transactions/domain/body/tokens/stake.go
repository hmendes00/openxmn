package tokens

import uuid "github.com/satori/go.uuid"

// Stake represents a stake of tokens to an organization
type Stake interface {
	GetID() *uuid.UUID
	FromUserID() *uuid.UUID
	ToOrganizationID() *uuid.UUID
	GetAmount() float64
}
