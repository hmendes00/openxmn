package branches

import uuid "github.com/satori/go.uuid"

// Save represents a save branch transaction
type Save interface {
	GetID() *uuid.UUID
	GetOrganizationID() *uuid.UUID
	GetTrustID() *uuid.UUID
	GetAddressID() *uuid.UUID
	GetTransit() int
}
