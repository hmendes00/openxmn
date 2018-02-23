package accounts

import uuid "github.com/satori/go.uuid"

// Save represents a save account transaction
type Save interface {
	GetBranchID() *uuid.UUID
	GetCurrencyID() *uuid.UUID
	GetNumber() int
	GetOrganizationID() *uuid.UUID
	GetAddressID() *uuid.UUID
}
