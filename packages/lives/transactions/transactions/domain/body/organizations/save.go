package organizations

import uuid "github.com/satori/go.uuid"

// Save represents a transaction used to save an organization
type Save interface {
	GetID() *uuid.UUID
	GetTokenID() *uuid.UUID
	GetUserID() *uuid.UUID
	HasParentOrganizationID() bool
	GetParentOrganizationID() *uuid.UUID
	GetName() string
	GetDescription() string
}
